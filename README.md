# Market Stats
[![Build Status](https://drone.element-43.com/api/badges/EVE-Tools/market-stats/status.svg)](https://drone.element-43.com/EVE-Tools/market-stats) [![Go Report Card](https://goreportcard.com/badge/github.com/EVE-Tools/market-stats)](https://goreportcard.com/report/github.com/EVE-Tools/market-stats) [![Docker Image](https://images.microbadger.com/badges/image/evetools/market-stats.svg)](https://microbadger.com/images/evetools/market-stats)

This service for [Element43](https://element-43.com) calculates various market metrics based on historical data. Every night it fetches the entire market history from ESI and calculates the following metrics:

Value | Description
--- | ---
region_id | The result's region's ID
type_id | The result's type's ID
generated_at | Date this dataset was fetched from the server (usually at midnight)
date | Date of the latest datapoint available (usually yesterday, see below)
highest | Highest price the type was traded for that day
lowest | Lowest price the type was traded for that day
average | Average price the type was traded for that day
volume | Total number of items of the type traded that day
order_count | Total number orders for the type that day
previous_date | Date of the second newest datapoint (usually the day before yesterday, see below)
previous_highest | Highest price the type was traded for that day
previous_lowest | Lowest price the type was traded for that day
previous_average | Average price the type was traded for that day
previous_volume | Total number of items of the type traded that day
previous_order_count | Total number orders for the type that day
week_price_weighted_average | Last week's *weighted* average of the average price - use this for calculations!
week_price_average | Last week's average average price (ignores volume)
week_price_average_standard_deviation | Last week's absolute standard deviation of the average price
week_price_average_relative_standard_deviation | Last week's absolute standard deviation of the average price, can be used for analyzing price's volatility
week_isk_volume_average | Last week's average daily ISK volume
week_isk_volume_average_standard_deviation | Last week's average daily ISK volume's standard deviation
week_isk_volume_average_relative_standard_deviation | Last week's relative daily ISK volume's standard deviation, can be used for analyzing general market's volatility
week_order_count_total | Total number of orders for that type last week
week_order_count_average | Average daily order count last week
week_order_count_standard_deviation | Average daily order count's standard deviation last week
week_order_count_relative_standard_deviation | Average daily order count's relative standard deviation last week, again can be used for volatility
week_volume_total | Total number of items of that type traded last week
week_volume_average | Average daily items traded of that type last week
week_volume_standard_deviation | Standard deviation of that value
week_volume_relative_standard_deviation | Relative standard deviation, interesting for checking if there are non-continuous patterns (volatility)

Data is stored in Postgres and can be retrieved via an API (see below). There might be gaps in the data as there is no history available if nothing was traded.

## Installation
Either use the prebuilt Docker images and pass the appropriate env vars (see below), or:

* Clone this repo into your gopath
* Run `go get`
* Run `go build`

## Deployment Info
Builds and releases are handled by Drone.

Environment Variable | Default | Description
--- | --- | ---
SEED_DB | false | Fetch all data from ESI on startup, ignoring the normal schedule
CRON | 0 5 1 * * * | Defines when to run data collection (1:05am UTC every day avoids getting cached data)
LOG_LEVEL | info | Threshold for logging messages to be printed
POSTGRES_URL| postgres://market-stats@localhost:5432/market-stats?sslmode=disable | URL to this service's Postgres database
PORT | 8000 | Port for the API server to listen on

## Todo
- [ ] Some tests to prove validity of calculations would be nice
- [ ] Move HTTP server and API into own modules
- [ ] Investigate update pipeline performance, it does a couple of hundred requests per second, yet it almost uses no system resources
- [ ] Investigate caching (we might not even need it right now, it is quite fast)

## Endpoints

Prefix: `/api/market-stats/v1/history`

URL Pattern | Description
--- | ---
`/region/:regionID/` | Get all stats in a region - this may return a lot of data
`/type/:typeID/` | Get all stats for a type
`/region/:regionID/type/:typeID/` | Get all stats of a type in a region

The return values share a similar schema. Truncated example output for `api/market-stats/v1/history/region/10000002/type/34/`:
```json
{
    "region_id": 10000002,
    "type_id": 34,
    "generated_at": "2017-03-20",
    "date": "2017-03-19",
    "highest": 4.3299999,
    "lowest": 4.1799998,
    "average": 4.3200002,
    "volume": 14884892223,
    "order_count": 2976,
    "previous_date": "2017-03-18",
    "previous_highest": 4.4299998,
    "previous_lowest": 4.29,
    "previous_average": 4.3800001,
    "previous_volume": 24129675845,
    "previous_order_count": 2717,
    "week_price_weighted_average": 4.317331841859094,
    "week_price_average": 4.341428552355085,
    "week_price_average_standard_deviation": 0.12888846938753978,
    "week_price_average_relative_standard_deviation": 0.029688031907751183,
    "week_isk_volume_average": 75411122089.5082,
    "week_isk_volume_average_standard_deviation": 22913268515.166306,
    "week_isk_volume_average_relative_standard_deviation": 0.30384468338728227,
    "week_order_count_total": 17529,
    "week_order_count_average": 2504.1428571428573,
    "week_order_count_standard_deviation": 234.0436031240019,
    "week_order_count_relative_standard_deviation": 0.013351794347880763,
    "week_volume_total": 122269465022,
    "week_volume_average": 17467066431.714287,
    "week_volume_standard_deviation": 5615453029.500905,
    "week_volume_relative_standard_deviation": 0.04592686349360009
}
```