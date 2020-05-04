package rek

import "net/http"

func Get(url string, opts ...Option) (*http.Response, error) {
	options := buildOptions(opts...)

	cl := makeClient(options)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	setHeaders(req, options)

	return cl.Do(req)
}
