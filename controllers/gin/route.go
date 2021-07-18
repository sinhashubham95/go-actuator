package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/sinhashubham95/go-actuator/commons"

	"github.com/sinhashubham95/go-actuator/models"
)

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
			actuator.GET(commons.HealthEndpoint, HandleInfo)
		case models.Metrics:
			actuator.GET(commons.HealthEndpoint, HandleMetrics)
		case models.Ping:
			actuator.GET(commons.HealthEndpoint, HandlePing)
		case models.Shutdown:
			actuator.GET(commons.HealthEndpoint, HandleShutdown)
		case models.ThreadDump:
			actuator.GET(commons.HealthEndpoint, HandleThreadDump)
		}
	}
}
