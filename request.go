package rek

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func makeRequest(method, endpoint string, opts *options) (*http.Request, error) {
	var body io.Reader
	var contentType string

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

	if opts.file != nil {
		b, ct, err := buildMultipartBody(opts)
		if err != nil {
			return nil, err
		}

		contentType = ct
		body = b
	}

	if opts.formData != nil {
		vals := url.Values{}

		for k, v := range opts.formData {
			vals.Set(k, v)
		}

		body = strings.NewReader(vals.Encode())
	}

	req, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		return nil, err
	}

	if opts.formData != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	if opts.file != nil {
		req.Header.Set("Content-Type", contentType)
	}

	if opts.jwt != "" {
		bearer := fmt.Sprintf("Bearer %s", opts.jwt)
		req.Header.Set("Authorization", bearer)
	}

	setHeaders(req, opts)

	setBasicAuth(req, opts)

	setCookies(req, opts)

	return req, nil
}
