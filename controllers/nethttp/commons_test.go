package nethttp_test

import (
	"errors"
	"github.com/sinhashubham95/go-actuator/commons"
	netHTTPControllers "github.com/sinhashubham95/go-actuator/controllers/nethttp"
	"github.com/sinhashubham95/go-actuator/core"
	"github.com/sinhashubham95/go-actuator/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var encodeJSON = commons.EncodeJSON
var getThreadDump = core.GetThreadDump

func setupMuxAndGetResponse(t *testing.T, endpoint int, path string) *httptest.ResponseRecorder {
	mux := &http.ServeMux{}
	netHTTPControllers.ConfigureHandlers(&models.Config{Endpoints: []int{endpoint}}, mux)

	request, err := http.NewRequest(http.MethodGet, path, nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	mux.ServeHTTP(w, request)

	return w
}

func mockEncodeJSONWithError() {
	encodeJSON = netHTTPControllers.EncodeJSON
	netHTTPControllers.EncodeJSON = func(interface{}) ([]byte, error) {
		return nil, errors.New("error")
	}
}

func unMockEncodeJSON() {
	netHTTPControllers.EncodeJSON = encodeJSON
}

func mockGetThreadDumpWithError() {
	getThreadDump = netHTTPControllers.GetThreadDump
	netHTTPControllers.GetThreadDump = func() ([]byte, error) {
		return nil, errors.New("error")
	}
}

func unMockGetThreadDump() {
	netHTTPControllers.GetThreadDump = getThreadDump
}
