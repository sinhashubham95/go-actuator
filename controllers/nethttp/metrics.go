package nethttp

import (
	"net/http"

	"github.com/sinhashubham95/go-actuator/commons"
	"github.com/sinhashubham95/go-actuator/core"
)

// HandleMetrics is the handler function for the metrics endpoint
func HandleMetrics(writer http.ResponseWriter, request *http.Request) {
	body, err := commons.EncodeJSON(core.GetMetrics())
	if err != nil {
		// some error occurred
		// send the error in the response
		writer.Header().Add(commons.ContentTypeHeader, commons.TextStringContentType)
		_, _ = writer.Write([]byte(err.Error()))
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	// now once we have the correct response
	writer.Header().Add(commons.ContentTypeHeader, commons.ApplicationJSONContentType)
	_, _ = writer.Write(body)
	writer.WriteHeader(http.StatusOK)
}
