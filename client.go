package rek

import "net/http"

func makeClient(options *Options) *http.Client {
	c := &http.Client{}

	if options.Timeout != 0 {
		c.Timeout = options.Timeout
	}

	return c
}
