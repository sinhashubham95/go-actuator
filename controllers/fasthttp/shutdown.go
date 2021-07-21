package fasthttp

import (
	"github.com/valyala/fasthttp"
	"net/http"

	"github.com/sinhashubham95/go-actuator/core"
)

// HandleShutdown is the handler function for the shutdown endpoint
func HandleShutdown(ctx *fasthttp.RequestCtx) {
	core.Shutdown()
	ctx.SetStatusCode(http.StatusOK)
}
