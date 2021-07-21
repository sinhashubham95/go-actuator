package flags_test

import (
	"github.com/sinhashubham95/go-actuator/commons"
	"github.com/sinhashubham95/go-actuator/flags"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnv(t *testing.T) {
	assert.Equal(t, commons.EnvDefaultValue, flags.Env())
}

func TestName(t *testing.T) {
	assert.Equal(t, commons.NameDefaultValue, flags.Name())
}

func TestPort(t *testing.T) {
	assert.Equal(t, commons.PortDefaultValue, flags.Port())
}

func TestVersion(t *testing.T) {
	assert.Equal(t, commons.VersionDefaultValue, flags.Version())
}
