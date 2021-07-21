package core

import (
	"github.com/gin-gonic/gin"
	"github.com/valyala/fasthttp"
	"net/http"
	"sync"
	"time"

	"github.com/sinhashubham95/go-actuator/commons"
	"github.com/sinhashubham95/go-actuator/models"
)

var httpTraceResultsMu sync.Mutex
var httpTraceResults []*models.HTTPTraceResult

func WrapFastHTTPHandler(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		// for each new request create a result and trace the request
		result := &models.HTTPTraceResult{
			Timestamp: time.Now(),
			Request: &models.HTTPTraceRequest{
				Method:  string(ctx.Method()),
				URL:     string(ctx.RequestURI()),
				Headers: getFastHTTPRequestHeaders(ctx),
			},
		}
		// now process the request
		handler(ctx)
		// now trace the response
		result.Duration = time.Now().Sub(result.Timestamp)
		result.Response = &models.HTTPTraceResponse{
			Status:  ctx.Response.StatusCode(),
			Headers: getFastHTTPResponseHeaders(ctx),
		}
		// save this result
		saveResult(result)
	}
}

func GINTracer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// for each new request create a result and trace the request
		result := &models.HTTPTraceResult{
			Timestamp: time.Now(),
			Request: &models.HTTPTraceRequest{
				Method:  ctx.Request.Method,
				URL:     ctx.Request.URL.String(),
				Headers: ctx.Request.Header,
			},
		}
		// now process the request
		ctx.Next()
		// now trace the response
		result.Duration = time.Now().Sub(result.Timestamp)
		result.Response = &models.HTTPTraceResponse{
			Status:  ctx.Writer.Status(),
			Headers: ctx.Writer.Header(),
		}
		// save this result
		saveResult(result)
	}
}

func WrapNetHTTPHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// for each new request create a result and trace the request
		result := &models.HTTPTraceResult{
			Timestamp: time.Now(),
			Request: &models.HTTPTraceRequest{
				Method:  request.Method,
				URL:     request.URL.String(),
				Headers: request.Header,
			},
		}
		// now process the request
		recorder := &models.HTTPStatusRecorder{
			ResponseWriter: writer,
			StatusCode:     http.StatusOK,
		}
		handler(recorder, request)
		// now trace the response
		result.Duration = time.Now().Sub(result.Timestamp)
		result.Response = &models.HTTPTraceResponse{
			Status:  recorder.StatusCode,
			Headers: writer.Header(),
		}
		// save this result
		saveResult(result)
	}
}

// GetHTTPTrace is used to get the list of http trace results available
func GetHTTPTrace() []*models.HTTPTraceResult {
	httpTraceResultsMu.Lock()
	defer httpTraceResultsMu.Unlock()
	return httpTraceResults
}

func saveResult(result *models.HTTPTraceResult) {
	httpTraceResultsMu.Lock()
	defer httpTraceResultsMu.Unlock()

	if len(httpTraceResults)+1 > commons.HTTPTraceResultsSize {
		// append removing the last element
		httpTraceResults = append([]*models.HTTPTraceResult{result},
			httpTraceResults[:(commons.HTTPTraceResultsSize-1)]...)
	} else {
		// append blindly
		httpTraceResults = append([]*models.HTTPTraceResult{result}, httpTraceResults...)
	}
}

func getFastHTTPRequestHeaders(ctx *fasthttp.RequestCtx) map[string][]string {
	headers := make(map[string][]string)
	ctx.Request.Header.VisitAll(func(key, value []byte) {
		headers[string(key)] = []string{string(value)}
	})
	return headers
}

func getFastHTTPResponseHeaders(ctx *fasthttp.RequestCtx) map[string][]string {
	headers := make(map[string][]string)
	ctx.Response.Header.VisitAll(func(key, value []byte) {
		headers[string(key)] = []string{string(value)}
	})
	return headers
}
