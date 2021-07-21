package gin_test

import (
	"github.com/gin-gonic/gin"
	ginControllers "github.com/sinhashubham95/go-actuator/controllers/gin"
	"github.com/sinhashubham95/go-actuator/core"
	"github.com/sinhashubham95/go-actuator/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupRouterAndGetResponse(t *testing.T, endpoint int, path string) *httptest.ResponseRecorder {
	router := gin.Default()
	router.Use(core.GINTracer())
	ginControllers.ConfigureHandlers(&models.Config{
		Endpoints: []int{endpoint},
	}, router)

	request, err := http.NewRequest(http.MethodGet, path, nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)

	return w
}
