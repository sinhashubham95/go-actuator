package core

import (
	"crypto/tls"
	"github.com/sinhashubham95/go-actuator/commons"
	"github.com/sinhashubham95/go-actuator/models"
	"net/http"
	"net/http/httptrace"
	"sync"
	"time"
)

var httpTraceResultsMu sync.Mutex
var httpTraceResults []*models.HTTPTraceResult

// WithClientTrace is used to attach the trace to the http request
// this is mandatory to populate the trace in the result set
func WithClientTrace(request *http.Request) *http.Request {
	// create a result
	result := &models.HTTPTraceResult{}

	// set the trace in the request context
	request.WithContext(httptrace.WithClientTrace(request.Context(), &httptrace.ClientTrace{
		DNSStart: func(httptrace.DNSStartInfo) {
			result.DNSStart = time.Now()
		},
		DNSDone: func(httptrace.DNSDoneInfo) {
			result.DNSDone = time.Now()
			result.DNSLookup = result.DNSDone.Sub(result.DNSStart)
		},
		ConnectStart: func(_, _ string) {
			result.TCPStart = time.Now()

			// when connecting to IP address
			if result.DNSStart.IsZero() {
				result.DNSStart = result.TCPStart
				result.DNSDone = result.TCPStart
			}
		},
		ConnectDone: func(_, _ string, _ error) {
			result.TCPDone = time.Now()
			result.TCPConnection = result.TCPDone.Sub(result.TCPStart)

			result.Connect = result.TCPDone.Sub(result.DNSStart)
		},
		TLSHandshakeStart: func() {
			result.IsTLS = true
			result.TLSStart = time.Now()
		},
		TLSHandshakeDone: func(_ tls.ConnectionState, _ error) {
			result.TLSDone = time.Now()
			result.TLSHandshake = result.TLSDone.Sub(result.TLSStart)

			result.PreTransfer = result.TLSDone.Sub(result.DNSStart)
		},
		GotConn: func(info httptrace.GotConnInfo) {
			if info.Reused {
				result.IsReused = true
			}
		},
		WroteRequest: func(_ httptrace.WroteRequestInfo) {
			result.ServerStart = time.Now()

			// when client does not use dial context or old package
			if result.DNSStart.IsZero() && result.TCPStart.IsZero() {
				result.DNSStart = result.ServerStart
				result.DNSDone = result.ServerStart
				result.TCPStart = result.ServerStart
				result.TCPDone = result.ServerStart
			}

			// when connection is reused, then dns lookup, tcp connection and tls handshake does not happen
			if result.IsReused {
				result.DNSStart = result.ServerStart
				result.DNSDone = result.ServerStart
				result.TCPStart = result.ServerStart
				result.TCPDone = result.ServerStart
				result.TLSStart = result.ServerStart
				result.TLSDone = result.ServerStart
			}
		},
		GotFirstResponseByte: func() {
			result.ServerDone = time.Now()
			result.ServerProcessing = result.ServerDone.Sub(result.ServerStart)

			result.Done = true
		},
	}))

	// save the result
	httpTraceResultsMu.Lock()
	defer httpTraceResultsMu.Unlock()
	httpTraceResults = append(httpTraceResults, result)
	cleanupHTTPTraceResults()

	// return the request
	return request
}

// GetHTTPTrace is used to get the list of http trace results available
func GetHTTPTrace() []*models.HTTPTraceResult {
	httpTraceResultsMu.Lock()
	defer httpTraceResultsMu.Unlock()
	cleanupHTTPTraceResults()
	return httpTraceResults
}

func cleanupHTTPTraceResults() {
	var numberOfElementsToBeRemoved = commons.HTTPTraceResultsSize - len(httpTraceResults)
	for i := range httpTraceResults {
		// check if we need to remove this result
		if numberOfElementsToBeRemoved <= 0 {
			break
		}
		// now check if we can remove this result
		if httpTraceResults[i].Done {
			// we can very well remove this
			numberOfElementsToBeRemoved--
			httpTraceResults = append(httpTraceResults[:i], httpTraceResults[i+1:]...)
		}
	}
}
