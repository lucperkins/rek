package rek

import (
	"net/http"
	"time"
)

type Options struct {
	Headers  map[string]string
	Timeout  time.Duration
	Username string
	Password string
}

type Option func(*Options)

func WithHeaders(headers map[string]string) Option {
	return func(opts *Options) {
		opts.Headers = headers
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(opts *Options) {
		opts.Timeout = timeout
	}
}

func WithBasicAuth(username, password string) Option {
	return func(opts *Options) {
		opts.Username = username
		opts.Password = password
	}
}

func buildOptions(opts ...Option) *Options {
	os := &Options{
		Headers: nil,
		Timeout: 10 * time.Second,
	}

	for _, opt := range opts {
		opt(os)
	}

	return os
}

func setHeaders(req *http.Request, opts *Options) *http.Request {
	if opts.Headers != nil {
		for k, v := range opts.Headers {
			req.Header.Set(k, v)
		}
	}

	return req
}
