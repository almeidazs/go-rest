package gorest

import (
	"context"
	"encoding/json"
	"net/http"
)


func (c *Client) do(
	method string,
	route string,
	query map[string]string,
	body any,
	attempt int,
) (any, error) {
	var zero any

	ctx, cancel := context.WithTimeout(context.Background(), c.opts.Timeout)
	defer cancel()

	req, err := c.newRequest(ctx, method, route, query, body)
	if err != nil {
		return zero, err
	}

	resp, err := c.http.Do(req)

	if err != nil {
		return c.retryTimeout(method, route, query, body, attempt)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return zero, nil
	}

	if resp.StatusCode >= 400 {
		return c.retryHTTP(resp, method, route, query, body, attempt)
	}

	var payload struct {
		Data  any                `json:"data"`
		Error *AbacatePayError `json:"error"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return zero, err
	}

	if payload.Error != nil {
		return zero, payload.Error
	}

	return payload.Data, nil
}

func (c *Client) Get(route string, query map[string]string) (any, error) {
	return c.do(http.MethodGet, route, query, nil, 0)
}

func (c *Client) Post(route string, body any) (any, error) {
	return c.do(http.MethodPost, route, nil, body, 0)
}

func (c *Client) Put(route string, body any) (any, error) {
	return c.do(http.MethodPut, route, nil, body, 0)
}

func (c *Client) Patch(route string, body any) (any, error) {
	return c.do(http.MethodPatch, route, nil, body, 0)
}

func (c *Client) Delete(route string) (any, error) {
	return c.do(http.MethodDelete, route, nil, nil, 0)
}

