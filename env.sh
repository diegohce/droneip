#!/bin/bash

# empty: no cache
# memory:// (in memory cache)  
# redis://:master@127.0.0.1/7 (redis as cache)
export DRONEIP_CACHE_URL="memory://"

# network_interface:port to bind to. Default 0.0.0.0:8080
export DRONEIP_BIND=":8080"

# network_interface:port to bind to. Default 0.0.0.0:8081
export DRONEIP_ADMIN_BIND=":8081"

# HTTP header to inspect for IP address. If empty will use
#remote tcp address.
export DRONEIP_INSPECT_HEADER=""

# Cache TTL in seconds. Default 24h.
export DRONEIP_CACHE_TTL=""

# Mandatory value in the form: http://host[:port]/path/to/destination
export DRONEIP_DESTINATION_URL=""

