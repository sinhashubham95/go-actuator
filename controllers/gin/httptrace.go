package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/sinhashubham95/go-actuator/core"
)

// HandleHTTPTrace is used to handle the http trace request
func HandleHTTPTrace(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, core.GetHTTPTrace())
}
