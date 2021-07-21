package gin_test

import (
	"github.com/sinhashubham95/go-actuator/commons"
	"github.com/sinhashubham95/go-actuator/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHandlePing(t *testing.T) {
	w := setupRouterAndGetResponse(t, models.Ping, commons.PingEndpoint)
	assert.Equal(t, http.StatusOK, w.Code)
}
