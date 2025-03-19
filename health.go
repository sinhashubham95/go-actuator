package actuator

import (
	"context"
	"net/http"
	"sync"
	"time"
)

// HealthCheckInfo is used as the health check information.
type HealthCheckInfo struct {
	Key         string `json:"key"`
	IsMandatory bool   `json:"isMandatory"`
	Success     bool   `json:"success"`
	Error       string `json:"error,omitempty"`
}

var healthCheckInfoLock sync.RWMutex
var lastHealthCheckStamp time.Time
var healthCheckInfo = make(map[string]HealthCheckInfo)

func getClonedHealthCheckInfo(info map[string]HealthCheckInfo) map[string]HealthCheckInfo {
	result := make(map[string]HealthCheckInfo)
	for k, v := range info {
		result[k] = HealthCheckInfo{
			Key:         v.Key,
			IsMandatory: v.IsMandatory,
			Success:     v.Success,
			Error:       v.Error,
		}
	}
	return result
}

func getHealthCheckInfoFromCacheIfValid(config *Config) (bool, map[string]HealthCheckInfo) {
	healthCheckInfoLock.RLock()
	defer healthCheckInfoLock.RUnlock()
	if lastHealthCheckStamp.Add(config.Health.CacheDuration).Before(time.Now()) {
		return false, nil
	}
	return true, getClonedHealthCheckInfo(healthCheckInfo)
}

func getHealthCheckInfo(ctx context.Context, c HealthChecker, ch chan HealthCheckInfo) {
	h := HealthCheckInfo{
		Key:         c.Key,
		IsMandatory: c.IsMandatory,
	}
	err := c.Func(ctx)
	if err != nil {
		h.Error = err.Error()
	} else {
		h.Success = true
	}
	ch <- h
}

func getHealthCheckInfoAndCacheIfSuccess(config *Config) (ok bool, result map[string]HealthCheckInfo) {
	result = make(map[string]HealthCheckInfo)
	ctx, cancel := context.WithTimeout(context.Background(), config.Health.Timeout)
	defer cancel()
	ch := make(chan HealthCheckInfo, len(config.Health.Checkers))
	for _, c := range config.Health.Checkers {
		go getHealthCheckInfo(ctx, c, ch)
	}
	ok = true
	for range config.Health.Checkers {
		h := <-ch
		result[h.Key] = h
		if !h.Success && h.IsMandatory {
			ok = false
		}
	}
	healthCheckInfoLock.Lock()
	defer healthCheckInfoLock.Unlock()
	lastHealthCheckStamp = time.Now()
	healthCheckInfo = getClonedHealthCheckInfo(result)
	return
}

func getHealthCheck(config *Config) (ok bool, result map[string]HealthCheckInfo) {
	ok, result = getHealthCheckInfoFromCacheIfValid(config)
	if ok {
		return
	}
	ok, result = getHealthCheckInfoAndCacheIfSuccess(config)
	return
}

func getHealthHandler(config *Config) http.HandlerFunc {
	return func(writer http.ResponseWriter, _ *http.Request) {
		ok, result := getHealthCheck(config)
		writer.Header().Add(contentTypeHeader, applicationJSONContentType)
		b, err := encodeJSON(result)
		if err != nil || !ok {
			writer.WriteHeader(http.StatusInternalServerError)
		} else {
			writer.WriteHeader(http.StatusOK)
		}
		_, _ = writer.Write(b)
	}
}
