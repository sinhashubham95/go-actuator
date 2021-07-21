package fasthttp_test

import (
	"fmt"
	"github.com/sinhashubham95/go-actuator/commons"
	fastHTTPControllers "github.com/sinhashubham95/go-actuator/controllers/fasthttp"
	"github.com/sinhashubham95/go-actuator/core"
	"github.com/sinhashubham95/go-actuator/models"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"net/http"
	"testing"
)

func TestInvalidRoute(t *testing.T) {
	port := getRandomPortNumber()

	go func(endpoint int) {
		assert.NoError(t, fasthttp.ListenAndServe(fmt.Sprintf(":%d", port),
			core.WrapFastHTTPHandler(func(ctx *fasthttp.RequestCtx) {
				fastHTTPControllers.HandleRequest(&models.Config{Endpoints: []int{endpoint}}, ctx)
			})))
	}(models.Ping)

	request, err := http.NewRequest(http.MethodHead, fmt.Sprintf("http://localhost:%d%s",
		port, commons.PingEndpoint), nil)
	assert.NoError(t, err)

	response, err := http.DefaultClient.Do(request)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, response.StatusCode)
}

func TestInvalidEndpoint(t *testing.T) {
	port := getRandomPortNumber()

	go func(endpoint int) {
		assert.NoError(t, fasthttp.ListenAndServe(fmt.Sprintf(":%d", port),
			core.WrapFastHTTPHandler(func(ctx *fasthttp.RequestCtx) {
				fastHTTPControllers.HandleRequest(&models.Config{Endpoints: []int{endpoint}}, ctx)
			})))
	}(models.Ping)

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%d%s",
		port, commons.EnvEndpoint), nil)
	assert.NoError(t, err)

	response, err := http.DefaultClient.Do(request)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, response.StatusCode)
}
