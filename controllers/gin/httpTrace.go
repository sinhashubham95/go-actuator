package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/sinhashubham95/go-actuator/controllers/core"
	"net/http"
)

// HandleHTTPTrace is used to handle the http trace request
func HandleHTTPTrace(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, core.GetHTTPTrace())
}
