// FIXME: This file is way too large and should be split into modules!

package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"runtime"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/EVE-Tools/element43/go/lib/transport"
	pb "github.com/EVE-Tools/market-stats/lib/marketStats"
	"github.com/EVE-Tools/market-stats/lib/types"
	"github.com/antihax/goesi"
	"github.com/antihax/goesi/esi"
	"github.com/antihax/goesi/optional"
	"github.com/golang/protobuf/ptypes"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/montanaflynn/stats"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	_ "gopkg.in/mattes/migrate.v1/driver/postgres"
	"gopkg.in/mattes/migrate.v1/migrate"
)

// Config holds the application's configuration info from the environment.
type Config struct {
	SeedDB      bool   `default:"false" envconfig:"seed_db"`
	Cron        string `default:"0 20 11 * * *" envconfig:"cron"`
	LogLevel    string `default:"debug" split_words:"true"`
	PostgresURL string `default:"postgres://market-stats@localhost:5432/market-stats?sslmode=disable" envconfig:"postgres_url"`
	Port        string `default:"43000" envconfig:"port"`
}

// Our server instance type
type marketStatsServer struct{}

// Global instances
// TODO: make these dependencies explicit!
var db *sql.DB

// For limiting requests to type API
var esiClient goesi.APIClient

func main() {
	const userAgent string = "Element43/market-stats (element-43.com)"
	const timeout time.Duration = time.Duration(time.Second * 30)

	httpClient := &http.Client{
		Timeout:   timeout,
		Transport: transport.NewESITransport(userAgent, timeout, 500),
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

	startGrpcServer(config)

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
	page := int32(1)
	var typeIDs []int32
	params := esi.GetUniverseTypesOpts{
		Page: optional.NewInt32(page),
	}

	typeResult, _, err := esiClient.ESI.UniverseApi.GetUniverseTypes(nil, &params)
	if err != nil {
		return nil, err
	}

	typeIDs = append(typeIDs, typeResult...)

	for len(typeResult) > 0 {
		page++
		params := esi.GetUniverseTypesOpts{
			Page: optional.NewInt32(page),
		}
		typeResult, _, err = esiClient.ESI.UniverseApi.GetUniverseTypes(nil, &params)
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
	typeInfo, _, err := esiClient.ESI.UniverseApi.GetUniverseTypesTypeId(nil, typeID, nil)
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
				RegionID: int64(regionID),
				TypeID:   int64(typeID),
			})
		}
	}

	return regionTypes
}

// Start the webserver.
func startGrpcServer(config Config) {
	var opts []grpc.ServerOption
	var logOpts []grpc_logrus.Option
	opts = append(opts, grpc_middleware.WithUnaryServerChain(
		grpc_ctxtags.UnaryServerInterceptor(),
		grpc_logrus.UnaryServerInterceptor(logrus.NewEntry(logrus.New()), logOpts...)))

	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", config.Port))

	if err != nil {
		log.Fatalf("could not listen: %v", err)
	}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterMarketStatsServer(grpcServer, &marketStatsServer{})
	grpcServer.Serve(listener)
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
	var bulkBacklog = make(chan *pb.HistoryStats, 5000)

	for i := 0; i < esiWorkers; i++ {
		go esiWorker(backlog, esiDone, esiTerminate, bulkBacklog)
	}

	// Create bulk packer done, termination and data forward channels
	const bulkSize = 500
	var bulkItemsLeft = len(regionTypes)
	var bulkDone = make(chan int, 100)
	var bulkTerminate = make(chan struct{})
	var dbBacklog = make(chan []*pb.HistoryStats, 50)

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
func esiWorker(backlog <-chan *types.RegionType, done chan<- error, terminate <-chan struct{}, bulkBacklog chan<- *pb.HistoryStats) {
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
		history, _, err := esiClient.ESI.MarketApi.GetMarketsRegionIdHistory(nil, int32(regionType.RegionID), int32(regionType.TypeID), nil)

		if err == nil {
			return history, nil
		}

		// Add slight randomized backoff e.g. if we're blocked
		logrus.WithError(err).Warn("could not download stats: ")
		sleepTime := (60000 / retries) + (rand.Int() % 3000)
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	}

	return nil, err
}

// Calculate stats
func calculateStats(regionType *types.RegionType, history []esi.GetMarketsRegionIdHistory200Ok) (*pb.HistoryStats, error) {
	retval := pb.HistoryStats{}
	retval.RegionId = regionType.RegionID
	retval.TypeId = regionType.TypeID
	retval.GeneratedAt = ptypes.TimestampNow()

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

	currentStatsDateProto, err := ptypes.TimestampProto(currentStatsDate)
	if err != nil {
		return nil, err
	}

	retval.CurrentStats = &pb.HistoryDataPoint{
		Date:       currentStatsDateProto,
		Average:    currentStats.Average,
		Highest:    currentStats.Highest,
		Lowest:     currentStats.Lowest,
		Volume:     currentStats.Volume,
		OrderCount: currentStats.OrderCount,
	}

	// Get previous (most of the time this should be yesterday) point
	yesterdaysStats := datapoints[len(datapoints)-2]

	yesterdaysStatsDate, err := time.Parse("2006-01-02", yesterdaysStats.Date)
	if err != nil {
		return nil, err
	}

	yesterdaysStatsDateProto, err := ptypes.TimestampProto(yesterdaysStatsDate)
	if err != nil {
		return nil, err
	}

	retval.PreviousStats = &pb.HistoryDataPoint{
		Date:       yesterdaysStatsDateProto,
		Average:    yesterdaysStats.Average,
		Highest:    yesterdaysStats.Highest,
		Lowest:     yesterdaysStats.Lowest,
		Volume:     yesterdaysStats.Volume,
		OrderCount: yesterdaysStats.OrderCount,
	}

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

	retval.WeekIskVolumeAverage = sumISK / float64(len(iskVolumeSeries))
	retval.WeekIskVolumeAverageStandardDeviation = stdDeviationISK
	retval.WeekIskVolumeAverageRelativeStandardDeviation = stdDeviationISK / retval.WeekIskVolumeAverage

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
func packRegionStats(size int, backlog <-chan *pb.HistoryStats, done chan<- int, terminate <-chan struct{}, dbBacklog chan<- []*pb.HistoryStats) {
	for {
		var packet []*pb.HistoryStats
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
func dbWorker(backlog <-chan []*pb.HistoryStats, done chan<- int, terminate <-chan struct{}) {
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
func buildQuery(regionStats []*pb.HistoryStats) (string, []interface{}, error) {
	placeholders := make([]string, 0, len(regionStats))
	values := make([]interface{}, 0, len(regionStats)*30)
	counter := 1

	for _, stat := range regionStats {
		var positionals []string

		for i := 0; i < 30; i++ {
			positionals = append(positionals, fmt.Sprintf("$%d", counter+i))
		}

		counter = counter + 30

		// Convert between protobuf typestamps and standard timestamps
		generatedAt, err := ptypes.Timestamp(stat.GeneratedAt)
		if err != nil {
			return "", nil, err
		}

		currentDate, err := ptypes.Timestamp(stat.CurrentStats.Date)
		if err != nil {
			return "", nil, err
		}

		previousDate, err := ptypes.Timestamp(stat.PreviousStats.Date)
		if err != nil {
			return "", nil, err
		}

		newPlaceholders := fmt.Sprintf("(%s)", strings.Join(positionals, ","))
		placeholders = append(placeholders, newPlaceholders)
		values = append(values, stat.RegionId)
		values = append(values, stat.TypeId)
		values = append(values, generatedAt)
		values = append(values, currentDate)
		values = append(values, stat.CurrentStats.Highest)
		values = append(values, stat.CurrentStats.Lowest)
		values = append(values, stat.CurrentStats.Average)
		values = append(values, stat.CurrentStats.Volume)
		values = append(values, stat.CurrentStats.OrderCount)
		values = append(values, previousDate)
		values = append(values, stat.PreviousStats.Highest)
		values = append(values, stat.PreviousStats.Lowest)
		values = append(values, stat.PreviousStats.Average)
		values = append(values, stat.PreviousStats.Volume)
		values = append(values, stat.PreviousStats.OrderCount)
		values = append(values, stat.WeekPriceWeightedAverage)
		values = append(values, stat.WeekPriceAverage)
		values = append(values, stat.WeekPriceAverageStandardDeviation)
		values = append(values, stat.WeekPriceAverageRelativeStandardDeviation)
		values = append(values, stat.WeekIskVolumeAverage)
		values = append(values, stat.WeekIskVolumeAverageStandardDeviation)
		values = append(values, stat.WeekIskVolumeAverageRelativeStandardDeviation)
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

	return query, values, nil
}

// Upsert the stats
func upsertStats(regionStats []*pb.HistoryStats) error {
	query, values, err := buildQuery(regionStats)
	if err != nil {
		return err
	}

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

func (s *marketStatsServer) GetRegionHistory(context context.Context, request *pb.GetRegionRequest) (*pb.GetHistoryStatsResponse, error) {
	regionID := request.GetRegionId()
	query := `SELECT * FROM "stats" WHERE "region_id" = $1`

	rows, err := db.Query(query, regionID)
	defer rows.Close()
	if err != nil {
		logrus.WithError(err).Warn("database query failed: ")
		return nil, status.Error(codes.Internal, "database query failed")
	}

	markets, err := convertRowsToStructs(rows)
	if err != nil {
		logrus.WithError(err).Warn("parsing DB response failed: ")
		return nil, status.Error(codes.Internal, "Parsing DB response failed")
	}

	if len(markets) == 0 {
		return nil, status.Error(codes.NotFound, "No matching stats found")
	}

	return &pb.GetHistoryStatsResponse{HistoryStats: markets}, nil
}

func (s *marketStatsServer) GetTypeHistory(context context.Context, request *pb.GetTypeRequest) (*pb.GetHistoryStatsResponse, error) {
	typeID := request.GetTypeId()
	query := `SELECT * FROM "stats" WHERE "type_id" = $1`

	rows, err := db.Query(query, typeID)
	defer rows.Close()
	if err != nil {
		logrus.WithError(err).Warn("database query failed: ")
		return nil, status.Error(codes.Internal, "database query failed")
	}

	markets, err := convertRowsToStructs(rows)
	if err != nil {
		logrus.WithError(err).Warn("parsing DB response failed: ")
		return nil, status.Error(codes.Internal, "Parsing DB response failed")
	}

	if len(markets) == 0 {
		return nil, status.Error(codes.NotFound, "No matching stats found")
	}

	return &pb.GetHistoryStatsResponse{HistoryStats: markets}, nil
}

func (s *marketStatsServer) GetRegionTypeHistory(context context.Context, request *pb.GetRegionTypeRequest) (*pb.HistoryStats, error) {
	regionID := request.GetRegionId()
	typeID := request.GetTypeId()
	query := `SELECT * FROM "stats" WHERE "region_id" = $1 AND "type_id" = $2`

	rows, err := db.Query(query, regionID, typeID)
	defer rows.Close()
	if err != nil {
		logrus.WithError(err).Warn("database query failed: ")
		return nil, status.Error(codes.Internal, "database query failed")
	}

	markets, err := convertRowsToStructs(rows)
	if err != nil {
		logrus.WithError(err).Warn("parsing DB response failed: ")
		return nil, status.Error(codes.Internal, "Parsing DB response failed")
	}

	if len(markets) == 0 {
		return nil, status.Error(codes.NotFound, "No matching stats found")
	}

	return markets[0], nil
}

// Load configuration from environment
func loadConfig() Config {
	config := Config{}
	envconfig.MustProcess("MARKET_STATS", &config)

	logLevel, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		panic(err)
	}

	logrus.SetLevel(logLevel)
	logrus.Debugf("Config: %+v", config)
	return config
}

func convertRowsToStructs(rows *sql.Rows) ([]*pb.HistoryStats, error) {
	defer rows.Close()
	markets := []*pb.HistoryStats{}

	for rows.Next() {
		var regionID int64
		var typeID int64
		var generatedAt time.Time
		var date time.Time
		var highest float64
		var lowest float64
		var average float64
		var volume int64
		var orderCount int64
		var previousDate time.Time
		var previousHighest float64
		var previousLowest float64
		var previousAverage float64
		var previousVolume int64
		var previousOrderCount int64
		var weekPriceWeightedAverage float64
		var weekPriceAverage float64
		var weekPriceAverageStandardDeviation float64
		var weekPriceAverageRelativeStandardDeviation float64
		var weekIskVolumeAverage float64
		var weekIskVolumeAverageStandardDeviation float64
		var weekIskVolumeAverageRelativeStandardDeviation float64
		var weekOrderCountTotal int64
		var weekOrderCountAverage float64
		var weekOrderCountStandardDeviation float64
		var weekOrderCountRelativeStandardDeviation float64
		var weekVolumeTotal int64
		var weekVolumeAverage float64
		var weekVolumeStandardDeviation float64
		var weekVolumeRelativeStandardDeviation float64

		err := rows.Scan(&regionID,
			&typeID,
			&generatedAt,
			&date,
			&highest,
			&lowest,
			&average,
			&volume,
			&orderCount,
			&previousDate,
			&previousHighest,
			&previousLowest,
			&previousAverage,
			&previousVolume,
			&previousOrderCount,
			&weekPriceWeightedAverage,
			&weekPriceAverage,
			&weekPriceAverageStandardDeviation,
			&weekPriceAverageRelativeStandardDeviation,
			&weekIskVolumeAverage,
			&weekIskVolumeAverageStandardDeviation,
			&weekIskVolumeAverageRelativeStandardDeviation,
			&weekOrderCountTotal,
			&weekOrderCountAverage,
			&weekOrderCountStandardDeviation,
			&weekOrderCountRelativeStandardDeviation,
			&weekVolumeTotal,
			&weekVolumeAverage,
			&weekVolumeStandardDeviation,
			&weekVolumeRelativeStandardDeviation)
		if err != nil {
			return nil, err
		}

		generatedAtProto, err := ptypes.TimestampProto(generatedAt)
		if err != nil {
			return nil, err
		}

		previousDateProto, err := ptypes.TimestampProto(date)
		if err != nil {
			return nil, err
		}

		currentDateProto, err := ptypes.TimestampProto(previousDate)
		if err != nil {
			return nil, err
		}

		market := pb.HistoryStats{
			RegionId:    regionID,
			TypeId:      typeID,
			GeneratedAt: generatedAtProto,
			CurrentStats: &pb.HistoryDataPoint{
				Date:    currentDateProto,
				Highest: highest,
				Lowest:  lowest,
				Average: average,
				Volume:  volume,
			},
			PreviousStats: &pb.HistoryDataPoint{
				Date:    previousDateProto,
				Highest: previousHighest,
				Lowest:  previousLowest,
				Average: previousAverage,
				Volume:  previousVolume,
			},
			WeekPriceWeightedAverage:                      weekPriceWeightedAverage,
			WeekPriceAverage:                              weekPriceAverage,
			WeekPriceAverageStandardDeviation:             weekPriceAverageStandardDeviation,
			WeekPriceAverageRelativeStandardDeviation:     weekPriceAverageRelativeStandardDeviation,
			WeekIskVolumeAverage:                          weekIskVolumeAverage,
			WeekIskVolumeAverageStandardDeviation:         weekIskVolumeAverageStandardDeviation,
			WeekIskVolumeAverageRelativeStandardDeviation: weekIskVolumeAverageRelativeStandardDeviation,
			WeekOrderCountTotal:                           weekOrderCountTotal,
			WeekOrderCountAverage:                         weekOrderCountAverage,
			WeekOrderCountStandardDeviation:               weekOrderCountStandardDeviation,
			WeekOrderCountRelativeStandardDeviation:       weekOrderCountRelativeStandardDeviation,
			WeekVolumeTotal:                               weekVolumeTotal,
			WeekVolumeAverage:                             weekVolumeAverage,
			WeekVolumeStandardDeviation:                   weekVolumeStandardDeviation,
			WeekVolumeRelativeStandardDeviation:           weekVolumeRelativeStandardDeviation,
		}

		markets = append(markets, &market)
	}

	return markets, nil
}
