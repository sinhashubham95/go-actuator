package netHTTP

import (
	"net/http"

	"github.com/sinhashubham95/go-actuator/commons"
	"github.com/sinhashubham95/go-actuator/controllers/core"
)

// HandleHTTPTrace is used to handle the http trace request
func HandleHTTPTrace(writer http.ResponseWriter, request *http.Request) {
	body, err := commons.EncodeJSON(core.GetHTTPTrace())
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
