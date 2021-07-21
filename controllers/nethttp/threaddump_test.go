package nethttp_test

import (
	"github.com/sinhashubham95/go-actuator/commons"
	"github.com/sinhashubham95/go-actuator/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHandleThreadDump(t *testing.T) {
	w := setupMuxAndGetResponse(t, models.ThreadDump, commons.ThreadDumpEndpoint)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body.String())
}
