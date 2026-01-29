package gorest

import "fmt"

type HTTPError struct {
	Message string
	Route   string
	Status  int
	Method  string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("[%s %s] %d: %s", e.Method, e.Route, e.Status, e.Message)
}

type AbacatePayError struct {
	Message string `json:"message"`
}

func (e *AbacatePayError) Error() string {
	return e.Message
}
