package models_test

import (
	"github.com/sinhashubham95/go-actuator/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidEndpoint(t *testing.T) {
	assert.True(t, models.IsValidEndpoint(models.Env))
	assert.True(t, models.IsValidEndpoint(models.HTTPTrace))
	assert.True(t, models.IsValidEndpoint(models.Info))
	assert.True(t, models.IsValidEndpoint(models.Metrics))
	assert.True(t, models.IsValidEndpoint(models.Ping))
	assert.True(t, models.IsValidEndpoint(models.Shutdown))
	assert.True(t, models.IsValidEndpoint(models.ThreadDump))
}
