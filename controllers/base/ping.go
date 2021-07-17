package base

import (
	"github.com/valyala/fasthttp"
	"net/http"
)

// HandlePing is the handler function for the ping endpoint
func HandlePing(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(http.StatusOK)
}
