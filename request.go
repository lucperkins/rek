package rek

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"io"
	"net/http"
)

func makeRequest(method, url string, opts *options) (*http.Request, error) {
	var body io.Reader

	if opts.data != nil {
		var buf bytes.Buffer

		if err := gob.NewEncoder(&buf).Encode(opts.data); err != nil {
			return nil, err
		}

		body = &buf
	}

	if opts.jsonObj != nil {
		js, err := json.Marshal(opts.jsonObj)
		if err != nil {
			return nil, err
		}

		body = bytes.NewBuffer(js)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	setHeaders(req, opts)

	setBasicAuth(req, opts)

	return req, nil
}
