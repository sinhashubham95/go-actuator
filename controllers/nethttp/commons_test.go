package nethttp_test

import (
	netHTTPControllers "github.com/sinhashubham95/go-actuator/controllers/nethttp"
	"github.com/sinhashubham95/go-actuator/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupMuxAndGetResponse(t *testing.T, endpoint int, path string) *httptest.ResponseRecorder {
	mux := &http.ServeMux{}
	netHTTPControllers.ConfigureHandlers(&models.Config{Endpoints: []int{endpoint}}, mux)

	request, err := http.NewRequest(http.MethodGet, path, nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	mux.ServeHTTP(w, request)

	return w
}
