package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/sinhashubham95/go-actuator/core"
)

// HandleShutdown is the handler function for the shutdown endpoint
func HandleShutdown(ctx *gin.Context) {
	core.Shutdown()
	ctx.Status(http.StatusOK)
}
