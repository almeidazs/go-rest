package gorest

import (
	"cmp"
	"net/url"
)

// makeURL builds the full request URL.
func (c *Client) makeURL(route string, query map[string]string) string {
	base := cmp.Or(c.opts.BaseURL, "https://api.abacatepay.com/v")

	u := base + string(rune('0'+c.opts.Version)) + route

	if len(query) == 0 {
		return u
	}

	q := url.Values{}

	for k, v := range query {
		q.Set(k, v)
	}

	return u + "?" + q.Encode()
}
