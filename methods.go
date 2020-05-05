package rek

import "net/http"

// GET request
func Get(url string, opts ...option) (*Response, error) {
	return call(http.MethodGet, url, opts...)
}

// POST request
func Post(url string, opts ...option) (*Response, error) {
	return call(http.MethodPost, url, opts...)
}

// PUT request
func Put(url string, opts ...option) (*Response, error) {
	return call(http.MethodPut, url, opts...)
}

// DELETE request
func Delete(url string, opts ...option) (*Response, error) {
	return call(http.MethodDelete, url, opts...)
}

// PATCH request
func Patch(url string, opts ...option) (*Response, error) {
	return call(http.MethodPatch, url, opts...)
}

// HEAD request
func Head(url string, opts ...option) (*Response, error) {
	options, err := buildOptions(opts...)

	cl := makeClient(options)

	res, err := cl.Head(url)
	if err != nil {
		return nil, err
	}

	return makeResponse(res)
}

func call(method, url string, opts ...option) (*Response, error) {
	options, err := buildOptions(opts...)
	if err != nil {
		return nil, err
	}

	cl := makeClient(options)

	req, err := makeRequest(method, url, options)

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
