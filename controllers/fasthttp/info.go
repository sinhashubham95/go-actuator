package fasthttp

import (
	"github.com/valyala/fasthttp"
	"net/http"

	"github.com/sinhashubham95/go-actuator/commons"
	"github.com/sinhashubham95/go-actuator/core"
)

// HandleInfo is the handler function for the info endpoint
func HandleInfo(ctx *fasthttp.RequestCtx) {
	body, err := commons.EncodeJSON(core.GetInfo())
	if err != nil {
		// some error occurred
		// send the error in the response
		ctx.SetContentType(commons.TextStringContentType)
		ctx.SetBody([]byte(err.Error()))
		ctx.SetStatusCode(http.StatusInternalServerError)
		return
	}
	// now once we have the correct response
	ctx.SetContentType(commons.ApplicationJSONContentType)
	ctx.SetBody(body)
	ctx.SetStatusCode(http.StatusOK)
}
