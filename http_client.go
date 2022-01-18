package aoe

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type opts struct {
	timeout time.Duration
}

var defaultOpts = &opts{
	timeout: 30 * time.Second,
}

type optfunc func(o *opts)

// WithTimeout will allow a custom timeout to be set within the http client
// allowing the default of 30 seconds to be overwritten.
func WithTimeout(d time.Duration) optfunc {
	return func(o *opts) {
		o.timeout = d
	}
}

type HttpClient interface {
	Do(ctx context.Context, method, url string, body interface{}, out interface{}) error
}

type httpclient struct {
	o *opts
	c *http.Client
}

// NewHTTPClient will setup and return a new http client instance that can be
// used to send and validate http requests and responses.
// It's just a slightly simplifed wrapper over net.http.
func NewHTTPClient(opts ...optfunc) *httpclient {
	o := defaultOpts
	for _, opt := range opts {
		opt(o)
	}
	return &httpclient{
		o: o,
		c: &http.Client{
			Timeout: o.timeout,
		},
	}
}

// Do will send an http request using the provided parameters.
// Body and out are optional.
func (h httpclient) Do(ctx context.Context, method, url string, body interface{}, out interface{}) error {
	if method == "" || url == "" {
		return errors.New("method and url should be provided")
	}
	buf := &bytes.Buffer{}
	if body != nil {
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return err
		}
	}
	req, err := http.NewRequestWithContext(ctx, method, url, buf)
	if err != nil {
		return err
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("accept", "application/json")
	resp, err := h.c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 300 {
		return fmt.Errorf("unexpected status code %s returned from %s %s", resp.Status, strings.ToUpper(method), url)
	}
	if out == nil {
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(&out)
}
