package core

import (
	"github.com/sinhashubham95/go-actuator/commons"
	"github.com/sinhashubham95/go-actuator/flags"
	"os"
	"strings"
)

// GetEnvironmentVariables is used to get the key-value pairs that are part of the environment variables
func GetEnvironmentVariables() map[string]string {
	variables := make(map[string]string)
	for _, e := range os.Environ() {
		keyValue := strings.SplitN(e, commons.Equals, 2)
		if len(keyValue) == 2 {
			variables[keyValue[0]] = keyValue[1]
		}
	}
	variables[commons.Env] = flags.Env()
	return variables
}
