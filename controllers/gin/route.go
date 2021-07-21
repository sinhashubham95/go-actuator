package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/sinhashubham95/go-actuator/commons"

	"github.com/sinhashubham95/go-actuator/models"
)

// ConfigureHandlers is used to add the respective handler functions for the configured actuator endpoints
// to the gin engine.
func ConfigureHandlers(config *models.Config, router *gin.Engine) {
	actuator := router.Group(config.Prefix)
	for _, e := range config.Endpoints {
		// now one by one add the handler of each endpoint
		switch e {
		case models.Env:
			actuator.GET(commons.EnvEndpoint, HandleEnv)
		case models.HTTPTrace:
			actuator.GET(commons.HTTPTraceEndpoint, HandleHTTPTrace)
		case models.Info:
			actuator.GET(commons.InfoEndpoint, HandleInfo)
		case models.Metrics:
			actuator.GET(commons.MetricsEndpoint, HandleMetrics)
		case models.Ping:
			actuator.GET(commons.PingEndpoint, HandlePing)
		case models.Shutdown:
			actuator.GET(commons.ShutdownEndpoint, HandleShutdown)
		case models.ThreadDump:
			actuator.GET(commons.ThreadDumpEndpoint, HandleThreadDump)
		}
	}
}
