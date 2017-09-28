FROM alpine:3.6

#
# Copy release to container and set command
#

# Add faster mirror and upgrade packages in base image, load ca certs, otherwise no TLS for us
RUN printf "http://mirror.leaseweb.com/alpine/v3.6/main\nhttp://mirror.leaseweb.com/alpine/v3.6/community" > etc/apk/repositories && \
    apk update && \
    apk upgrade && \
    apk add ca-certificates && \
    rm -rf /var/cache/apk/*

# Copy build
COPY migrations migrations
COPY market-stats market-stats

ENV PORT 80
EXPOSE 80

CMD ["/market-stats"]