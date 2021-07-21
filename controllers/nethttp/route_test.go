package nethttp_test

import (
	"github.com/sinhashubham95/go-actuator/commons"
	netHTTPControllers "github.com/sinhashubham95/go-actuator/controllers/nethttp"
	"github.com/sinhashubham95/go-actuator/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInvalidRoute(t *testing.T) {
	w := setupMuxAndGetResponse(t, 9999, commons.PingEndpoint)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestInvalidMethod(t *testing.T) {
	mux := &http.ServeMux{}
	netHTTPControllers.ConfigureHandlers(&models.Config{Endpoints: []int{models.Ping}}, mux)

	request, err := http.NewRequest(http.MethodHead, commons.PingEndpoint, nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	mux.ServeHTTP(w, request)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
