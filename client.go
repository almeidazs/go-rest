package gorest

import (
	"cmp"
	"errors"
	"net/http"
	"os"
	"time"
)

// Client is the AbacatePay REST API client.
type Client struct {
	opts Options
	http *http.Client
}

// New creates a new REST client with the given options.
func New(opts Options) *Client {
	if opts.Timeout == 0 {
		opts.Timeout = 5 * time.Second
	}

	if opts.Version == 0 {
		opts.Version = 1
	}

	if opts.Retry.Max == 0 {
		opts.Retry.Max = 3
	}

	return &Client{
		opts: opts,
		http: &http.Client{},
	}
}

// SetSecret sets the API secret used for authentication.
func (c *Client) SetSecret(secret string) {
	c.opts.Secret = secret
}

// secret returns the configured API secret or error if none is found.
func (c *Client) secret() (string, error) {
	secret := cmp.Or(c.opts.Secret, os.Getenv("ABACATEPAY_SECRET"))

	if secret != "" {
		return secret, nil
	}

	return "", errors.New("abacatepay: missing API secret")
}
