#
# Build project in separate container
#

FROM golang:alpine3.8 AS build

RUN apk update && \
    apk upgrade && \
    apk add git
COPY . /go/src/github.com/EVE-Tools/market-stats

WORKDIR /go/src/github.com/EVE-Tools/market-stats
RUN go get -d -v ./...
RUN go build
RUN cp /go/src/github.com/EVE-Tools/market-stats/market-stats /market-stats

#
# Copy release to fresh container and set command
#

FROM alpine:3.8

# Update base system, load ca certs, otherwise no TLS for us
RUN apk update && \
    apk upgrade && \
    apk add ca-certificates && \
    rm -rf /var/cache/apk/*

# Do not run as root
RUN addgroup -g 1000 -S element43 && \
    adduser -u 1000 -S element43 -G element43
USER element43:element43

# Copy build
COPY --from=build /market-stats /market-stats

ENV PORT 43000
EXPOSE 43000

CMD ["/market-stats"]