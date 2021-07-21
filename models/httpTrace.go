package models

import (
	"net/http"
	"time"
)

// HTTPTraceRequest is the set of useful information to trace the http request
type HTTPTraceRequest struct {
	Method  string              `json:"method"`
	URL     string              `json:"url"`
	Headers map[string][]string `json:"headers"`
}

// HTTPTraceResponse is the set of useful information to trace the http response
type HTTPTraceResponse struct {
	Status  int                 `json:"status"`
	Headers map[string][]string `json:"headers"`
}

// HTTPTraceResult is the set of useful information to trace the http request to your application
// and response from your application
type HTTPTraceResult struct {
	Timestamp time.Time          `json:"timestamp"`
	Duration  time.Duration      `json:"duration"`
	Request   *HTTPTraceRequest  `json:"request"`
	Response  *HTTPTraceResponse `json:"response"`
}

// HTTPStatusRecorder extends the response writer to handle the tracing of the status code value
type HTTPStatusRecorder struct {
	http.ResponseWriter
	StatusCode int
}

// WriteHeader saves an HTTP response header with the provided
// status code.
func (r *HTTPStatusRecorder) WriteHeader(statusCode int) {
	r.StatusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}
