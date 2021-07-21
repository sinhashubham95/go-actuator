package core_test

import (
	"github.com/sinhashubham95/go-actuator/core"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShutdown(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotNil(t, r)
	}()
	core.Shutdown()
}
