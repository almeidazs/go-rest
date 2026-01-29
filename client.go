package gorest

import (
	"cmp"
	"errors"
	"net/http"
	"os"
	"time"
)

type Client struct {
	opts  Options
	http  *http.Client
}

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

func (c *Client) SetSecret(secret string) {
	c.opts.Secret = secret
}

func (c *Client) secret() (string, error) {
	secret := cmp.Or(c.opts.Secret, os.Getenv("ABACATEPAY_SECRET"))

	if secret != "" {
		return secret, nil
	}

	return "", errors.New("abacatepay: missing API secret")
}
