package rek

import "net/http"

func makeClient(options *Options) *http.Client {
	c := &http.Client{}

	if options.timeout != 0 {
		c.Timeout = options.timeout
	}

	return c
}
