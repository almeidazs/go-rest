package gorest

import "fmt"

// HTTPError represents a low-level HTTP error returned after retries are exhausted
// or when a non-retryable status code is received.
type HTTPError struct {
	Message string
	Route   string
	Status  int
	Method  string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("[%s %s] %d: %s", e.Method, e.Route, e.Status, e.Message)
}

// AbacatePayError represents a business error returned by the AbacatePay API.
type AbacatePayError struct {
	Message string `json:"message"`
}

func (e *AbacatePayError) Error() string {
	return e.Message
}
