package rek

import "net/http"

func makeClient(opts *options) *http.Client {
	c := &http.Client{}

	if opts.cookieJar != nil {
		c.Jar = *opts.cookieJar
	}

	if opts.timeout != 0 {
		c.Timeout = opts.timeout
	}

	return c
}
