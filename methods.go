package rek

import (
	"net/http"
	"net/url"
)

// GET request
func Get(url string, opts ...Option) (*Response, error) {
	return call(http.MethodGet, url, opts...)
}

// POST request
func Post(url string, opts ...Option) (*Response, error) {
	return call(http.MethodPost, url, opts...)
}

// PUT request
func Put(url string, opts ...Option) (*Response, error) {
	return call(http.MethodPut, url, opts...)
}

// DELETE request
func Delete(url string, opts ...Option) (*Response, error) {
	return call(http.MethodDelete, url, opts...)
}

// PATCH request
func Patch(url string, opts ...Option) (*Response, error) {
	return call(http.MethodPatch, url, opts...)
}

// HEAD request
func Head(url string, opts ...Option) (*Response, error) {
	options, err := buildOptions(opts...)
	if err != nil {
		return nil, err
	}

	cl := makeClient(options)

	res, err := cl.Head(url)
	if err != nil {
		return nil, err
	}

	return makeResponse(res)
}

// Make a request with an arbitrary HTTP method, i.e. not GET, POST, PUT, DELETE, etc.
func Do(method, url string, opts ...Option) (*Response, error) {
	return call(method, url, opts...)
}

func call(method, endpoint string, opts ...Option) (*Response, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	options, err := buildOptions(opts...)
	if err != nil {
		return nil, err
	}

	cl := makeClient(options)

	req, err := makeRequest(method, u.String(), options)

	res, err := cl.Do(req)
	if err != nil {
		return nil, err
	}

	resp, err := makeResponse(res)
	if err != nil {
		return nil, err
	}

	if options.callback != nil {
		options.callback(resp)
	}

	return resp, nil
}
