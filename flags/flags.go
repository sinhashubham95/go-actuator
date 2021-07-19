package flags

import (
	"github.com/sinhashubham95/go-actuator/commons"
	flag "github.com/spf13/pflag"
)

var (
	env     = flag.String(commons.Env, commons.EnvDefaultValue, commons.EnvUsage)
	name    = flag.String(commons.Name, commons.NameDefaultValue, commons.NameUsage)
	port    = flag.Int(commons.Port, commons.PortDefaultValue, commons.PortUsage)
	version = flag.String(commons.Version, commons.VersionDefaultValue, commons.VersionUsage)
)

func init() {
	flag.Parse()
}

// Env is the environment where the application is running
func Env() string {
	return *env
}

// Name is the application name
func Name() string {
	return *name
}

// Port is the port number where the application is running
func Port() int {
	return *port
}

// Version is the application build version
func Version() string {
	return *version
}
