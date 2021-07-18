package netHTTP

import (
	"net/http"
	"path/filepath"

	"github.com/sinhashubham95/go-actuator/commons"
	"github.com/sinhashubham95/go-actuator/models"
)

func ConfigureHandlers(config *models.Config, mux *http.ServeMux) {
	for _, e := range config.Endpoints {
		// now one by one add the handler of each endpoint
		switch e {
		case models.Env:
			mux.HandleFunc(filepath.Join(config.Prefix, commons.EnvEndpoint), handle(HandleEnv))
		case models.Health:
			mux.HandleFunc(filepath.Join(config.Prefix, commons.HealthEndpoint), handle(HandleHealth))
		case models.HTTPTrace:
			mux.HandleFunc(filepath.Join(config.Prefix, commons.HTTPTraceEndpoint), handle(HandleHTTPTrace))
		case models.Info:
			mux.HandleFunc(filepath.Join(config.Prefix, commons.InfoEndpoint), handle(HandleInfo))
		case models.Metrics:
			mux.HandleFunc(filepath.Join(config.Prefix, commons.MetricsEndpoint), handle(HandleMetrics))
		case models.Ping:
			mux.HandleFunc(filepath.Join(config.Prefix, commons.PingEndpoint), handle(HandlePing))
		case models.Shutdown:
			mux.HandleFunc(filepath.Join(config.Prefix, commons.ShutdownEndpoint), handle(HandleShutdown))
		case models.ThreadDump:
			mux.HandleFunc(filepath.Join(config.Prefix, commons.ThreadDumpEndpoint), handle(HandleThreadDump))
		}
	}
}

func handle(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// first make sure the request is a get request
		if request.Method == http.MethodGet {
			handler(writer, request)
			return
		}
		writer.WriteHeader(http.StatusNotFound)
	}
}
