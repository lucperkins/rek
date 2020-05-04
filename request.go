package rek

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func makeRequest(method, url string, opts *Options) (*http.Request, error) {
	var body io.Reader

	if opts.data != nil {
		js, err := json.Marshal(opts.data)
		if err != nil {
			return nil, err
		}

		body = bytes.NewReader(js)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	setHeaders(req, opts)

	setBasicAuth(req, opts)

	return req, nil
}
