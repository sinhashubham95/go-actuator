package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HandlePing is the handler function for the ping endpoint
func HandlePing(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}
