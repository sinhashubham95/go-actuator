package commons_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/sinhashubham95/go-actuator/commons"
)

func TestEncodeJSON(t *testing.T) {
	data := map[string]string{
		"naruto": "rocks",
	}
	bytes, err := commons.EncodeJSON(data)
	assert.NoError(t, err)
	assert.Contains(t, string(bytes), "naruto")
}
