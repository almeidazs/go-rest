package gorest

import (
	"net/http"
	"time"
)

const (
	StatusRateLimit      = 429
)

var RetryableStatus = map[int]struct{}{
	http.StatusRequestTimeout: {},
	StatusRateLimit: {},
	http.StatusInternalServerError: {},
	http.StatusBadGateway: {},
	http.StatusServiceUnavailable: {},
	http.StatusGatewayTimeout: {},
}

func DefaultBackoff(attempt int) time.Duration {
	return time.Duration(1<<attempt) * time.Second
}
