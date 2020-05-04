package rek

import "net/http"

func makeRequest(method, url string, opts *Options) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	setHeaders(req, opts)

	return req, nil
}
