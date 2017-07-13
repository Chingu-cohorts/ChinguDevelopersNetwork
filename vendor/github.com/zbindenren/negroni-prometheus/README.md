# negroni-prometheus [![GoDoc](http://godoc.org/github.com/zbinderen/negroni-prometheus?status.svg)](http://godoc.org/github.com/zbindenren/negroni-prometheus) [![Go Report Card](https://goreportcard.com/badge/github.com/zbindenren/negroni-prometheus)](https://goreportcard.com/report/github.com/zbindenren/negroni-prometheus)
[Prometheus](http://prometheus.io) middleware for [negroni](https://github.com/codegangsta/negroni).

## Why
[Logging v. instrumentation](http://peter.bourgon.org/blog/2016/02/07/logging-v-instrumentation.html)

Instead of logging request times, it is considered best practice to provide an endpoint for instrumentation tools (like prometheus).

## Installation

```bash
$ go get github.com/zbindenren/negroni-prometheus
```

## Usage

Use this middleware like the `negroni.Logger` middleware (after `negroni.Recovery`, before every other middleware).

Take a look at the [example](./example/main.go).

## What do you get

An endpoint with the following information (stripped output):
```
...
# HELP negroni_request_duration_milliseconds How long it took to process the request, partitioned by status code, method and HTTP path.
# TYPE negroni_request_duration_milliseconds histogram
negroni_request_duration_milliseconds_bucket{code="OK",method="GET",path="/metrics",service="serviceName",le="300"} 1
negroni_request_duration_milliseconds_bucket{code="OK",method="GET",path="/metrics",service="serviceName",le="1200"} 1
negroni_request_duration_milliseconds_bucket{code="OK",method="GET",path="/metrics",service="serviceName",le="5000"} 1
negroni_request_duration_milliseconds_bucket{code="OK",method="GET",path="/metrics",service="serviceName",le="+Inf"} 1
negroni_request_duration_milliseconds_sum{code="OK",method="GET",path="/metrics",service="serviceName"} 2.003123
negroni_request_duration_milliseconds_count{code="OK",method="GET",path="/metrics",service="serviceName"} 1
negroni_request_duration_milliseconds_bucket{code="OK",method="GET",path="/ok",service="serviceName",le="300"} 0
negroni_request_duration_milliseconds_bucket{code="OK",method="GET",path="/ok",service="serviceName",le="1200"} 0
negroni_request_duration_milliseconds_bucket{code="OK",method="GET",path="/ok",service="serviceName",le="5000"} 2
negroni_request_duration_milliseconds_bucket{code="OK",method="GET",path="/ok",service="serviceName",le="+Inf"} 2
negroni_request_duration_milliseconds_sum{code="OK",method="GET",path="/ok",service="serviceName"} 4747.529026
negroni_request_duration_milliseconds_count{code="OK",method="GET",path="/ok",service="serviceName"} 2
# HELP negroni_requests_total How many HTTP requests processed, partitioned by status code, method and HTTP path.
# TYPE negroni_requests_total counter
negroni_requests_total{code="OK",method="GET",path="/metrics",service="serviceName"} 1
negroni_requests_total{code="OK",method="GET",path="/ok",service="serviceName"} 2
...
```
