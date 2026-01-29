package gorest

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// shouldRetry reports whether another retry attempt is allowed.
func (r RetryOptions) shouldRetry(attempt int) bool {
	return attempt < r.Max
}

// delay returns the backoff delay for a given attempt.
func (r RetryOptions) delay(attempt int) time.Duration {
	if r.Backoff != nil {
		return r.Backoff(attempt)
	}

	return DefaultBackoff(attempt)
}

// retryTimeout handles retry logic when a request times out.
func (c *Client) retryTimeout(
	method, route string,
	query map[string]string,
	body any,
	attempt int,
) (any, error) {

	var zero any

	if !c.opts.Retry.shouldRetry(attempt) {
		return zero, &HTTPError{
			Message: "timeout after max retries",
			Route:   route,
			Status:  http.StatusServiceUnavailable,
			Method:  method,
		}
	}

	if c.opts.Retry.OnRetry != nil {
		c.opts.Retry.OnRetry(attempt)
	}

	time.Sleep(c.opts.Retry.delay(attempt))

	return c.do(method, route, query, body, attempt+1)
}

// retryHTTP handles retry logic for HTTP error responses.
func (c *Client) retryHTTP(
	resp *http.Response,
	method, route string,
	query map[string]string,
	body any,
	attempt int,
) (any, error) {

	var zero any

	if _, ok := RetryableStatus[resp.StatusCode]; !ok {
		var errPayload struct {
			Error AbacatePayError `json:"error"`
		}

		_ = json.NewDecoder(resp.Body).Decode(&errPayload)

		return zero, &errPayload.Error
	}

	if !c.opts.Retry.shouldRetry(attempt) {
		return zero, &HTTPError{
			Message: "max retries exceeded",
			Route:   route,
			Status:  resp.StatusCode,
			Method:  method,
		}
	}

	if resp.StatusCode == StatusRateLimit && c.opts.OnRateLimit != nil {
		c.opts.OnRateLimit(resp.StatusCode)
	}

	if c.opts.Retry.OnRetry != nil {
		c.opts.Retry.OnRetry(attempt)
	}

	time.Sleep(c.opts.Retry.delay(attempt))

	return c.do(method, route, query, body, attempt+1)
}

// newRequest builds an HTTP request with headers, body and context.
func (c *Client) newRequest(
	ctx context.Context,
	method, route string,
	query map[string]string,
	body any,
) (*http.Request, error) {
	secret, err := c.secret()

	if err != nil {
		return nil, err
	}

	var buf *bytes.Reader

	if body != nil {
		raw, _ := json.Marshal(body)

		buf = bytes.NewReader(raw)
	} else {
		buf = bytes.NewReader(nil)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.makeURL(route, query), buf)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+secret)

	for k, v := range c.opts.Headers {
		req.Header.Set(k, v)
	}

	return req, nil
}
