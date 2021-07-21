package core_test

import (
	"github.com/sinhashubham95/go-actuator/commons"
	"github.com/sinhashubham95/go-actuator/core"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetInfo(t *testing.T) {
	info := core.GetInfo()
	assert.NotEmpty(t, info[commons.Application])
	assert.NotEmpty(t, info[commons.Git])
	assert.NotEmpty(t, info[commons.Runtime])
}
