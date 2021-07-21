package fasthttp_test

import (
	"github.com/sinhashubham95/go-actuator/commons"
	"github.com/sinhashubham95/go-actuator/models"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHandleThreadDump(t *testing.T) {
	response := setupFastHTTPHandlersAndGetResponse(t, models.ThreadDump, commons.ThreadDumpEndpoint)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	bytes, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, string(bytes))
}
