package fasthttp_test

import (
	"fmt"
	fastHTTPControllers "github.com/sinhashubham95/go-actuator/controllers/fasthttp"
	"github.com/sinhashubham95/go-actuator/core"
	"github.com/sinhashubham95/go-actuator/models"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"math/rand"
	"net/http"
	"testing"
)

func getRandomPortNumber() int {
	return rand.Intn(9800) + 100
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
