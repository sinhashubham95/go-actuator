package gin_test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupRouterAndGetResponse(t *testing.T, endpoint string, handler gin.HandlerFunc) *httptest.ResponseRecorder {
	router := gin.Default()
	router.GET(endpoint, handler)

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)

	return w
}
