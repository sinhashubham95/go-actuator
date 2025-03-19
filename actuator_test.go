package actuator

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEmptyConfig(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			assert.Failf(t, "incorrect config", "%+v", r)
		}
	}()
	c := &Config{}
	c.setDefaultsAndValidate()
}

func TestValidateConfigWithIncorrectEndpoint(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				assert.Equal(t, fmt.Sprintf("invalid endpoint %d provided", 20), e.Error())
			}
		}
	}()
	c := &Config{
		Endpoints: []int{20},
	}
	c.setDefaultsAndValidate()
}

func TestSetDefaultsInConfig(t *testing.T) {
	c := &Config{}
	c.setDefaultsAndValidate()
	assert.Equal(t, defaultEndpoints, c.Endpoints)
}

func TestValidateConfigHealthNoConfiguration(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				assert.Equal(t, "health checker not configured", e.Error())
			}
		}
	}()
	c := &Config{Endpoints: []int{Health}}
	c.setDefaultsAndValidate()
}

func TestValidateConfigHealthNoChecker(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				assert.Equal(t, "no health checkers provided", e.Error())
			}
		}
	}()
	c := &Config{Endpoints: []int{Health}, Health: &HealthConfig{}}
	c.setDefaultsAndValidate()
}

func TestValidateConfigHealthCheckerKeyNotProvided(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				assert.Equal(t, "health checker key not provided", e.Error())
			}
		}
	}()
	c := &Config{Endpoints: []int{Health}, Health: &HealthConfig{
		Checkers: []HealthChecker{{}},
	}}
	c.setDefaultsAndValidate()
}

func TestValidateConfigHealthCheckerFunctionNotProvided(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				assert.Equal(t, "health checker function not provided for naruto", e.Error())
			}
		}
	}()
	c := &Config{Endpoints: []int{Health}, Health: &HealthConfig{
		Checkers: []HealthChecker{{Key: "naruto"}},
	}}
	c.setDefaultsAndValidate()
}

func TestValidateConfigHealthCheckerDuplicateKey(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				assert.Equal(t, "duplicate health checker key: naruto", e.Error())
			}
		}
	}()
	c := &Config{Endpoints: []int{Health}, Health: &HealthConfig{
		Checkers: []HealthChecker{
			{Key: "naruto", Func: func(_ context.Context) error { return nil }},
			{Key: "naruto", Func: func(_ context.Context) error { return nil }},
		},
	}}
	c.setDefaultsAndValidate()
}

func TestEnv(t *testing.T) {
	w := setupMuxAndGetResponse(t, Env, envEndpoint)
	assert.Equal(t, http.StatusOK, w.Code)

	var data map[string]string
	getTypedJSONBody(t, w.Body, &data)
	assert.NotEmpty(t, data)
}

func TestEnvNotConfigured(t *testing.T) {
	w := setupMuxAndGetResponse(t, Info, envEndpoint)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, notFoundError, w.Body.String())
}

func TestEnvInvalidMethod(t *testing.T) {
	w := setupMuxWithConfigAndGetResponseForMethod(t, &Config{Endpoints: []int{Env}}, http.MethodHead, envEndpoint)
	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
	assert.Equal(t, methodNotAllowedError, w.Body.String())
}

func TestEnvWithoutConfig(t *testing.T) {
	w := setupMuxWithConfigAndGetResponseForMethod(t, nil, http.MethodGet, envEndpoint)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, notFoundError, w.Body.String())
}

func TestEnvWithConfig(t *testing.T) {
	w := setupMuxWithConfigAndGetResponseForMethod(t, &Config{Endpoints: []int{Env}}, http.MethodGet, envEndpoint)
	assert.Equal(t, http.StatusOK, w.Code)

	var data map[string]string
	getTypedJSONBody(t, w.Body, &data)
	assert.NotEmpty(t, data)
}

func TestInfo(t *testing.T) {
	w := setupMuxAndGetResponse(t, Info, infoEndpoint)
	assert.Equal(t, http.StatusOK, w.Code)

	var data map[string]interface{}
	getTypedJSONBody(t, w.Body, &data)
	assert.NotEmpty(t, data)
}

func TestInfoNotConfigured(t *testing.T) {
	w := setupMuxAndGetResponse(t, Env, infoEndpoint)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, notFoundError, w.Body.String())
}

func TestInfoInvalidMethod(t *testing.T) {
	w := setupMuxWithConfigAndGetResponseForMethod(t, &Config{Endpoints: []int{Info}}, http.MethodHead, infoEndpoint)
	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
	assert.Equal(t, methodNotAllowedError, w.Body.String())
}

func TestInfoWithoutConfig(t *testing.T) {
	w := setupMuxWithConfigAndGetResponseForMethod(t, nil, http.MethodGet, infoEndpoint)
	assert.Equal(t, http.StatusOK, w.Code)

	var data map[string]interface{}
	getTypedJSONBody(t, w.Body, &data)
	assert.NotEmpty(t, data)
}

func TestHealth(t *testing.T) {
	defer clearHealthCheckCache()
	w := setupMuxWithConfigAndGetResponseForMethod(t,
		&Config{Endpoints: []int{Health}, Health: &HealthConfig{Checkers: []HealthChecker{{
			Key: "naruto",
			Func: func(_ context.Context) error {
				return nil
			},
		}}}},
		http.MethodGet, healthEndpoint)
	assert.Equal(t, http.StatusOK, w.Code)

	var data map[string]HealthCheckInfo
	getTypedJSONBody(t, w.Body, &data)
	assert.Equal(t, 1, len(data))
	for _, v := range data {
		assert.False(t, v.IsMandatory && !v.Success)
	}
}

func TestHealthFromCache(t *testing.T) {
	defer clearHealthCheckCache()

	w := setupMuxWithConfigAndGetResponseForMethod(t,
		&Config{Endpoints: []int{Health}, Health: &HealthConfig{Checkers: []HealthChecker{{
			Key: "naruto",
			Func: func(_ context.Context) error {
				return nil
			},
		}}}},
		http.MethodGet, healthEndpoint)
	assert.Equal(t, http.StatusOK, w.Code)
	var data map[string]HealthCheckInfo
	getTypedJSONBody(t, w.Body, &data)
	assert.Equal(t, 1, len(data))
	for _, v := range data {
		assert.False(t, v.IsMandatory && !v.Success)
	}

	w = setupMuxWithConfigAndGetResponseForMethod(t,
		&Config{Endpoints: []int{Health}, Health: &HealthConfig{Checkers: []HealthChecker{{
			Key: "naruto",
			Func: func(_ context.Context) error {
				return nil
			},
		}}}},
		http.MethodGet, healthEndpoint)
	assert.Equal(t, http.StatusOK, w.Code)
	getTypedJSONBody(t, w.Body, &data)
	assert.Equal(t, 1, len(data))
	for _, v := range data {
		assert.False(t, v.IsMandatory && !v.Success)
	}
}

func TestHealthNonMandatoryFailure(t *testing.T) {
	defer clearHealthCheckCache()
	w := setupMuxWithConfigAndGetResponseForMethod(t,
		&Config{Endpoints: []int{Health}, Health: &HealthConfig{Checkers: []HealthChecker{{
			Key: "naruto",
			Func: func(_ context.Context) error {
				return nil
			},
		}, {
			Key: "another naruto",
			Func: func(_ context.Context) error {
				return errors.New("some error")
			},
		}}}},
		http.MethodGet, healthEndpoint)
	assert.Equal(t, http.StatusOK, w.Code)
	var data map[string]HealthCheckInfo
	getTypedJSONBody(t, w.Body, &data)
	assert.Equal(t, 2, len(data))
	for _, v := range data {
		assert.False(t, v.IsMandatory && !v.Success)
	}
}

func TestHealthMandatoryFailure(t *testing.T) {
	defer clearHealthCheckCache()
	w := setupMuxWithConfigAndGetResponseForMethod(t,
		&Config{Endpoints: []int{Health}, Health: &HealthConfig{Checkers: []HealthChecker{{
			Key: "naruto",
			Func: func(_ context.Context) error {
				return nil
			},
		}, {
			Key: "another naruto",
			Func: func(_ context.Context) error {
				return errors.New("some error")
			},
			IsMandatory: true,
		}}}},
		http.MethodGet, healthEndpoint)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	var data map[string]HealthCheckInfo
	getTypedJSONBody(t, w.Body, &data)
	assert.Equal(t, 2, len(data))
}

func TestHealthNotConfigured(t *testing.T) {
	defer clearHealthCheckCache()
	w := setupMuxAndGetResponse(t, Info, healthEndpoint)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, notFoundError, w.Body.String())
}

func TestHealthInvalidMethod(t *testing.T) {
	defer clearHealthCheckCache()
	w := setupMuxWithConfigAndGetResponseForMethod(t,
		&Config{Endpoints: []int{Health}, Health: &HealthConfig{Checkers: []HealthChecker{{
			Key: "naruto",
			Func: func(_ context.Context) error {
				return nil
			},
		}}}},
		http.MethodHead, healthEndpoint)
	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
	assert.Equal(t, methodNotAllowedError, w.Body.String())
}

func TestHealthWithoutConfig(t *testing.T) {
	defer clearHealthCheckCache()
	w := setupMuxWithConfigAndGetResponseForMethod(t, nil, http.MethodGet, healthEndpoint)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, notFoundError, w.Body.String())
}

func TestHealthWithConfig(t *testing.T) {
	defer clearHealthCheckCache()
	w := setupMuxWithConfigAndGetResponseForMethod(t,
		&Config{Endpoints: []int{Health}, Health: &HealthConfig{Checkers: []HealthChecker{{
			Key: "naruto",
			Func: func(_ context.Context) error {
				return nil
			},
		}}}},
		http.MethodGet, healthEndpoint)
	assert.Equal(t, http.StatusOK, w.Code)

	var data map[string]HealthCheckInfo
	getTypedJSONBody(t, w.Body, &data)
	for _, v := range data {
		assert.False(t, v.IsMandatory && !v.Success)
	}
}

func TestHealthEncodeJSONError(t *testing.T) {
	mockEncodeJSONWithError()
	defer unMockEncodeJSON()

	defer clearHealthCheckCache()
	w := setupMuxWithConfigAndGetResponseForMethod(t,
		&Config{Endpoints: []int{Health}, Health: &HealthConfig{Checkers: []HealthChecker{{
			Key: "naruto",
			Func: func(_ context.Context) error {
				return nil
			},
		}}}},
		http.MethodGet, healthEndpoint)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "error", w.Body.String())
	assert.Equal(t, textStringContentType, w.Header().Get(contentTypeHeader))
}

func TestMetrics(t *testing.T) {
	w := setupMuxAndGetResponse(t, Metrics, metricsEndpoint)
	assert.Equal(t, http.StatusOK, w.Code)

	var data MetricsResponse
	getTypedJSONBody(t, w.Body, &data)
	assert.NotEmpty(t, data.MemStats)
}

func TestMetricsNotConfigured(t *testing.T) {
	w := setupMuxAndGetResponse(t, Info, metricsEndpoint)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, notFoundError, w.Body.String())
}

func TestMetricsInvalidMethod(t *testing.T) {
	w := setupMuxWithConfigAndGetResponseForMethod(t, &Config{Endpoints: []int{Metrics}}, http.MethodHead, metricsEndpoint)
	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
	assert.Equal(t, methodNotAllowedError, w.Body.String())
}

func TestMetricsWithoutConfig(t *testing.T) {
	w := setupMuxWithConfigAndGetResponseForMethod(t, nil, http.MethodGet, metricsEndpoint)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, notFoundError, w.Body.String())
}

func TestMetricsWithConfig(t *testing.T) {
	w := setupMuxWithConfigAndGetResponseForMethod(t, &Config{Endpoints: []int{Metrics}}, http.MethodGet, metricsEndpoint)
	assert.Equal(t, http.StatusOK, w.Code)

	var data MetricsResponse
	getTypedJSONBody(t, w.Body, &data)
	assert.NotEmpty(t, data.MemStats)
}

func TestMetricsEncodeJSONError(t *testing.T) {
	mockEncodeJSONWithError()
	defer unMockEncodeJSON()

	w := setupMuxAndGetResponse(t, Metrics, metricsEndpoint)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "error", w.Body.String())
	assert.Equal(t, textStringContentType, w.Header().Get(contentTypeHeader))
}

func TestPing(t *testing.T) {
	w := setupMuxAndGetResponse(t, Ping, pingEndpoint)
	assert.Equal(t, http.StatusOK, w.Code)

	assert.Empty(t, w.Body)
}

func TestPingNotConfigured(t *testing.T) {
	w := setupMuxAndGetResponse(t, Info, pingEndpoint)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, notFoundError, w.Body.String())
}

func TestPingInvalidMethod(t *testing.T) {
	w := setupMuxWithConfigAndGetResponseForMethod(t, &Config{Endpoints: []int{Ping}}, http.MethodHead, pingEndpoint)
	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
	assert.Equal(t, methodNotAllowedError, w.Body.String())
}

func TestPingWithoutConfig(t *testing.T) {
	w := setupMuxWithConfigAndGetResponseForMethod(t, nil, http.MethodGet, pingEndpoint)
	assert.Equal(t, http.StatusOK, w.Code)

	assert.Empty(t, w.Body)
}

func TestShutdown(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			// do nothing here, just to handle shutdown gracefully
		}
	}()
	w := setupMuxAndGetResponse(t, Shutdown, shutdownEndpoint)
	assert.Equal(t, http.StatusOK, w.Code)

	assert.Empty(t, w.Body)
}

func TestShutdownNotConfigured(t *testing.T) {
	w := setupMuxAndGetResponse(t, Info, shutdownEndpoint)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, notFoundError, w.Body.String())
}

func TestShutdownInvalidMethod(t *testing.T) {
	w := setupMuxWithConfigAndGetResponseForMethod(t, &Config{Endpoints: []int{Shutdown}}, http.MethodHead, shutdownEndpoint)
	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
	assert.Equal(t, methodNotAllowedError, w.Body.String())
}

func TestShutdownWithoutConfig(t *testing.T) {
	w := setupMuxWithConfigAndGetResponseForMethod(t, nil, http.MethodGet, shutdownEndpoint)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, notFoundError, w.Body.String())
}

func TestShutdownWithConfig(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			// do nothing here, just to handle shutdown gracefully
		}
	}()
	w := setupMuxWithConfigAndGetResponseForMethod(t, &Config{Endpoints: []int{Shutdown}}, http.MethodGet, shutdownEndpoint)
	assert.Equal(t, http.StatusOK, w.Code)

	assert.Empty(t, w.Body)
}

func TestThreadDump(t *testing.T) {
	w := setupMuxAndGetResponse(t, ThreadDump, threadDumpEndpoint)
	assert.Equal(t, http.StatusOK, w.Code)

	assert.NotEmpty(t, w.Body)
}

func TestThreadDumpNotConfigured(t *testing.T) {
	w := setupMuxAndGetResponse(t, Info, threadDumpEndpoint)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, notFoundError, w.Body.String())
}

func TestThreadDumpInvalidMethod(t *testing.T) {
	w := setupMuxWithConfigAndGetResponseForMethod(t, &Config{Endpoints: []int{ThreadDump}}, http.MethodHead, threadDumpEndpoint)
	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
	assert.Equal(t, methodNotAllowedError, w.Body.String())
}

func TestThreadDumpWithoutConfig(t *testing.T) {
	w := setupMuxWithConfigAndGetResponseForMethod(t, nil, http.MethodGet, threadDumpEndpoint)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, notFoundError, w.Body.String())
}

func TestThreadDumpWithConfig(t *testing.T) {
	w := setupMuxWithConfigAndGetResponseForMethod(t, &Config{Endpoints: []int{ThreadDump}}, http.MethodGet, threadDumpEndpoint)
	assert.Equal(t, http.StatusOK, w.Code)

	assert.NotEmpty(t, w.Body)
}

func TestThreadDumpWithError(t *testing.T) {
	mockPprofLookupWithError()
	defer unMockPprofLookup()

	w := setupMuxAndGetResponse(t, ThreadDump, threadDumpEndpoint)
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	assert.Equal(t, profileNotFoundError, w.Body.String())
}

func TestGetLastStringAfterDelim(t *testing.T) {
	assert.Equal(t, "", getLastStringAfterDelimiter("", slash))
	assert.Equal(t, "a", getLastStringAfterDelimiter("a", slash))
	assert.Equal(t, "", getLastStringAfterDelimiter("", ""))
	assert.Equal(t, "c", getLastStringAfterDelimiter("a/b/c", slash))
	assert.Equal(t, "", getLastStringAfterDelimiter("a/", slash))
}
