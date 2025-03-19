# GO Actuator

[![GoDoc](https://godoc.org/github.com/sinhashubham95/go-actuator?status.svg)](https://pkg.go.dev/github.com/sinhashubham95/go-actuator)
[![Release](https://img.shields.io/github/v/release/sinhashubham95/go-actuator?sort=semver)](https://github.com/sinhashubham95/go-actuator/releases)
[![Report](https://goreportcard.com/badge/github.com/sinhashubham95/go-actuator)](https://goreportcard.com/report/github.com/sinhashubham95/go-actuator)
[![Coverage Status](https://coveralls.io/repos/github/sinhashubham95/go-actuator/badge.svg?branch=master)](https://coveralls.io/github/sinhashubham95/go-actuator?branch=master)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go#utilities)

GO actuator configures the set of actuator endpoints for your application. It is very extensible and can be configured with `Go's native HTTP Server Mux`, or with any `3rd party web framework` as well.

## Project Versioning

Go actuator uses [semantic versioning](http://semver.org/). API should not change between patch and minor releases. New minor versions may add additional features to the API.

## Installation

To install `Go Actuator` package, you need to install Go and set your Go workspace first.

1. The first need Go installed (version 1.18+ is required), then you can use the below Go command to install Go Actuator.

```shell
go get github.com/sinhashubham95/go-actuator
```

2. Import it in your code:

```go
import "github.com/sinhashubham95/go-actuator"
```

## How to Use

The actuator library exposes a plain native handler function, and it is the responsibility of the application to put this handler to use. This can be used either directly with `Go's native HTTP Server Mux`, or with any `3rd party web framework` as well.

### Configuration

The configuration contains the following:-

1. **Endpoints** - This is the list of endpoints which will be enabled. This is not a mandatory parameter. If not provided, then only `info` and `ping` endpoints will be enabled. The possible endpoints are - `/env`, `/info`, `/health`, `/metrics`, `/ping`, `/shutdown` and `/threadDump`. You can find the description of each of these endpoints below.
2. **Env** - This is the environment where the application is running. For example, `dev`, `stg`, `prod`, etc.
3. **Name** - This is the name of the application which is using this actuator library.
4. **Port** - This is the port where the application is running.
5. **Version** - This is the current application version.
6. **Health** - This is used to provide a set of health checkers to derive the health of an application using its dependencies, some of which might be mandatory and some non-mandatory, which helps maintain a robust eye on the application.

```go
import actuator "github.com/sinhashubham95/go-actuator"

config := &actuator.Config{
	Endpoints: []int{
		actuator.Env,
		actuator.Info,
		actuator.Metrics,
		actuator.Ping,
		actuator.Shutdown,
		actuator.ThreadDump,
    },
    Env: "dev",
    Name: "Naruto Rocks",
    Port: 8080,
    Version: "0.1.0",
}
```

### Using with [Go's Native Server Mux](https://pkg.go.dev/net/http)

```go
import (
    actuator "github.com/sinhashubham95/go-actuator"
    "net/http"
)

// create a server
mux := &http.ServeMux{}

// get the handler for actuator
actuatorHandler := actuator.GetActuatorHandler(&actuator.Config{})
// configure the handler at this base endpoint
mux.Handle("/actuator", actuatorHandler)

// configure other handlers
....
```

### Using with [Fast HTTP](https://github.com/valyala/fasthttp)

```go
import (
	"strings"
	
    "github.com/valyala/fasthttp"
    actuator "github.com/sinhashubham95/go-actuator"
)

// get the handler for actuator
actuatorHandler := fasthttp.NewFastHTTPHandlerFunc(actuator.GetActuatorHandler(&actuator.Config{}))

// create a fast http handler
handler := func(ctx *fasthttp.RequestCtx) {
    if strings.HasPrefix(ctx.Path(), "/actuator") {
        // use the actuator handler
    	actuatorHandler(ctx)
    	return
    }
    // other request handler calls
}
fasthttp.ListenAndServe(":8080", handler)
```

### Using with [GIN](https://github.com/gin-gonic/gin)

```go
import (
    "github.com/gin-gonic/gin"
    actuator "github.com/sinhashubham95/go-actuator"
	"github.com/sinhashubham95/go-actuator/models"
)

// create the gin engine
engine := gin.Default()

// get the handler for actuator
actuatorHandler := actuator.GetActuatorHandler(&actuator.Config{})
ginActuatorHandler := func(ctx *gin.Context) {
	actuatorHandler(ctx.Writer, ctx.Request)
}

engine.GET("/actuator/*endpoint", ginActuatorHandler)
```

### Using with [Hertz](https://github.com/cloudwego/hertz)

```go
import (
    "context"
    "github.com/cloudwego/hertz/pkg/app"
    "github.com/cloudwego/hertz/pkg/app/server"
    "github.com/cloudwego/hertz/pkg/common/adaptor"
    "github.com/sinhashubham95/go-actuator"
)

var actuatorHandler = actuator.GetActuatorHandler(&actuator.Config{})

func handleActuator(ctx context.Context, c *app.RequestContext) {
	request, err := adaptor.GetCompatRequest(&c.Request)
	if err != nil {
        c.Status(http.StatusInternalServerError)
		return
	}
	response := adaptor.GetCompatResponseWriter(&c.Response)
	actuatorHandler(response, request)
}

s := server.Default()
s.Any("/actuator/*any", handleActuator)
```

## Endpoints

### Env - `/actuator/env`

This is used to get all the environment variables for the runtime where the application is running. Note that to use this, you need to pass the runtime environment as an application flag as follows.

```shell
go build
./${APPLICATION_NAME}
```

```json
{
  "env_key_1": "env_value_1",
  "env_key_2": "env_value_2"
}
```

### Info - `/actuator/info`

This is used to get the basic information for an application. To get the correct and relevant information for your application you need to change the build script as well as the run script for your application as follows.

```shell
buildStamp=$(date -u '+%Y-%m-%d_%I:%M:%S%p')
commitId=$(git rev-list -1 HEAD)
commitTime=$(git show -s --format=%ci "$commitId")
commitAuthor=$(git --no-pager show -s --format='%an <%ae>' "$commitId")
gitUrl=$(git config --get remote.origin.url)
userName=$(whoami)
hostName=$(hostname)
go build -ldflags "<other linking params> -X github.com/sinhashubham95/go-actuator.BuildStamp=$buildStamp -X github.com/sinhashubham95/go-actuator.GitCommitID=$commitId -X github.com/sinhashubham95/go-actuator.GitPrimaryBranch=$2 -X github.com/sinhashubham95/go-actuator.GitURL=$gitUrl -X github.com/sinhashubham95/go-actuator.Username=$userName -X github.com/sinhashubham95/go-actuator.HostName=$hostName  -X \"github.com/sinhashubham95/go-actuator.GitCommitTime=$commitTime\" -X \"github.com/sinhashubham95/go-actuator.GitCommitAuthor=$commitAuthor\""
./${APPLICATION_NAME}
```

```json
{
  "application": {
    "env": "ENVIRONMENT",
    "name": "APPLICATION_NAME",
    "version": "APPLICATION_VERSION"
  },
  "git": {
    "username": "",
    "hostName": "",
    "buildStamp": "",
    "commitAuthor": "Shubham Sinha ",
    "commitId": "",
    "commitTime": "",
    "branch": "",
    "url": "",
    "startupStamp": ""
  },
  "runtime": {
    "arch": "",
    "os": "",
    "port": 8080,
    "runtimeVersion": ""
  }
}
```

### Health - `/actuator/health`

This is used to provide the most health of an application, validating the underlying dependencies, via the various health checks.

These health checks might be costly to the application. That's why there is an additional cache which is implemented internally for the same. A successful health check is cached for the provided cache duration(described below). During this duration, the same cached health check will be provided as part of the response.

The health check can be configured with the following set of details, which needs to be provided only if the `Health` endpoint is enabled.
1. **CacheDuration** - This is an optional parameter, which defines the duration for which the health check once performed will be served from cache if success. The default value is 1 hour.
2. **Timeout** - This is an optional parameter, which signifies the time within which the health check should be completed. The context passed in the health check functions, will cancel post this time duration. The implementation of the health check function should honour the context cancellation.
3. **Checkers** - These are the set of dependencies of the application which needs to be validated as part of the health check. Each health checker can have the following set of configurations.
   1. **Key** - This is the unique key which can identify the dependency of this application uniquely.
   2. **Func** - This is the implementation function which will be called as part of the health check.
   3. **IsMandatory** - This tells whether this dependency is a mandatory dependency for this application or not. If this value is set to false, even if the provided health check function fails, it won't fail the overall health check.

```go
import "github.com/sinhashubham95/go-actuator"

cfg := &actuator.Config{
	Endpoints: []int{actuator.Health},
    Health: &actuator.HealthConfig{
		Checkers: []actuator.HealthChecker{
		    {
			    Key: "test1",
				Func: func(ctx context.Context) error {
					// validate connectivity to database
					// or validate connectivity to redis
					// or validate downstream api service
					// etc
                },
				IsMandatory: true,
            },
        },
    },
}
```

```json
{
  "test1": {
    "key": "test1",
    "isMandatory": false,
    "success": false,
    "error": "some error"
  },
  "test2": {
    "key": "test1",
    "isMandatory": true,
    "success": true
  }
}
```

### Metrics - `/actuator/metrics`

This is used to get the runtime memory statistics for your application. You can find the definition of each of the fields [here](./models/memStats.go).

```json
{
  "alloc": 2047816,
  "totalAlloc": 2850832,
  "sys": 73942024,
  "lookups": 0,
  "mAllocations": 15623,
  "frees": 9223,
  "heapAlloc": 2047816,
  "heapSys": 66551808,
  "heapIdle": 62832640,
  "heapInUse": 3719168,
  "heapReleased": 62570496,
  "heapObjects": 6400,
  "stackInUse": 557056,
  "stackSys": 557056,
  "mSpanInUse": 81056,
  "mSpanSys": 81920,
  "MCacheInUse": 19200,
  "mCacheSys": 32768,
  "buckHashSys": 1446250,
  "gcSys": 4225056,
  "otherSys": 1047166,
  "nextGC": 4194304,
  "lastGC": 1627102938524536000,
  "pauseTotalNs": 35655,
  "pauseNs": [
    35655
  ],
  "pauseEnd": [
    1627102938524536000
  ],
  "numGC": 1,
  "numForcedGC": 0,
  "gcCPUFraction": 0.000005360999257331059,
  "enableGC": true,
  "debugGC": false,
  "BySize": [
    {
      "Size": 0,
      "MAllocations": 0,
      "Frees": 0
    }
  ]
}
```

### Ping - `/actuator/ping`

This is the lightweight ping endpoint that can be used along with your load balancer. This is used to know the running status of your application.

### Shutdown - `/actuator/shutdown`

This is used to bring the application down.

### Thread dump - `/actuator/threadDump`

This is used to get the trace of all the goroutines.

```text
goroutine profile: total 1
1 @ 0x103af45 0x10337fb 0x10688f5 0x10c4de5 0x10c58b5 0x10c5897 0x1117e0f 0x1124391 0x11355e8 0x113576f 0x12037a5 0x1203676 0x1217025 0x1217007 0x121db9a 0x121e5b5 0x106e3e1
#	0x10688f4	internal/poll.runtime_pollWait+0x54				/Users/s0s01qp/go/go1.16.6/src/runtime/netpoll.go:222
#	0x10c4de4	internal/poll.(*pollDesc).wait+0x44				/Users/s0s01qp/go/go1.16.6/src/internal/poll/fd_poll_runtime.go:87
#	0x10c58b4	internal/poll.(*pollDesc).waitRead+0x1d4			/Users/s0s01qp/go/go1.16.6/src/internal/poll/fd_poll_runtime.go:92
#	0x10c5896	internal/poll.(*FD).Read+0x1b6					/Users/s0s01qp/go/go1.16.6/src/internal/poll/fd_unix.go:166
#	0x1117e0e	net.(*netFD).Read+0x4e						/Users/s0s01qp/go/go1.16.6/src/net/fd_posix.go:55
#	0x1124390	net.(*conn).Read+0x90						/Users/s0s01qp/go/go1.16.6/src/net/net.go:183
#	0x11355e7	bufio.(*Reader).fill+0x107					/Users/s0s01qp/go/go1.16.6/src/bufio/bufio.go:101
#	0x113576e	bufio.(*Reader).Peek+0x4e					/Users/s0s01qp/go/go1.16.6/src/bufio/bufio.go:139
#	0x12037a4	github.com/valyala/fasthttp.(*RequestHeader).tryRead+0x64	/Users/s0s01qp/go/pkg/mod/github.com/valyala/fasthttp@v1.28.0/header.go:1520
#	0x1203675	github.com/valyala/fasthttp.(*RequestHeader).readLoop+0x55	/Users/s0s01qp/go/pkg/mod/github.com/valyala/fasthttp@v1.28.0/header.go:1506
#	0x1217024	github.com/valyala/fasthttp.(*RequestHeader).Read+0x1ae4	/Users/s0s01qp/go/pkg/mod/github.com/valyala/fasthttp@v1.28.0/header.go:1497
#	0x1217006	github.com/valyala/fasthttp.(*Server).serveConn+0x1ac6		/Users/s0s01qp/go/pkg/mod/github.com/valyala/fasthttp@v1.28.0/server.go:2112
#	0x121db99	github.com/valyala/fasthttp.(*workerPool).workerFunc+0xb9	/Users/s0s01qp/go/pkg/mod/github.com/valyala/fasthttp@v1.28.0/workerpool.go:223
#	0x121e5b4	github.com/valyala/fasthttp.(*workerPool).getCh.func1+0x34	/Users/s0s01qp/go/pkg/mod/github.com/valyala/fasthttp@v1.28.0/workerpool.go:195
```
