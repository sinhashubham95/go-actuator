package core_test

import (
	"github.com/sinhashubham95/go-actuator/commons"
	"github.com/sinhashubham95/go-actuator/core"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetEnvironmentVariables(t *testing.T) {
	_ = os.Setenv("key1", "value1")
	_ = os.Setenv("key2", "value2")
	environ := core.GetEnvironmentVariables()
	assert.Equal(t, "value1", environ["key1"])
	assert.Equal(t, "value2", environ["key2"])
	assert.Equal(t, commons.EnvDefaultValue, environ[commons.Env])
}
