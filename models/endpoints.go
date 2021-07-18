package models

const (
	Env = iota
	HTTPTrace
	Info
	Metrics
	Ping
	Shutdown
	ThreadDump
)

// Endpoints is the list of endpoints supported
var Endpoints = []int{Env, HTTPTrace, Info, Metrics, Ping, Shutdown, ThreadDump}

// IsValidEndpoint is used to check whether an endpoint is valid or not
func IsValidEndpoint(endpoint int) bool {
	for _, e := range Endpoints {
		if endpoint == e {
			return true
		}
	}
	return false
}
