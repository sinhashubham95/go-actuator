package actuator

import (
	"github.com/gin-gonic/gin"
	"github.com/valyala/fasthttp"
	"net/http"

	fastHTTPControllers "github.com/sinhashubham95/go-actuator/controllers/fasthttp"
	ginControllers "github.com/sinhashubham95/go-actuator/controllers/gin"
	netHTTPControllers "github.com/sinhashubham95/go-actuator/controllers/nethttp"
	"github.com/sinhashubham95/go-actuator/models"
)

// GetFastHTTPActuatorHandler is used to get the request handler for fast http
func GetFastHTTPActuatorHandler(config *models.Config) fasthttp.RequestHandler {
	handleConfigs(config)
	return func(ctx *fasthttp.RequestCtx) {
		fastHTTPControllers.HandleRequest(config, ctx)
	}
}

// ConfigureGINActuatorEngine is used to configure the gin engine with the actuator handlers
func ConfigureGINActuatorEngine(config *models.Config, engine *gin.Engine) {
	handleConfigs(config)
	ginControllers.ConfigureHandlers(config, engine)
}

// ConfigureNetHTTPHandler is used to configure the net http mux with the actuator handlers
func ConfigureNetHTTPHandler(config *models.Config, mux *http.ServeMux) {
	handleConfigs(config)
	netHTTPControllers.ConfigureHandlers(config, mux)
}

func handleConfigs(config *models.Config) {
	config.Validate()
	config.Default()
}
