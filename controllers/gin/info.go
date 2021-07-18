package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/sinhashubham95/go-actuator/core"
)

// HandleInfo is the handler function for the info endpoint
func HandleInfo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, core.GetInfo())
}
