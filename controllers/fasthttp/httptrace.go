package fasthttp

import (
	"github.com/valyala/fasthttp"
	"net/http"

	"github.com/sinhashubham95/go-actuator/commons"
	"github.com/sinhashubham95/go-actuator/core"
)

// HandleHTTPTrace is used to handle the http trace request
func HandleHTTPTrace(ctx *fasthttp.RequestCtx) {
	body, err := EncodeJSON(core.GetHTTPTrace())
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
