package nethttp

import (
	"net/http"

	"github.com/sinhashubham95/go-actuator/commons"
)

// HandleThreadDump is the handler to get the thread dump
func HandleThreadDump(writer http.ResponseWriter, _ *http.Request) {
	body, err := GetThreadDump()
	if err != nil {
		// some error occurred
		// send the error in the response
		writer.Header().Add(commons.ContentTypeHeader, commons.TextStringContentType)
		_, _ = writer.Write([]byte(err.Error()))
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	// now once we have the correct response
	writer.Header().Add(commons.ContentTypeHeader, commons.TextStringContentType)
	_, _ = writer.Write(body)
	writer.WriteHeader(http.StatusOK)
}
