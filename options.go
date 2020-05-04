package rek

import "net/http"

type Options struct {
	Headers map[string]string
}

type Option func(*Options)

func WithHeaders(headers map[string]string) Option {
	return func(opts *Options) {
		opts.Headers = headers
	}
}

func buildOptions(opts ...Option) *Options {
	os := &Options{}

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
}