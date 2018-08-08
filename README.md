# Market Stats
[![Build Status](https://semaphoreci.com/api/v1/zweizeichen/market-stats/branches/master/badge.svg)](https://semaphoreci.com/zweizeichen/market-stats) [![Go Report Card](https://goreportcard.com/badge/github.com/EVE-Tools/market-stats)](https://goreportcard.com/report/github.com/EVE-Tools/market-stats) [![Docker Image](https://images.microbadger.com/badges/image/evetools/market-stats.svg)](https://microbadger.com/images/evetools/market-stats)

This service for [Element43](https://element-43.com) calculates various market metrics based on historical data. Data is stored in Postgres and can be retrieved via a gRPC API. Every night it fetches the entire market's history from ESI and calculates the metrics described in its [gRPC service description](https://github.com/EVE-Tools/element43/services/marketStats/marketStats.proto). There might be gaps in the data as there is no history available if nothing was traded.

Issues can be filed [here](https://github.com/EVE-Tools/element43). Pull requests can be made in this repo.

## Interface
The service's gRPC description can be found [here](https://github.com/EVE-Tools/element43/blob/master/services/marketStats/marketStats.proto.

## Installation
Either use the prebuilt Docker images and pass the appropriate env vars (see below), or:

* Install Go, clone this repo into your gopath
* Run `go get ./...` to fetch the service's dependencies
* Run `bash generateProto.sh` to generate the necessary gRPC-related code
* Run `go build`
* Run `./market-stats` to start the service

## Deployment Info
Builds and releases are handled by Drone.

Environment Variable | Default | Description
--- | --- | ---
SEED_DB | false | Fetch all data from ESI on startup, ignoring the normal schedule
CRON | 0 52 * * * | Defines when to run data collection (0:52am UTC every day avoids getting cached data)
LOG_LEVEL | info | Threshold for logging messages to be printed
POSTGRES_URL| postgres://market-stats@localhost:5432/market-stats?sslmode=disable | URL to this service's Postgres database
PORT | 43000 | Port for the grpc server to listen on