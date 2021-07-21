package models_test

import (
	"github.com/sinhashubham95/go-actuator/commons"
	"github.com/sinhashubham95/go-actuator/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigValidateInvalidPrefix(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				assert.Error(t, err)
				assert.Equal(t, "invalid prefix provided", err.Error())
			}
		}
	}()
	config := &models.Config{
		Prefix: "postgres://user:abc{DEf1=ghi@example.com:5432/db?sslmode=require",
	}
	config.Validate()
}

func TestConfigValidateInvalidEndpoints(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				assert.Error(t, err)
				assert.Equal(t, "invalid endpoint 999 provided", err.Error())
			}
		}
	}()
	config := &models.Config{
		Endpoints: []int{999},
	}
	config.Validate()
}

func TestConfigDefault(t *testing.T) {
	config := &models.Config{}
	config.Default()
	assert.Equal(t, commons.DefaultActuatorPrefix, config.Prefix)
	assert.Equal(t, models.Endpoints, config.Endpoints)
}
