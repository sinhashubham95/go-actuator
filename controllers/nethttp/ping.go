package nethttp

import (
	"net/http"
)

// HandlePing is the handler function for the ping endpoint
func HandlePing(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusOK)
}
