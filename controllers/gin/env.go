package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/sinhashubham95/go-actuator/controllers/base"
)

// HandleEnv is the handler function for the env endpoint
func HandleEnv(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, base.GetEnvironmentVariables())
}
