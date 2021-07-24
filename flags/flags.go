package flags

import (
	flag "github.com/spf13/pflag"
	"os"

	"github.com/sinhashubham95/go-actuator/commons"
)

var flagSet = flag.NewFlagSet(commons.ActuatorFlagSetName, flag.ContinueOnError)

var (
	env     = flagSet.String(commons.Env, commons.EnvDefaultValue, commons.EnvUsage)
	name    = flagSet.String(commons.Name, commons.NameDefaultValue, commons.NameUsage)
	port    = flagSet.Int(commons.Port, commons.PortDefaultValue, commons.PortUsage)
	version = flagSet.String(commons.Version, commons.VersionDefaultValue, commons.VersionUsage)
)

func init() {
	err := flagSet.Parse(os.Args[1:])
	if err != nil {
		return
	}
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
