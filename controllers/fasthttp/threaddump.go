package fasthttp

import (
	"github.com/valyala/fasthttp"
	"net/http"

	"github.com/sinhashubham95/go-actuator/commons"
	"github.com/sinhashubham95/go-actuator/core"
)

// HandleThreadDump is the handler for getting thread dump
func HandleThreadDump(ctx *fasthttp.RequestCtx) {
	body, err := core.GetThreadDump()
	if err != nil {
		// some error occurred
		// send the error in the response
		ctx.SetContentType(commons.TextStringContentType)
		ctx.SetBody([]byte(err.Error()))
		ctx.SetStatusCode(http.StatusInternalServerError)
		return
	}
	// now once we have the correct response
	ctx.SetContentType(commons.TextStringContentType)
	ctx.SetBody(body)
	ctx.SetStatusCode(http.StatusOK)
}
