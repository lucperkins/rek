package rek

import (
	"net/http"
	"time"
)

type options struct {
	headers   map[string]string
	timeout   time.Duration
	username  string
	password  string
	data      map[string]interface{}
	userAgent string
	jsonObj   interface{}
	callback  func(*Response)
}

type Option func(*options)

func Headers(headers map[string]string) Option {
	return func(opts *options) {
		opts.headers = headers
	}
}

func Timeout(timeout time.Duration) Option {
	return func(opts *options) {
		opts.timeout = timeout
	}
}

func BasicAuth(username, password string) Option {
	return func(opts *options) {
		opts.username = username
		opts.password = password
	}
}

func Data(data map[string]interface{}) Option {
	return func(opts *options) {
		opts.data = data
	}
}

func UserAgent(agent string) Option {
	return func(opts *options) {
		opts.userAgent = agent
	}
}

func Struct(v interface{}) Option {
	return func(opts *options) {
		opts.jsonObj = v
	}
}

func Callback(cb func(*Response)) Option {
	return func(opts *options) {
		opts.callback = cb
	}
}

func buildOptions(opts ...Option) (*options, error) {
	os := &options{
		headers: nil,
		timeout: 10 * time.Second,
	}

	for _, opt := range opts {
		opt(os)
	}

	if os.jsonObj != nil && os.data != nil {
		return nil, ErrMultipleBodies
	}

	return os, nil
}

func setHeaders(req *http.Request, opts *options) *http.Request {
	if opts.headers != nil {
		for k, v := range opts.headers {
			req.Header.Set(k, v)
		}
	}

	if opts.userAgent != "" {
		req.Header.Set("User-Agent", opts.userAgent)
	}

	if opts.jsonObj != nil {
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	}

	if opts.data != nil {
		req.Header.Set("Content-Type", "application/octet-stream")
	}

	return req
}

func setBasicAuth(req *http.Request, opts *options) *http.Request {
	if opts.username != "" && opts.password != "" {
		req.SetBasicAuth(opts.username, opts.password)
	}

	return req
}
