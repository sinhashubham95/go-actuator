package nethttp_test

import (
	"encoding/json"
	"github.com/sinhashubham95/go-actuator/commons"
	"github.com/sinhashubham95/go-actuator/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHandleHTTPTrace(t *testing.T) {
	w := setupMuxAndGetResponse(t, models.HTTPTrace, commons.HTTPTraceEndpoint)
	assert.Equal(t, http.StatusOK, w.Code)

	var traces []*models.HTTPTraceResult
	err := json.NewDecoder(w.Body).Decode(&traces)
	assert.NoError(t, err)
	assert.Nil(t, traces)
}

func TestHandleHTTPTraceEncodeJSONError(t *testing.T) {
	mockEncodeJSONWithError()
	defer unMockEncodeJSON()
	setupMuxAndGetResponse(t, models.HTTPTrace, commons.HTTPTraceEndpoint)
}
