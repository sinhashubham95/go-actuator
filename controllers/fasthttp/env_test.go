package fasthttp_test

import (
	"encoding/json"
	"github.com/sinhashubham95/go-actuator/commons"
	"github.com/sinhashubham95/go-actuator/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHandleEnv(t *testing.T) {
	response := setupFastHTTPHandlersAndGetResponse(t, models.Env, commons.EnvEndpoint)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	var env map[string]interface{}
	err := json.NewDecoder(response.Body).Decode(&env)
	assert.NoError(t, err)
	assert.NotEmpty(t, env)
}
