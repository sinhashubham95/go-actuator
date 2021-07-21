package fasthttp_test

import (
	"encoding/json"
	"github.com/sinhashubham95/go-actuator/commons"
	"github.com/sinhashubham95/go-actuator/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHandleMetrics(t *testing.T) {
	response := setupFastHTTPHandlersAndGetResponse(t, models.Metrics, commons.MetricsEndpoint)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	var metrics *models.MemStats
	err := json.NewDecoder(response.Body).Decode(&metrics)
	assert.NoError(t, err)
	assert.NotNil(t, metrics)
}

func TestHandleMetricsEncodeJSONError(t *testing.T) {
	mockEncodeJSONWithError()
	defer unMockEncodeJSON()
	response := setupFastHTTPHandlersAndGetResponse(t, models.Metrics, commons.MetricsEndpoint)
	assert.Equal(t, http.StatusInternalServerError, response.StatusCode)
}
