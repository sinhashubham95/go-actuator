package gin_test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"

	"github.com/sinhashubham95/go-actuator/commons"
	ginControllers "github.com/sinhashubham95/go-actuator/controllers/gin"
)

func TestHandleEnv(t *testing.T) {
	w := setupRouterAndGetResponse(t, commons.EnvEndpoint, ginControllers.HandleEnv)
	assert.Equal(t, http.StatusOK, w.Code)
}
