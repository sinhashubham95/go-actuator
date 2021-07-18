package commons

// Common Constants
const (
	Application           = "app"
	Arch                  = "arch"
	BuildStamp            = "buildStamp"
	DefaultActuatorPrefix = "/actuator"
	Env                   = "env"
	EnvDefaultValue       = "no environment specified"
	EnvUsage              = "env is the environment where the application is running"
	Equals                = "="
	Git                   = "git"
	GitCommitAuthor       = "commitAuthor"
	GitCommitID           = "commitId"
	GitCommitTime         = "commitTime"
	GitPrimaryBranch      = "branch"
	GitURL                = "url"
	GoRoutines            = "goroutine"
	HostName              = "hostName"
	HTTPTraceResultsSize  = 100
	OS                    = "os"
	Name                  = "name"
	NameDefaultValue      = "no application name specified"
	NameUsage             = "name is the application name"
	Port                  = "port"
	PortDefaultValue      = 8080
	PortUsage             = "port is the port number on which the application is running"
	Runtime               = "runtime"
	RuntimeVersion        = "runtimeVersion"
	Username              = "username"
	Version               = "version"
	VersionDefaultValue   = "no version specified"
	VersionUsage          = "version is the current application build version"
)

// Endpoints
const (
	EnvEndpoint        = "/env"
	HealthEndpoint     = "/health"
	HTTPTraceEndpoint  = "/httpTrace"
	InfoEndpoint       = "/info"
	MetricsEndpoint    = "/metrics"
	PingEndpoint       = "/ping"
	ShutdownEndpoint   = "/shutdown"
	ThreadDumpEndpoint = "/threadDump"
)

// Response constants
const (
	ContentTypeHeader          = "Content-Type"
	ApplicationJSONContentType = "application/json"
	TextStringContentType      = "text/string"
)
