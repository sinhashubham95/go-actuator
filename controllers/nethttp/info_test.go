package nethttp_test

import (
	"encoding/json"
	"github.com/sinhashubham95/go-actuator/commons"
	"github.com/sinhashubham95/go-actuator/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHandleInfo(t *testing.T) {
	w := setupMuxAndGetResponse(t, models.Info, commons.InfoEndpoint)
	assert.Equal(t, http.StatusOK, w.Code)

	var info map[string]interface{}
	err := json.NewDecoder(w.Body).Decode(&info)
	assert.NoError(t, err)
	assert.NotEmpty(t, info)
}
