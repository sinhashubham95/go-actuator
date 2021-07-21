package core_test

import (
	"github.com/sinhashubham95/go-actuator/core"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetThreadDump(t *testing.T) {
	dump, err := core.GetThreadDump()
	assert.NoError(t, err)
	assert.NotEmpty(t, string(dump))
}
