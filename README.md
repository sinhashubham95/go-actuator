# GO Actuator

[![GoDoc](https://godoc.org/github.com/sinhashubham95/go-actuator?status.svg)](https://pkg.go.dev/github.com/sinhashubham95/go-actuator)
[![Release](https://img.shields.io/github/v/release/sinhashubham95/go-actuator?sort=semver)](https://github.com/sinhashubham95/go-actuator/releases)
[![Report](https://goreportcard.com/badge/github.com/sinhashubham95/go-actuator)](https://goreportcard.com/report/github.com/sinhashubham95/go-actuator)
[![Coverage Status](https://coveralls.io/repos/github/sinhashubham95/go-actuator/badge.svg?branch=master)](https://coveralls.io/github/sinhashubham95/go-actuator?branch=master)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go#utilities)

GO actuator configures the set of actuator endpoints for your application. It is compatible with [Fast HTTP](https://github.com/valyala/fasthttp), [GIN](https://github.com/gin-gonic/gin) and [NET/HTTP](https://pkg.go.dev/net/http).

## Installation

```shell
go get github.com/sinhashubham95/go-actuator
```

## How to Use

The actuator library is compatible with the most famous web frameworks. This is highly configurable and each endpoint can be enabled or disabled during initialization. You can also specify a prefix path for each of these configured endpoints(with default value `/actuator`).

### Configuration

The configuration contains the following:-

1. **Endpoints** - This is the list of endpoints which will be enabled. This is not a mandatory parameter. If not provided, then all the endpoints will be enabled. The possible endpoints are - `/env`, `/httpTrace`, `/info`, `/metrics`, `/ping`, `/shutdown` and `/threadDump`. You can find the description of each of these endpoints below.

2. **Prefix** - This is the prefix request path for all the configured endpoints.

```go
import "github.com/sinhashubham95/go-actuator/models"

config := &models.Config{
	Endpoints: []int{
		models.Env, models.HTTPTrace, models.Info, models.Metrics, models.Ping, models.Shutdown, models.ThreadDump
    },
    Prefix: "/actuator"
}
```

### Using with [Fast HTTP](https://github.com/valyala/fasthttp)

```go
import (
    "github.com/valyala/fasthttp"
    actuator "github.com/sinhashubham95/go-actuator"
	"github.com/sinhashubham95/go-actuator/models"
)

actuatorHandler := actuator.GetFastHTTPActuatorHandler(&models.Config{})
handler := func(ctx *fasthttp.RequestCtx) {
	switch(ctx.Path()) {
	// your configured paths
    default:
    	actuatorHandler(ctx)
    }
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

engine := gin.Default()
actuator.ConfigureGINActuatorEngine(&models.Config{}, engine)
```

### Using with [Net HTTP](https://pkg.go.dev/net/http)

```go
import (
    actuator "github.com/sinhashubham95/go-actuator"
	"github.com/sinhashubham95/go-actuator/models"
    "net/http"
)

mux := &http.ServeMux{}
actuator.ConfigureNetHTTPHandler(&models.Config{}, mux)
```

## Endpoints

### Env - `/actuator/env`

This is used to get all the environment variables for the runtime where the application is running. Note that to use this, you need to pass the runtime environment as an application flag as follows.

```shell
go build
./${APPLICATION_NAME} -env=${ENVIRONMENT_NAME}
```

```json
{
  "env_key_1": "env_value_1",
  "env_key_2": "env_value_2"
}
```

### HTTP Trace - `/actuator/httpTrace`

This is used to get the trace for the last 100 HTTP requests. Now if this has to be used, then the HTTP requests should be wrapped with the trace as follows.

```go
import (
    actuatorCore "github.com/sinhashubham95/go-actuator/core"
    "net/http"
    "strings"
)

request, err := http.NewRequest(http.MethodGet, "https://sample.com", strings.NewReader("sample"))
if err != nil {
	// incorrect request
}

request = actuatorCore.WithClientTrace(request)

// now use this request with whichever http client you want to use
response, err := http.DefaultClient.Do(request)
```

```json
[
  {
    "host": "https://sample.com",
    "dnsLookupTimeTakenInNanos": 1234,
    "tcpConnectionTimeTakenInNanos": 1234,
    "connectTimeTakenInNanos": 1234,
    "preTransferTimeTakenInNanos": 1234,
    "isTLSEnabled": true,
    "tlsHandshakeTimeTakenInNanos": 1234,
    "serverProcessingTimeTakenInNanos": 1234,
    "isConnectionReused": true
  }
]
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
go build -ldflags "<other linking params> -X github.com/sinhashubham95/info.BuildStamp=$buildStamp -X github.com/sinhashubham95/info.GitCommitID=$commitId -X github.com/sinhashubham95/info.GitPrimaryBranch=$2 -X github.com/sinhashubham95/info.GitURL=$gitUrl -X github.com/sinhashubham95/info.UserName=$userName -X github.com/sinhashubham95/info.HostName=$hostName  -X \"github.com/sinhashubham95/info.GitCommitTime=$commitTime\" -X \"github.com/sinhashubham95/info.GitCommitAuthor=$commitAuthor\""
./${APPLICATION_NAME} -env=${ENVIRONMENT_NAME} -name=${APPLICATION_NAME} -port=${APPLICATION_PORT} -version=${APPLICATION_VERSION}
```

```json
{
  "application": {
    "env": "ENVIRONMENT",
    "name": "APPLICATION_NAME",
    "version": "APPLICATION_VERSION"
  },
  "git": {
    "username": "s0s01qp",
    "hostName": "m-C02WV1L6HTD5",
    "buildStamp": "2019-08-22_09:44:04PM",
    "commitAuthor": "Shubham Sinha ",
    "commitId": "836475215e3ecf0ef26e0d5b65a9db626568ef89",
    "commitTime": "2019-08-23 02:27:26 +0530",
    "branch": "master",
    "url": "https://gecgithub01.walmart.com/RT-Integrated-Fulfillment/gif-ui-bff.git"
  },
  "runtime": {
    "arch": "",
    "os": "",
    "port": 8080,
    "runtimeVersion": ""
  }
}
```

### Metrics - `/actuator/metrics`

This is used to get the runtime memory statistics for your application. You can find the definition of each of the fields [here](./models/memStats.go).

```json
{}
```

### Ping - `/actuator/ping`

This is the lightweight ping endpoint that can be used along with your load balancer. This is used to know the running status of your application.

### Shutdown - `/actuator/shutdown`

This is used to bring the application down.

### Thread dump - `/actuator/threadDump`

This is used to get the trace of all the goroutines.
