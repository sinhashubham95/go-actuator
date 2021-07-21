package nethttp

import (
	"net/http"

	"github.com/sinhashubham95/go-actuator/core"
)

// HandleShutdown is the handler function for the shutdown endpoint
func HandleShutdown(writer http.ResponseWriter, request *http.Request) {
	core.Shutdown()
	writer.WriteHeader(http.StatusOK)
}
