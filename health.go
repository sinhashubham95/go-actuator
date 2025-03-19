package actuator

import "net/http"

func getHealthHandler(_ *Config) http.HandlerFunc {
	return func(_ http.ResponseWriter, _ *http.Request) {}
}
