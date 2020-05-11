package rek

import (
	"net/http"
	"net/url"
)

// GET request
func Get(url string, opts ...Option) (*Response, error) {
	return do(http.MethodGet, url, opts...)
}

// POST request
func Post(url string, opts ...Option) (*Response, error) {
	return do(http.MethodPost, url, opts...)
}

// PUT request
func Put(url string, opts ...Option) (*Response, error) {
	return do(http.MethodPut, url, opts...)
}

// DELETE request
func Delete(url string, opts ...Option) (*Response, error) {
	return do(http.MethodDelete, url, opts...)
}

// PATCH request
func Patch(url string, opts ...Option) (*Response, error) {
	return do(http.MethodPatch, url, opts...)
}

// HEAD request
func Head(url string, opts ...Option) (*Response, error) {
	options, err := buildOptions(opts...)
	if err != nil {
		return nil, err
	}

	cl := buildClient(options)

	res, err := cl.Head(url)
	if err != nil {
		return nil, err
	}

	return buildResponse(res)
}

// Make a request with an arbitrary HTTP method, i.e. not GET, POST, PUT, DELETE, etc.
func Do(method, url string, opts ...Option) (*Response, error) {
	return do(method, url, opts...)
}

func do(method, endpoint string, opts ...Option) (*Response, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	options, err := buildOptions(opts...)
	if err != nil {
		return nil, err
	}

	cl := buildClient(options)

	req, err := buildRequest(method, u.String(), options)

	res, err := cl.Do(req)
	if err != nil {
		return nil, err
	}

	resp, err := buildResponse(res)
	if err != nil {
		return nil, err
	}

	if options.callback != nil {
		options.callback(resp)
	}

	return resp, nil
}
