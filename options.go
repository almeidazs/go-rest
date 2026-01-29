package gorest

import "time"

// RetryOptions controls how the client retries failed requests.
type RetryOptions struct {
	// Max is the maximum number of retry attempts.
	Max int

	// OnRetry is called before each retry attempt.
	OnRetry func(attempt int)

	// Backoff returns the delay before the next retry attempt.
	// If nil, a default exponential backoff is used.
	Backoff func(attempt int) time.Duration
}

// Options configures the REST client.
type Options struct {
	// BaseURL overrides the default API base URL.
	BaseURL string

	// Version specifies the API version (default: 1).
	Version int

	// Secret is the AbacatePay API secret.
	// If empty, the client will try to read ABACATEPAY_SECRET from the environment.
	Secret string

	// Headers are additional headers sent with every request.
	Retry RetryOptions

	// Timeout is the request timeout (default: 5s).
	Timeout time.Duration

	// OnRateLimit is called when a 429 response is received.
	OnRateLimit func(status int)

	// Retry configures retry behavior.
	Headers map[string]string
}
