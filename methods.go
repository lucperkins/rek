package rek

import "net/http"

func Get(url string, opts ...Option) (*Response, error) {
	options := buildOptions(opts...)

	cl := makeClient(options)

	req, err := makeRequest(http.MethodGet, url, options)

	res, err := cl.Do(req)
	if err != nil {
		return nil, err
	}

	return makeResponse(res)
}
