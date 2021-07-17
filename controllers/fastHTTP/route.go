package fastHTTP

import (
	"github.com/sinhashubham95/go-actuator/models"
	"github.com/valyala/fasthttp"
	"net/http"
	"path/filepath"

	"github.com/sinhashubham95/go-actuator/commons"
)

// HandleRequest is the generic request handler for actuator
func HandleRequest(config *models.Config, ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case filepath.Join(config.Prefix, commons.EnvEndpoint):
		handle(config.Endpoints, models.Env, ctx, HandleEnv)
	case filepath.Join(config.Prefix, commons.HealthEndpoint):
		handle(config.Endpoints, models.Health, ctx, HandleHealth)
	case filepath.Join(config.Prefix, commons.InfoEndpoint):
		handle(config.Endpoints, models.Info, ctx, HandleInfo)
	case filepath.Join(config.Prefix, commons.MetricsEndpoint):
		handle(config.Endpoints, models.Metrics, ctx, HandleMetrics)
	case filepath.Join(config.Prefix, commons.PingEndpoint):
		handle(config.Endpoints, models.Ping, ctx, HandlePing)
	case filepath.Join(config.Prefix, commons.ShutdownEndpoint):
		handle(config.Endpoints, models.Shutdown, ctx, HandleShutdown)
	case filepath.Join(config.Prefix, commons.ThreadDumpEndpoint):
		handle(config.Endpoints, models.ThreadDump, ctx, HandleThreadDump)
	}
}

func handle(endpoints []int, endpoint int, ctx *fasthttp.RequestCtx, handler fasthttp.RequestHandler) {
	for _, e := range endpoints {
		if e == endpoint {
			// this endpoint is configured
			handler(ctx)
			return
		}
	}
	ctx.SetStatusCode(http.StatusNotFound)
}
