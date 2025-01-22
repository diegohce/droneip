# Man in the Middle for dronebl service

Checks for ip address ban in dronebl.org and forwards request to destination.

## Build

```bash
go generate ./... && go build -o drone-ip .
```

## Config env vars (and default values)

```bash
# empty: no cache
# memory:// (in memory cache)  
# redis://:master@127.0.0.1/7 (redis as cache)
DRONEIP_CACHE_URL="memory://"

# network_interface:port to bind to. Default 0.0.0.0:8080
DRONEIP_BIND=":8080"

# network_interface:port to bind to. Default 0.0.0.0:8081
DRONEIP_ADMIN_BIND=":8081"

# HTTP header to inspect for IP address. If empty will use
#remote tcp address.
DRONEIP_INSPECT_HEADER=""

# Cache TTL in seconds. Default 24h.
DRONEIP_CACHE_TTL=""

# Mandatory value in the form: http://host[:port]/path/to/destination
DRONEIP_DESTINATION_URL=""

## Admin port requests

### Cached ips

`GET /droneip/keys`

Response:

```json
[
    "droneip-10.10.0.1",
    "droneip-10.10.0.2",
    "droneip-10.10.0.3",
    "droneip-10.10.0.4",
    "droneip-10.10.0.5",
    "droneip-10.10.0.6",
]
```


### Version
`GET /droneip/version`

Response:

```json
{
    "commit": "28d0aa38f63f558b1bbc7b8b1bfe3dc01b757a50",
    "date": "Tue Nov 15 16:10:36 2022 -0300",
    "version": "v0.1.0"
}
```



## Testing

### Run tests:

```bash
you@hal9000:~$ go generate ./...
you@hal9000:~$ go test -tags test -count=1 -coverprofile=coverage.out .
```

### Check code coverage

```bash
you@hal9000:~$ go tool cover -func=coverage.out
```

or

```bash
you@hal9000:~$ go tool cover -html=coverage.out
```

## Docker image

```bash
you@hal9000:~$ docker build -t mim-drone-ip:latest .
```
## Sequence Diagram

TODO


