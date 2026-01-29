package gorest

import "time"

type RetryOptions struct {
	Max     int
	OnRetry func(attempt int)
	Backoff func(attempt int) time.Duration
}

type Options struct {
	BaseURL     string
	Version     int
	Secret      string
	Retry       RetryOptions
	Timeout     time.Duration
	Headers     map[string]string
	OnRateLimit func(status int)
}
