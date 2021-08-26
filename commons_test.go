package actuator

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"runtime/pprof"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

var (
	tempEncodeJSONFunction  = encodeJSONFunction
	tempPprofLookupFunction = pprofLookupFunction
)

func setupMuxAndGetResponse(t *testing.T, endpoint int, path string) *httptest.ResponseRecorder {
	return setupMuxWithConfigAndGetResponseForMethod(t, &Config{Endpoints: []int{endpoint}}, http.MethodGet, path)
}

func setupMuxWithConfigAndGetResponseForMethod(t *testing.T, config *Config, method, path string) *httptest.ResponseRecorder {
	mux := &http.ServeMux{}
	mux.Handle("/", GetActuatorHandler(config))

	request, err := http.NewRequest(method, path, nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	mux.ServeHTTP(w, request)

	return w
}

func getTypedJSONBody(t *testing.T, body *bytes.Buffer, v interface{}) {
	decoder := jsoniter.NewDecoder(bytes.NewReader(body.Bytes()))
	err := decoder.Decode(v)
	assert.NoError(t, err)
}

func mockEncodeJSONWithError() {
	tempEncodeJSONFunction = encodeJSONFunction
	encodeJSONFunction = func(interface{}) ([]byte, error) {
		return nil, errors.New("error")
	}
}

func unMockEncodeJSON() {
	encodeJSONFunction = tempEncodeJSONFunction
}

func mockPprofLookupWithError() {
	tempPprofLookupFunction = pprofLookupFunction
	pprofLookupFunction = func(name string) *pprof.Profile {
		return nil
	}
}

func unMockPprofLookup() {
	pprofLookupFunction = tempPprofLookupFunction
}
