FROM golang:1.23.3 as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build
RUN go generate ./... && go build -o main .



FROM ubuntu:20.04

RUN apt update && apt install -y ca-certificates

# empty: no cache
# memory:// (in memory cache)  
# redis://:master@127.0.0.1/7 (redis as cache)
ENV DRONEIP_CACHE_URL="memory://"

# network_interface:port to bind to. Default 0.0.0.0:8080
ENV DRONEIP_BIND=":8080"

# network_interface:port to bind to. Default 0.0.0.0:8081
ENV DRONEIP_ADMIN_BIND=":8081"

# HTTP header to inspect for IP address. If empty will use
#remote tcp address.
ENV DRONEIP_INSPECT_HEADER=""

# Cache TTL in seconds. Default 24h.
ENV DRONEIP_CACHE_TTL=""

# Mandatory value in the form: http://host[:port]/path/to/destination
ENV DRONEIP_DESTINATION_URL=""


RUN mkdir /app
COPY --from=builder /build/main /app/

WORKDIR /app

CMD ["./main"]
