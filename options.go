package rek

import (
	"net/http"
	"time"
)

type Options struct {
	headers  map[string]string
	timeout  time.Duration
	username string
	password string
	data     map[string]interface{}
}

type Option func(*Options)

func WithHeaders(headers map[string]string) Option {
	return func(opts *Options) {
		opts.headers = headers
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(opts *Options) {
		opts.timeout = timeout
	}
}

func WithBasicAuth(username, password string) Option {
	return func(opts *Options) {
		opts.username = username
		opts.password = password
	}
}

func WithData(data map[string]interface{}) Option {
	return func(opts *Options) {
		opts.data = data
	}
}

func buildOptions(opts ...Option) *Options {
	os := &Options{
		headers: nil,
		timeout: 10 * time.Second,
	}

	for _, opt := range opts {
		opt(os)
	}

	return os
}

func setHeaders(req *http.Request, opts *Options) *http.Request {
	if opts.headers != nil {
		for k, v := range opts.headers {
			req.Header.Set(k, v)
		}
	}

	return req
}

func setBasicAuth(req *http.Request, opts *Options) *http.Request {
	if opts.username != "" && opts.password != "" {
		req.SetBasicAuth(opts.username, opts.password)
	}

	return req
}
