package actuator

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Endpoints enumeration
const (
	Env = iota
	Info
	Health
	Metrics
	Ping
	Shutdown
	ThreadDump
)

// AllEndpoints is the list of endpoints supported
var AllEndpoints = []int{Env, Info, Health, Metrics, Ping, Shutdown, ThreadDump}

var defaultEndpoints = []int{Info, Ping, Health}

// HealthCheckFunc is the implementation to be called in case of a health check.
type HealthCheckFunc func(ctx context.Context) error

// HealthChecker is the set of details corresponding to a health check.
// For the health check, a custom function `HealthChecker.Func` has to be passed, which will be called during the health check.
// `HealthChecker.IsMandatory` decides whether this check will create an impact on the overall health check result.
type HealthChecker struct {
	Key         string
	Func        HealthCheckFunc
	IsMandatory bool
}

// HealthConfig is the set of configurable parameters for the health endpoint setup.
type HealthConfig struct {
	CacheDuration time.Duration
	Timeout       time.Duration
	Checkers      []HealthChecker
}

// Config is the set of configurable parameters for the actuator setup
type Config struct {
	Endpoints []int
	Env       string
	Name      string
	Port      int
	Version   string
	Health    *HealthConfig // optional, health check config
}

func (config *Config) validate() {
	for _, endpoint := range config.Endpoints {
		if !isValidEndpoint(endpoint) {
			panic(fmt.Errorf("invalid endpoint %d provided", endpoint))
		}
	}
}

// Default is used to fill the default configs in case of any missing ones
func (config *Config) setDefaults() {
	if config.Endpoints == nil {
		config.Endpoints = defaultEndpoints
	}
}

// GetActuatorHandler is used to get the handler function for the actuator endpoints
// This single handler is sufficient for handling all the endpoints.
func GetActuatorHandler(config *Config) http.HandlerFunc {
	if config == nil {
		config = &Config{}
	}
	handleConfigs(config)
	handlerMap := getHandlerMap(config)
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			// method not allowed for the requested resource
			sendStringResponse(writer, http.StatusMethodNotAllowed, methodNotAllowedError)
			return
		}
		endpoint := fmt.Sprintf("/%s", getLastStringAfterDelimiter(request.URL.Path, slash))
		if handler, ok := handlerMap[endpoint]; ok {
			handler(writer, request)
			return
		}
		// incorrect endpoint
		// or endpoint not enabled
		sendStringResponse(writer, http.StatusNotFound, notFoundError)
	}
}

func handleConfigs(config *Config) {
	config.validate()
	config.setDefaults()
}

func getHandlerMap(config *Config) map[string]http.HandlerFunc {
	handlerMap := make(map[string]http.HandlerFunc, len(config.Endpoints))
	for _, e := range config.Endpoints {
		// now one by one add the handler of each endpoint
		switch e {
		case Env:
			handlerMap[envEndpoint] = getEnvHandler(config)
		case Info:
			handlerMap[infoEndpoint] = getInfoHandler(config)
		case Metrics:
			handlerMap[metricsEndpoint] = handleMetrics
		case Ping:
			handlerMap[pingEndpoint] = handlePing
		case Shutdown:
			handlerMap[shutdownEndpoint] = handleShutdown
		case ThreadDump:
			handlerMap[threadDumpEndpoint] = handleThreadDump
		}
	}
	return handlerMap
}
