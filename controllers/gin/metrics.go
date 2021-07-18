package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/sinhashubham95/go-actuator/controllers/core"
)

// HandleMetrics is the handler function for the metrics endpoint
func HandleMetrics(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, core.GetMetrics())
}
