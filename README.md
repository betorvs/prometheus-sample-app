# prometheus-sample-app

## testing

```bash
TEST=true ; count=1 ; while $TEST; do echo $count ; curl -I localhost:8080/ping; sleep 0.1  ; count=$((count+1)) ; if [[ $count = "11" ]]; then export TEST=false ; fi ; done
```


default
```
http_response_time_seconds_bucket{path="/ping",le="0.005"} 0
http_response_time_seconds_bucket{path="/ping",le="0.01"} 0
http_response_time_seconds_bucket{path="/ping",le="0.025"} 0
http_response_time_seconds_bucket{path="/ping",le="0.05"} 0
http_response_time_seconds_bucket{path="/ping",le="0.1"} 0
http_response_time_seconds_bucket{path="/ping",le="0.25"} 4
http_response_time_seconds_bucket{path="/ping",le="0.5"} 4
http_response_time_seconds_bucket{path="/ping",le="1"} 10
http_response_time_seconds_bucket{path="/ping",le="2.5"} 10
http_response_time_seconds_bucket{path="/ping",le="5"} 10
http_response_time_seconds_bucket{path="/ping",le="10"} 10
http_response_time_seconds_bucket{path="/ping",le="+Inf"} 10
http_response_time_seconds_sum{path="/ping"} 4.609575167
http_response_time_seconds_count{path="/ping"} 10
...
response_status{status="200"} 10
```

FEATURE_FLAG=instable
```
http_response_time_seconds_bucket{path="/ping",le="0.005"} 0
http_response_time_seconds_bucket{path="/ping",le="0.01"} 0
http_response_time_seconds_bucket{path="/ping",le="0.025"} 0
http_response_time_seconds_bucket{path="/ping",le="0.05"} 0
http_response_time_seconds_bucket{path="/ping",le="0.1"} 0
http_response_time_seconds_bucket{path="/ping",le="0.25"} 0
http_response_time_seconds_bucket{path="/ping",le="0.5"} 0
http_response_time_seconds_bucket{path="/ping",le="1"} 0
http_response_time_seconds_bucket{path="/ping",le="2.5"} 4
http_response_time_seconds_bucket{path="/ping",le="5"} 4
http_response_time_seconds_bucket{path="/ping",le="10"} 10
http_response_time_seconds_bucket{path="/ping",le="+Inf"} 10
http_response_time_seconds_sum{path="/ping"} 46.01143983199999
http_response_time_seconds_count{path="/ping"} 10
...
response_status{status="200"} 6
response_status{status="500"} 4
```

FEATURE_FLAG=broken
```
http_response_time_seconds_bucket{path="/ping",le="0.005"} 0
http_response_time_seconds_bucket{path="/ping",le="0.01"} 0
http_response_time_seconds_bucket{path="/ping",le="0.025"} 4
http_response_time_seconds_bucket{path="/ping",le="0.05"} 4
http_response_time_seconds_bucket{path="/ping",le="0.1"} 10
http_response_time_seconds_bucket{path="/ping",le="0.25"} 10
http_response_time_seconds_bucket{path="/ping",le="0.5"} 10
http_response_time_seconds_bucket{path="/ping",le="1"} 10
http_response_time_seconds_bucket{path="/ping",le="2.5"} 10
http_response_time_seconds_bucket{path="/ping",le="5"} 10
http_response_time_seconds_bucket{path="/ping",le="10"} 10
http_response_time_seconds_bucket{path="/ping",le="+Inf"} 10
http_response_time_seconds_sum{path="/ping"} 0.470832958
http_response_time_seconds_count{path="/ping"} 10
...
response_status{status="200"} 1
response_status{status="500"} 9
```

FEATURE_FLAG=quick
```
http_response_time_seconds_bucket{path="/ping",le="0.005"} 2
http_response_time_seconds_bucket{path="/ping",le="0.01"} 2
http_response_time_seconds_bucket{path="/ping",le="0.025"} 4
http_response_time_seconds_bucket{path="/ping",le="0.05"} 4
http_response_time_seconds_bucket{path="/ping",le="0.1"} 10
http_response_time_seconds_bucket{path="/ping",le="0.25"} 10
http_response_time_seconds_bucket{path="/ping",le="0.5"} 10
http_response_time_seconds_bucket{path="/ping",le="1"} 10
http_response_time_seconds_bucket{path="/ping",le="2.5"} 10
http_response_time_seconds_bucket{path="/ping",le="5"} 10
http_response_time_seconds_bucket{path="/ping",le="10"} 10
http_response_time_seconds_bucket{path="/ping",le="+Inf"} 10
http_response_time_seconds_sum{path="/ping"} 0.447974331
http_response_time_seconds_count{path="/ping"} 10
...
response_status{status="200"} 9
response_status{status="500"} 1
```