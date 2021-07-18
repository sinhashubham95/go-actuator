package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/sinhashubham95/go-actuator/core"
)

// HandleThreadDump is the handler to get the thread dump
func HandleThreadDump(ctx *gin.Context) {
	body, err := core.GetThreadDump()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "%v", err)
	}
	ctx.String(http.StatusOK, string(body))
}
