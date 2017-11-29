package main

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"time"

	"database/sql"

	"github.com/EVE-Tools/element43/go/lib/transport"
	"github.com/EVE-Tools/market-stats/lib/types"
	"github.com/antihax/goesi"
	"github.com/antihax/goesi/esi"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/montanaflynn/stats"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	_ "gopkg.in/mattes/migrate.v1/driver/postgres"
	"gopkg.in/mattes/migrate.v1/migrate"
)

// Config holds the application's configuration info from the environment.
type Config struct {
	SeedDB      bool   `default:"false" envconfig:"seed_db"`
	Cron        string `default:"0 52 0 * * *" envconfig:"cron"`
	LogLevel    string `default:"debug" split_words:"true"`
	PostgresURL string `default:"postgres://market-stats@localhost:5432/market-stats?sslmode=disable" envconfig:"postgres_url"`
	Port        string `default:"8000" envconfig:"port"`
}

var db *sql.DB

// For limiting requests to type API
var esiClient goesi.APIClient
var esiSemaphore chan struct{}

func main() {
	const userAgent string = "Element43/market-stats (element-43.com)"
	const timeout time.Duration = time.Duration(time.Second * 30)

	httpClient := &http.Client{
		Timeout:   timeout,
		Transport: transport.NewESITransport(userAgent, timeout),
	}

	esiClient = *goesi.NewAPIClient(httpClient, "Element43/market-stats (element-43.com)")

	config := loadConfig()
	connectToDB(config)
	migrateDB(config)
	startUpdateCron(config)

	// Seed DB if flag is set
	if config.SeedDB {
		go updateHistoryStats()
	}

	startWebServer(config)

	// Terminate this goroutine, crash if all other goroutines exited
	runtime.Goexit()
}

// Initialize databse connection pool
func connectToDB(config Config) {
	var err error
	db, err = sql.Open("postgres", config.PostgresURL)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(50)
}

// Check database and migrate if needed.
func migrateDB(config Config) {
	errors, ok := migrate.UpSync(config.PostgresURL, "./migrations")
	if !ok {
		logrus.Error("Migrating the database failed!")
		for _, err := range errors {
			logrus.Error(err.Error())
		}
		panic(0)
	}
}

// Schedule full update every night 5 minutes after stat generation
func startUpdateCron(config Config) {
	job := cron.New()
	job.AddFunc(config.Cron, updateHistoryStats)
	job.Start()
}

// Seed the database with all region/types available
func updateHistoryStats() {
	// Get all regions
	// Get all types
	// Filter for types on market
	// Store and calculate stats in pipeline
	logrus.Info("Started market update.")
	updateStart := time.Now()

	regionIDs, err := getMarketRegions()
	if err != nil {
		logrus.Warnf("Failed to get regions from ESI: %s", err.Error())
		return
	}

	typeIDs, err := getMarketTypes()
	if err != nil {
		logrus.Warnf("Failed to get types from ESI: %s", err.Error())
		return
	}

	logrus.Debugf("Got %v regions and %v types", len(regionIDs), len(typeIDs))

	regionTypes := makeRegionTypes(regionIDs, typeIDs)

	generateStats(regionTypes)

	timeElapsed := time.Since(updateStart)
	logrus.WithFields(logrus.Fields{
		"time": timeElapsed,
	}).Info("Finished market update.")
}

// Get all regionIDs from ESI
func getRegionIDs() ([]int32, error) {
	regionIDs, _, err := esiClient.ESI.UniverseApi.GetUniverseRegions(nil, nil)
	if err != nil {
		return nil, err
	}

	return regionIDs, nil
}

// Get all regions with a market (filter WH)
func getMarketRegions() ([]int32, error) {
	regionIDs, err := getRegionIDs()
	if err != nil {
		return nil, err
	}

	var marketRegionIDs []int32
	for _, regionID := range regionIDs {
		if regionID < 11000000 {
			marketRegionIDs = append(marketRegionIDs, regionID)
		}
	}

	return marketRegionIDs, nil
}

// Get all typeIDs from ESI
// TODO: move to static-data RPC
func getTypeIDs() ([]int32, error) {
	var typeIDs []int32
	params := make(map[string]interface{})
	params["page"] = int32(1)

	typeResult, _, err := esiClient.ESI.UniverseApi.GetUniverseTypes(nil, params)
	if err != nil {
		return nil, err
	}

	typeIDs = append(typeIDs, typeResult...)

	for len(typeResult) > 0 {
		params["page"] = params["page"].(int32) + 1
		typeResult, _, err = esiClient.ESI.UniverseApi.GetUniverseTypes(nil, params)
		if err != nil {
			return nil, err
		}

		typeIDs = append(typeIDs, typeResult...)
	}

	return typeIDs, nil
}

// Get all types on market
// TODO: move to static-data RPC
func getMarketTypes() ([]int32, error) {
	typeIDs, err := getTypeIDs()
	if err != nil {
		return nil, err
	}

	marketTypes := make(chan int32)
	nonMarketTypes := make(chan int32)
	failure := make(chan error)

	typesLeft := len(typeIDs)

	for _, id := range typeIDs {
		go checkIfMarketTypeAsyncRetry(id, marketTypes, nonMarketTypes, failure)
	}

	var marketTypeIDs []int32

	for typesLeft > 0 {
		select {
		case typeID := <-marketTypes:
			marketTypeIDs = append(marketTypeIDs, typeID)
		case <-nonMarketTypes:
		case err := <-failure:
			logrus.Warnf("Error fetching type from ESI: %s", err.Error())
		}

		typesLeft--
	}

	return marketTypeIDs, nil
}

// Async check if market type, retry 3 times
func checkIfMarketTypeAsyncRetry(typeID int32, marketTypes chan int32, nonMarketTypes chan int32, failure chan error) {
	var isMarketType bool
	var err error
	retries := 3

	for retries > 0 {
		isMarketType, err = checkIfMarketType(typeID)
		if err != nil {
			retries--
		} else {
			err = nil
			retries = 0
		}
	}

	if err != nil {
		failure <- err
		return
	}

	if isMarketType {
		marketTypes <- typeID
		return
	}

	nonMarketTypes <- typeID
}

// Check if type is market type
func checkIfMarketType(typeID int32) (bool, error) {
	esiSemaphore <- struct{}{}
	typeInfo, _, err := esiClient.ESI.UniverseApi.GetUniverseTypesTypeId(nil, typeID, nil)
	<-esiSemaphore
	if err != nil {
		return false, err
	}

	// If it is published and has a market group it is a market type!
	if typeInfo.Published && (typeInfo.MarketGroupId != 0) {
		return true, nil
	}

	return false, nil
}

// Make region types from region and type typeIDs
func makeRegionTypes(regionIDs []int32, typeIDs []int32) []types.RegionType {
	var regionTypes []types.RegionType

	for _, regionID := range regionIDs {
		for _, typeID := range typeIDs {
			regionTypes = append(regionTypes, types.RegionType{
				RegionID: regionID,
				TypeID:   typeID,
			})
		}
	}

	return regionTypes
}

// Start the webserver.
func startWebServer(config Config) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(cors.Default())
	router.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true))

	v1 := router.Group("/api/market-stats/v1/history")
	v1.GET("/region/:regionID/", getRegion)
	v1.GET("/type/:typeID/", getType)
	v1.GET("/region/:regionID/type/:typeID/", getRegionType)

	router.Run(":" + config.Port)
}

// Get data from ESI and store in Postgres
func generateStats(regionTypes []types.RegionType) {

	// Stages
	// 1: Calculate ESI
	// 2: Bulk results
	// 3: DB upserter

	// 800k items ---> 1000 x ESI + calculation --(10,000)--> 1 x packer to 1,000 --(50)--> 50 x DB upserter

	logrus.Debugf("Processing %d regionTypes.", len(regionTypes))

	// Create backlog channel and push all regionTypes
	var backlog = make(chan *types.RegionType, len(regionTypes))
	for _, workItem := range regionTypes {
		var item = workItem
		backlog <- &item
	}

	// Create ESI download/calculator response, termination and data forward channels
	const esiWorkers = 200
	var esiItemsLeft = len(regionTypes)
	var esiDone = make(chan error, esiWorkers)
	var esiTerminate = make(chan struct{})
	var bulkBacklog = make(chan *types.RegionStats, 5000)

	for i := 0; i < esiWorkers; i++ {
		go esiWorker(backlog, esiDone, esiTerminate, bulkBacklog)
	}

	// Create bulk packer done, termination and data forward channels
	const bulkSize = 500
	var bulkItemsLeft = len(regionTypes)
	var bulkDone = make(chan int, 100)
	var bulkTerminate = make(chan struct{})
	var dbBacklog = make(chan []*types.RegionStats, 50)

	go packRegionStats(bulkSize, bulkBacklog, bulkDone, bulkTerminate, dbBacklog)

	// Create database workers and done channels
	const dbWorkers = 30
	var dbItemsLeft = len(regionTypes)
	var dbDone = make(chan int, dbWorkers)
	var dbTerminate = make(chan struct{})

	for i := 0; i < dbWorkers; i++ {
		go dbWorker(dbBacklog, dbDone, dbTerminate)
	}

	ticker := time.NewTicker(time.Second)

PipelineMonitor:
	for {
		select {
		case esiStatus := <-esiDone:
			esiItemsLeft--
			if esiStatus != nil {
				bulkItemsLeft--
				dbItemsLeft--
				if esiStatus.Error() != "Not enough datapoints" {
					logrus.WithError(esiStatus).Warn("Error getting stats from ESI.")
				}
				if esiItemsLeft < 1 && bulkItemsLeft < 1 && dbItemsLeft < 1 {
					break PipelineMonitor
				}
			}
		case bulkProcessed := <-bulkDone:
			bulkItemsLeft = bulkItemsLeft - bulkProcessed
		case dbProcessed := <-dbDone:
			dbItemsLeft = dbItemsLeft - dbProcessed
			if esiItemsLeft < 1 && bulkItemsLeft < 1 && dbItemsLeft < 1 {
				break PipelineMonitor
			}
		case <-ticker.C:
			logrus.WithFields(logrus.Fields{
				"esi_backlog":  esiItemsLeft,
				"bulk_waiting": bulkItemsLeft - esiItemsLeft,
				"bulk_backlog": bulkItemsLeft,
				"db_waiting":   dbItemsLeft - bulkItemsLeft,
				"db_backlog":   dbItemsLeft,
			}).Info("Processing markets.")
		}
	}

	ticker.Stop()

	logrus.Info("[1/3] Terminating ESI stage.")
	for i := 0; i < esiWorkers; i++ {
		esiTerminate <- struct{}{}
	}

	logrus.Info("[2/3] Terminating bulk stage.")
	bulkTerminate <- struct{}{}

	logrus.Info("[3/3] Terminating db stage.")
	for i := 0; i < dbWorkers; i++ {
		dbTerminate <- struct{}{}
	}
}

// Fetch a regionType from queue and get data/calculate statistics, then forward to bulk. Run until terminated.
func esiWorker(backlog <-chan *types.RegionType, done chan<- error, terminate <-chan struct{}, bulkBacklog chan<- *types.RegionStats) {
	for {
		select {
		case regionType := <-backlog:
			history, err := downloadStats(regionType)
			if err != nil {
				done <- err
				break
			}

			regionStats, err := calculateStats(regionType, history)
			if err != nil {
				done <- err
				break
			}

			bulkBacklog <- regionStats
			done <- nil
		case <-terminate:
			return
		}
	}
}

// Download stats from ESI - try three times
func downloadStats(regionType *types.RegionType) ([]esi.GetMarketsRegionIdHistory200Ok, error) {
	var err error

	for retries := 3; retries > 0; retries-- {
		history, _, err := esiClient.ESI.MarketApi.GetMarketsRegionIdHistory(nil, regionType.RegionID, regionType.TypeID, nil)

		if err == nil {
			return history, nil
		}
	}

	return nil, err
}

// Calculate stats
func calculateStats(regionType *types.RegionType, history []esi.GetMarketsRegionIdHistory200Ok) (*types.RegionStats, error) {
	retval := types.RegionStats{}
	retval.RegionID = regionType.RegionID
	retval.TypeID = regionType.TypeID
	retval.GeneratedAt = time.Now()

	aWeekAgo := time.Now().AddDate(0, 0, -8)
	var datapoints []esi.GetMarketsRegionIdHistory200Ok

	var priceSeries []float64
	var volumeSeries []float64
	var iskVolumeSeries []float64
	var orderCountSeries []float64

	for _, point := range history {

		historyDate, err := time.Parse("2006-01-02", point.Date)
		if err != nil {
			return nil, err
		}

		if aWeekAgo.Before(historyDate) {
			datapoints = append(datapoints, point)
			priceSeries = append(priceSeries, float64(point.Average))
			volumeSeries = append(volumeSeries, float64(point.Volume))
			iskVolumeSeries = append(iskVolumeSeries, float64(point.Average)*float64(point.Volume))
			orderCountSeries = append(orderCountSeries, float64(point.OrderCount))
		}
	}

	if len(datapoints) > 7 {
		logrus.Info(len(datapoints))
		logrus.Warn(retval)
	}

	if len(datapoints) < 2 {
		return nil, errors.New("Not enough datapoints")
	}

	// Get most recent point
	currentStats := datapoints[len(datapoints)-1]

	currentStatsDate, err := time.Parse("2006-01-02", currentStats.Date)
	if err != nil {
		return nil, err
	}

	retval.Date = currentStatsDate
	retval.Average = currentStats.Average
	retval.Highest = currentStats.Highest
	retval.Lowest = currentStats.Lowest
	retval.Volume = currentStats.Volume
	retval.OrderCount = currentStats.OrderCount

	// Get previous (most of the time this should be yesterday) point
	yesterdaysStats := datapoints[len(datapoints)-2]

	yesterdaysStatsDate, err := time.Parse("2006-01-02", yesterdaysStats.Date)
	if err != nil {
		return nil, err
	}

	retval.PreviousDate = yesterdaysStatsDate
	retval.PreviousAverage = yesterdaysStats.Average
	retval.PreviousHighest = yesterdaysStats.Highest
	retval.PreviousLowest = yesterdaysStats.Lowest
	retval.PreviousVolume = yesterdaysStats.Volume
	retval.PreviousOrderCount = yesterdaysStats.OrderCount

	//
	// Calculate stats
	//

	// Order Counts
	sum, err := stats.Sum(orderCountSeries)
	if err != nil {
		return nil, err
	}

	stdDeviation, err := stats.StandardDeviation(orderCountSeries)
	if err != nil {
		return nil, err
	}

	retval.WeekOrderCountTotal = int64(sum)
	retval.WeekOrderCountAverage = sum / float64(len(orderCountSeries))
	retval.WeekOrderCountStandardDeviation = stdDeviation
	retval.WeekOrderCountRelativeStandardDeviation = stdDeviation / retval.WeekOrderCountAverage

	// Volume
	sum, err = stats.Sum(volumeSeries)
	if err != nil {
		return nil, err
	}

	stdDeviation, err = stats.StandardDeviation(volumeSeries)
	if err != nil {
		return nil, err
	}

	retval.WeekVolumeTotal = int64(sum)
	retval.WeekVolumeAverage = sum / float64(len(volumeSeries))
	retval.WeekVolumeStandardDeviation = stdDeviation
	retval.WeekVolumeRelativeStandardDeviation = stdDeviation / retval.WeekVolumeAverage

	// ISK Volume
	sumISK, err := stats.Sum(iskVolumeSeries)
	if err != nil {
		return nil, err
	}

	stdDeviationISK, err := stats.StandardDeviation(iskVolumeSeries)
	if err != nil {
		return nil, err
	}

	retval.WeekISKVolumeAverage = sumISK / float64(len(iskVolumeSeries))
	retval.WeekISKVolumeAverageStandardDeviation = stdDeviationISK
	retval.WeekISKVolumeAverageRelativeStandardDeviation = stdDeviationISK / retval.WeekISKVolumeAverage

	// Prices
	sum, err = stats.Sum(priceSeries)
	if err != nil {
		return nil, err
	}

	stdDeviation, err = stats.StandardDeviation(priceSeries)
	if err != nil {
		return nil, err
	}

	retval.WeekPriceWeightedAverage = sumISK / float64(len(priceSeries)) / retval.WeekVolumeAverage
	retval.WeekPriceAverage = sum / float64(len(priceSeries))
	retval.WeekPriceAverageStandardDeviation = stdDeviation
	retval.WeekPriceAverageRelativeStandardDeviation = stdDeviation / retval.WeekPriceAverage

	return &retval, nil
}

// Take size region stats or wait a second, then forward items
func packRegionStats(size int, backlog <-chan *types.RegionStats, done chan<- int, terminate <-chan struct{}, dbBacklog chan<- []*types.RegionStats) {
	for {
		var packet []*types.RegionStats
	PackageCollector:
		for {
			select {
			case regionStat := <-backlog:
				packet = append(packet, regionStat)
				if len(packet) == size {
					dbBacklog <- packet
					done <- len(packet)
					break PackageCollector
				}
			case <-time.After(time.Second * 1):
				if len(packet) > 0 {
					dbBacklog <- packet
					done <- len(packet)
					break PackageCollector
				}
			case <-terminate:
				return
			}
		}
	}
}

// Store item in DB
func dbWorker(backlog <-chan []*types.RegionStats, done chan<- int, terminate <-chan struct{}) {
	for {
		select {
		case packet := <-backlog:
			err := upsertStats(packet)
			if err != nil {
				logrus.WithError(err).Warnf("Could not store stats.")
			}
			done <- len(packet)
		case <-terminate:
			return
		}
	}
}

// Build the query for upserting the stats
func buildQuery(regionStats []*types.RegionStats) (string, []interface{}) {
	placeholders := make([]string, 0, len(regionStats))
	values := make([]interface{}, 0, len(regionStats)*30)
	counter := 1

	for _, stat := range regionStats {
		var positionals []string

		for i := 0; i < 30; i++ {
			positionals = append(positionals, fmt.Sprintf("$%d", counter+i))
		}

		counter = counter + 30

		newPlaceholders := fmt.Sprintf("(%s)", strings.Join(positionals, ","))
		placeholders = append(placeholders, newPlaceholders)
		values = append(values, stat.RegionID)
		values = append(values, stat.TypeID)
		values = append(values, stat.GeneratedAt)
		values = append(values, stat.Date)
		values = append(values, stat.Highest)
		values = append(values, stat.Lowest)
		values = append(values, stat.Average)
		values = append(values, stat.Volume)
		values = append(values, stat.OrderCount)
		values = append(values, stat.PreviousDate)
		values = append(values, stat.PreviousHighest)
		values = append(values, stat.PreviousLowest)
		values = append(values, stat.PreviousAverage)
		values = append(values, stat.PreviousVolume)
		values = append(values, stat.PreviousOrderCount)
		values = append(values, stat.WeekPriceWeightedAverage)
		values = append(values, stat.WeekPriceAverage)
		values = append(values, stat.WeekPriceAverageStandardDeviation)
		values = append(values, stat.WeekPriceAverageRelativeStandardDeviation)
		values = append(values, stat.WeekISKVolumeAverage)
		values = append(values, stat.WeekISKVolumeAverageStandardDeviation)
		values = append(values, stat.WeekISKVolumeAverageRelativeStandardDeviation)
		values = append(values, stat.WeekOrderCountTotal)
		values = append(values, stat.WeekOrderCountAverage)
		values = append(values, stat.WeekOrderCountStandardDeviation)
		values = append(values, stat.WeekOrderCountRelativeStandardDeviation)
		values = append(values, stat.WeekVolumeTotal)
		values = append(values, stat.WeekVolumeAverage)
		values = append(values, stat.WeekVolumeStandardDeviation)
		values = append(values, stat.WeekVolumeRelativeStandardDeviation)
	}

	query := fmt.Sprintf(`INSERT INTO stats ("region_id", "type_id", "generated_at", "date", "highest", "lowest", "average", "volume", "order_count", "previous_date", "previous_highest", "previous_lowest", "previous_average", "previous_volume", "previous_order_count", "week_price_weighted_average", "week_price_average", "week_price_average_standard_deviation", "week_price_average_relative_standard_deviation", "week_isk_volume_average", "week_isk_volume_average_standard_deviation", "week_isk_volume_average_relative_standard_deviation", "week_order_count_total", "week_order_count_average", "week_order_count_standard_deviation", "week_order_count_relative_standard_deviation", "week_volume_total", "week_volume_average", "week_volume_standard_deviation", "week_volume_relative_standard_deviation") VALUES %s ON CONFLICT ("region_id", "type_id") DO UPDATE SET "generated_at" = EXCLUDED."generated_at", "date" = EXCLUDED."date", "highest" = EXCLUDED."highest", "lowest" = EXCLUDED."lowest", "average" = EXCLUDED."average", "volume" = EXCLUDED."volume", "order_count" = EXCLUDED."order_count", "previous_date" = EXCLUDED."previous_date", "previous_highest" = EXCLUDED."previous_highest", "previous_lowest" = EXCLUDED."previous_lowest", "previous_average" = EXCLUDED."previous_average", "previous_volume" = EXCLUDED."previous_volume", "previous_order_count" = EXCLUDED."previous_order_count", "week_price_weighted_average" = EXCLUDED."week_price_weighted_average", "week_price_average" = EXCLUDED."week_price_average", "week_price_average_standard_deviation" = EXCLUDED."week_price_average_standard_deviation", "week_price_average_relative_standard_deviation" = EXCLUDED."week_price_average_relative_standard_deviation", "week_isk_volume_average" = EXCLUDED."week_isk_volume_average", "week_isk_volume_average_standard_deviation" = EXCLUDED."week_isk_volume_average_standard_deviation", "week_isk_volume_average_relative_standard_deviation" = EXCLUDED."week_isk_volume_average_relative_standard_deviation", "week_order_count_total" = EXCLUDED."week_order_count_total", "week_order_count_average" = EXCLUDED."week_order_count_average", "week_order_count_standard_deviation" = EXCLUDED."week_order_count_standard_deviation", "week_order_count_relative_standard_deviation" = EXCLUDED."week_order_count_relative_standard_deviation", "week_volume_total" = EXCLUDED."week_volume_total", "week_volume_average" = EXCLUDED."week_volume_average", "week_volume_standard_deviation" = EXCLUDED."week_volume_standard_deviation", "week_volume_relative_standard_deviation" = EXCLUDED."week_volume_relative_standard_deviation"`, strings.Join(placeholders, ","))

	return query, values
}

// Upsert the stats
func upsertStats(regionStats []*types.RegionStats) error {
	query, values := buildQuery(regionStats)

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, values...)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func getRegion(context *gin.Context) {
	regionID := context.Param("regionID")

	query := `SELECT row_to_json(row) FROM (SELECT * FROM "stats" WHERE "region_id" = $1) row`

	rows, err := db.Query(query, regionID)
	defer rows.Close()
	if err != nil {
		context.AbortWithError(500, err)
		return
	}

	numRows, markets, err := convertRowsToJSON(rows)
	if err != nil {
		context.AbortWithError(500, err)
	}

	if numRows == 0 {
		context.AbortWithStatus(404)
		return
	}

	context.Header("Access-Control-Allow-Origin", "*")
	context.Data(200, "application/json; charset=utf-8", markets)
}

func getType(context *gin.Context) {
	typeID := context.Param("typeID")

	query := `SELECT row_to_json(rows) FROM (SELECT * FROM "stats" WHERE "type_id" = $1) rows`

	rows, err := db.Query(query, typeID)
	defer rows.Close()
	if err != nil {
		context.AbortWithError(500, err)
		return
	}

	numRows, markets, err := convertRowsToJSON(rows)
	if err != nil {
		context.AbortWithError(500, err)
	}

	if numRows == 0 {
		context.AbortWithStatus(404)
		return
	}

	context.Header("Access-Control-Allow-Origin", "*")
	context.Data(200, "application/json; charset=utf-8", markets)
}

func getRegionType(context *gin.Context) {
	regionID := context.Param("regionID")
	typeID := context.Param("typeID")

	query := `SELECT row_to_json(row) FROM (SELECT * FROM "stats" WHERE "region_id" = $1 AND "type_id" = $2) row`

	rows, err := db.Query(query, regionID, typeID)
	defer rows.Close()
	if err != nil {
		context.AbortWithError(500, err)
		return
	}

	var market []byte

	for rows.Next() {
		err = rows.Scan(&market)
		if err != nil {
			context.AbortWithError(500, err)
		}
	}

	if len(market) == 0 {
		context.AbortWithStatus(404)
		return
	}

	context.Header("Access-Control-Allow-Origin", "*")
	context.Data(200, "application/json; charset=utf-8", market)
}

func convertRowsToJSON(rows *sql.Rows) (int, []byte, error) {
	defer rows.Close()
	markets := []byte("[")
	numRows := 0

	for rows.Next() {
		var market []byte
		numRows++
		err := rows.Scan(&market)
		if err != nil {
			return numRows, nil, err
		}

		markets = append(markets, market...)
		markets = append(markets, ","...)
	}
	markets = markets[:len(markets)-1]
	markets = append(markets, "]"...)

	return numRows, markets, nil
}

// Load configuration from environment
func loadConfig() Config {
	config := Config{}
	envconfig.MustProcess("MARKET_STATS", &config)

	logLevel, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		panic(err)
	}

	esiSemaphore = make(chan struct{}, 200)

	logrus.SetLevel(logLevel)
	logrus.Debugf("Config: %q", config)
	return config
}
