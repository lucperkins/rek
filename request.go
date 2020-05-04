package rek

import (
	"io"
	"net/http"
)

func makeRequest(method, url string, opts *options) (*http.Request, error) {
	var body io.Reader

	if opts.data != nil {
		data, err := getData(opts)
		if err != nil {
			return nil, err
		}

		body = data
	}

	if opts.jsonObj != nil {
		js, err := getJson(opts)
		if err != nil {
			return nil, err
		}

		body = js
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	setHeaders(req, opts)

	setBasicAuth(req, opts)

	setCookies(req, opts)

	return req, nil
}
