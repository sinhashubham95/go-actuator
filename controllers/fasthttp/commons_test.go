package fasthttp_test

import (
	"errors"
	"fmt"
	"github.com/sinhashubham95/go-actuator/commons"
	fastHTTPControllers "github.com/sinhashubham95/go-actuator/controllers/fasthttp"
	"github.com/sinhashubham95/go-actuator/core"
	"github.com/sinhashubham95/go-actuator/models"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"net/http"
	"sync"
	"testing"
)

var encodeJSON = commons.EncodeJSON
var getThreadDump = core.GetThreadDump

var portMu sync.Mutex
var port = 1201

func getRandomPortNumber() int {
	portMu.Lock()
	defer portMu.Unlock()
	port += 10
	return port
}

func setupFastHTTPHandlersAndGetResponse(t *testing.T, endpoint int, path string) *http.Response {
	port := getRandomPortNumber()

	go func(endpoint int) {
		assert.NoError(t, fasthttp.ListenAndServe(fmt.Sprintf(":%d", port),
			core.WrapFastHTTPHandler(func(ctx *fasthttp.RequestCtx) {
				fastHTTPControllers.HandleRequest(&models.Config{Endpoints: []int{endpoint}}, ctx)
			})))
	}(endpoint)

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%d%s", port, path), nil)
	assert.NoError(t, err)

	response, err := http.DefaultClient.Do(request)
	assert.NoError(t, err)

	return response
}

func mockEncodeJSONWithError() {
	encodeJSON = fastHTTPControllers.EncodeJSON
	fastHTTPControllers.EncodeJSON = func(interface{}) ([]byte, error) {
		return nil, errors.New("error")
	}
}

func unMockEncodeJSON() {
	fastHTTPControllers.EncodeJSON = encodeJSON
}

func mockGetThreadDumpWithError() {
	getThreadDump = fastHTTPControllers.GetThreadDump
	fastHTTPControllers.GetThreadDump = func() ([]byte, error) {
		return nil, errors.New("error")
	}
}

func unMockGetThreadDump() {
	fastHTTPControllers.GetThreadDump = getThreadDump
}
