package core_test

import (
	"github.com/sinhashubham95/go-actuator/core"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMetrics(t *testing.T) {
	assert.NotNil(t, core.GetMetrics())
}
