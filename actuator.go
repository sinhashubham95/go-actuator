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

var defaultEndpoints = []int{Info, Ping}

// HealthCheckFunc is the implementation to be called in case of a health check.
type HealthCheckFunc func(ctx context.Context) error

// HealthChecker is the set of details corresponding to a health check.
// For the health check, a custom function HealthChecker.Func has to be passed, which will be called during the health check.
// HealthChecker.IsMandatory decides whether this check will create an impact on the overall health check result.
type HealthChecker struct {
	Key         string
	Func        HealthCheckFunc
	IsMandatory bool
}

// HealthConfig is the set of configurable parameters for the health endpoint setup.
//
// HealthConfig.CacheDuration is the duration for which the health check details will be cached,
// during which the cached response of the prior health check performed will be reused.
// Defaults to 1 hour.
//
// HealthConfig.Timeout is the timeout which will be set in the context passed to the health check functions.
// It is the responsibility of the function implementation to honour the context cancellation.
// Defaults to 5 seconds.
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

func (config *Config) setDefaultsAndValidate() {
	if config.Endpoints == nil {
		config.Endpoints = defaultEndpoints
	}
	isHealthEnabled := false
	for _, endpoint := range config.Endpoints {
		if !isValidEndpoint(endpoint) {
			panic(fmt.Errorf("invalid endpoint %d provided", endpoint))
		}
		if endpoint == Health {
			isHealthEnabled = true
		}
	}
	if isHealthEnabled {
		if config.Health == nil {
			panic("health checker not configured")
		}
		if config.Health.CacheDuration == 0 {
			config.Health.CacheDuration = defaultHealthCheckCacheDuration
		}
		if config.Health.Timeout == 0 {
			config.Health.Timeout = defaultHealthCheckTimeout
		}
		if len(config.Health.Checkers) == 0 {
			panic("no health checkers provided")
		}
		keys := make(map[string]struct{})
		for _, checker := range config.Health.Checkers {
			if checker.Key == "" {
				panic("health checker key not provided")
			}
			if checker.Func == nil {
				panic(fmt.Errorf("health checker function not provided for %s", checker.Key))
			}
			if _, ok := keys[checker.Key]; ok {
				panic(fmt.Errorf("duplicate health checker key: %s", checker.Key))
			}
			keys[checker.Key] = struct{}{}
		}
	}
}

// GetActuatorHandler is used to get the handler function for the actuator endpoints
// This single handler is sufficient for handling all the endpoints.
func GetActuatorHandler(config *Config) http.HandlerFunc {
	if config == nil {
		config = &Config{}
	}
	config.setDefaultsAndValidate()
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

func getHandlerMap(config *Config) map[string]http.HandlerFunc {
	handlerMap := make(map[string]http.HandlerFunc, len(config.Endpoints))
	for _, e := range config.Endpoints {
		// now one by one add the handler of each endpoint
		switch e {
		case Env:
			handlerMap[envEndpoint] = getEnvHandler(config)
		case Info:
			handlerMap[infoEndpoint] = getInfoHandler(config)
		case Health:
			handlerMap[healthEndpoint] = getHealthHandler(config)
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
