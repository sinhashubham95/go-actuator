package fastHTTP

import (
	"github.com/sinhashubham95/go-actuator/commons"
	"github.com/sinhashubham95/go-actuator/flags"
	"github.com/valyala/fasthttp"
	"net/http"
	"os"
	"strings"
)

// HandleEnv is the handler function for the env endpoint
func HandleEnv(ctx *fasthttp.RequestCtx) {
	body, err := commons.EncodeJSON(getEnvironmentVariables())
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

func getEnvironmentVariables() map[string]string {
	variables := make(map[string]string)
	for _, e := range os.Environ() {
		keyValue := strings.SplitN(e, commons.Equals, 2)
		if len(keyValue) == 2 {
			variables[keyValue[0]] = keyValue[1]
		}
	}
	variables[commons.Env] = flags.Env()
	return variables
}
