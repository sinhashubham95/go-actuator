package commons

// Common Constants
const (
	Colon                 = ":"
	DefaultActuatorPrefix = "/actuator"
	Env                   = "env"
	EnvDefaultValue       = "no environment specified"
	EnvUsage              = "env is the environment where the application is running"
	Equals                = "="
	Port                  = "port"
	PortDefaultValue      = 8080
	PortUsage             = "port is the port number on which the application is running"
)

// Endpoints
const (
	EnvEndpoint        = "/env"
	HealthEndpoint     = "/health"
	InfoEndpoint       = "/info"
	MetricsEndpoint    = "/metrics"
	PingEndpoint       = "/ping"
	ShutdownEndpoint   = "/shutdown"
	ThreadDumpEndpoint = "/threadDump"
)

// Response constants
const (
	ApplicationJSONContentType = "application/json"
	TextStringContentType      = "text/string"
)
