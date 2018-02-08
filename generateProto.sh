#!/bin/bash
mkdir -p ./lib/marketStats

go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/golang/protobuf/protoc-gen-go

protoc -I../element43/services/marketStats \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--go_out=plugins=grpc:./lib/marketStats \
../element43/services/marketStats/marketStats.proto