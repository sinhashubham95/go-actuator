package fasthttp_test

import (
	"encoding/json"
	"github.com/sinhashubham95/go-actuator/commons"
	"github.com/sinhashubham95/go-actuator/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHandleHTTPTrace(t *testing.T) {
	response := setupFastHTTPHandlersAndGetResponse(t, models.HTTPTrace, commons.HTTPTraceEndpoint)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	var traces []*models.HTTPTraceResult
	err := json.NewDecoder(response.Body).Decode(&traces)
	assert.NoError(t, err)
	assert.NotEmpty(t, traces)
}
