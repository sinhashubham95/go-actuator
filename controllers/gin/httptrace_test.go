package gin_test

import (
	"encoding/json"
	"github.com/sinhashubham95/go-actuator/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"

	"github.com/sinhashubham95/go-actuator/commons"
)

func TestHandleHTTPTrace(t *testing.T) {
	w := setupRouterAndGetResponse(t, models.HTTPTrace, commons.HTTPTraceEndpoint)
	assert.Equal(t, http.StatusOK, w.Code)

	var traces []*models.HTTPTraceResult
	err := json.NewDecoder(w.Body).Decode(&traces)
	assert.NoError(t, err)
	assert.NotEmpty(t, traces)
}
