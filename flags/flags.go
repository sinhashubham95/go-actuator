package flags

import (
	"flag"
	"github.com/sinhashubham95/go-actuator/commons"
)

var (
	env  = flag.String(commons.Env, commons.EnvDefaultValue, commons.EnvUsage)
	port = flag.Int(commons.Port, commons.PortDefaultValue, commons.PortUsage)
)

func init() {
	flag.Parse()
}

// Env is the environment where the application is running
func Env() string {
	return *env
}

// Port is the port number where the application is running
func Port() int {
	return *port
}
