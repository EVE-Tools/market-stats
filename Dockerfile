FROM alpine:latest

MAINTAINER zweizeichen@element-43.com

#
# Copy release to container and set command
#

# Copy build
COPY migrations migrations
COPY market-stats market-stats

ENV PORT 80
EXPOSE 80

CMD ["/market-stats"]