package gin_test

import (
	"github.com/sinhashubham95/go-actuator/commons"
	"github.com/sinhashubham95/go-actuator/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHandleShutdown(t *testing.T) {
	w := setupRouterAndGetResponse(t, models.Shutdown, commons.ShutdownEndpoint)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
