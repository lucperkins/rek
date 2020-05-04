package rek

import "net/http"

func Get(url string, opts ...Option) (*Response, error) {
	return call(http.MethodGet, url, opts...)
}

func Post(url string, opts ...Option) (*Response, error) {
	return call(http.MethodPost, url, opts...)
}

func Put(url string, opts ...Option) (*Response, error) {
	return call(http.MethodPut, url, opts...)
}

func Delete(url string, opts ...Option) (*Response, error) {
	return call(http.MethodDelete, url, opts...)
}

func Patch(url string, opts ...Option) (*Response, error) {
	return call(http.MethodPatch, url, opts...)
}

func Head(url string, opts ...Option) (*Response, error) {
	return call(http.MethodHead, url, opts...)
}

func call(method, url string, opts ...Option) (*Response, error) {
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
