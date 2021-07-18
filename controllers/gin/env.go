package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/sinhashubham95/go-actuator/core"
)

// HandleEnv is the handler function for the env endpoint
func HandleEnv(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, core.GetEnvironmentVariables())
}
