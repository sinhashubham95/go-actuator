package models

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/sinhashubham95/go-actuator/commons"
)

// Config is the set of configurable parameters for the actuator setup
type Config struct {
	Prefix    string
	Endpoints []int
}

// Validate is used to validate the configurations passed.
// This method panics in case of any error.
func (config *Config) Validate() {
	if config.Prefix != "" {
		// validate the prefix path
		if _, err := url.Parse(config.Prefix); err != nil {
			// invalid prefix path
			panic(errors.New("invalid prefix provided"))
		}
	}
	for _, endpoint := range config.Endpoints {
		if !IsValidEndpoint(endpoint) {
			panic(fmt.Errorf("invalid endpoint %d provided", endpoint))
		}
	}
}

// Default is used to fill the default configs in case of any missing ones
func (config *Config) Default() {
	if config.Prefix == "" {
		config.Prefix = commons.DefaultActuatorPrefix
	}
	if config.Endpoints == nil {
		config.Endpoints = Endpoints
	}
}
