package gorest

import (
	"net/http"
	"time"
)

// Common HTTP status codes used by the SDK.
const (
	StatusRateLimit = 429
)

// RetryableStatus defines which HTTP status codes are eligible for retry.
var RetryableStatus = map[int]struct{}{
	http.StatusRequestTimeout:      {},
	StatusRateLimit:                {},
	http.StatusInternalServerError: {},
	http.StatusBadGateway:          {},
	http.StatusServiceUnavailable:  {},
	http.StatusGatewayTimeout:      {},
}

// DefaultBackoff implements an exponential backoff strategy.
func DefaultBackoff(attempt int) time.Duration {
	return time.Duration(1<<attempt) * time.Second
}
