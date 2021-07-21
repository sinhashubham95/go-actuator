package fasthttp_test

import (
	"encoding/json"
	"github.com/sinhashubham95/go-actuator/commons"
	"github.com/sinhashubham95/go-actuator/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHandleInfo(t *testing.T) {
	response := setupFastHTTPHandlersAndGetResponse(t, models.Info, commons.InfoEndpoint)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	var info map[string]interface{}
	err := json.NewDecoder(response.Body).Decode(&info)
	assert.NoError(t, err)
	assert.NotEmpty(t, info)
}
