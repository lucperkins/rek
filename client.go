package rek

import "net/http"

func makeClient(opts *options) *http.Client {
	if opts.client != nil {
		return opts.client
	} else {
		c := &http.Client{}

		if opts.cookieJar != nil {
			c.Jar = *opts.cookieJar
		}

		if opts.timeout != 0 {
			c.Timeout = opts.timeout
		}

		if opts.disallowRedirects {
			c.CheckRedirect = func(_ *http.Request, _ []*http.Request) error {
				return http.ErrUseLastResponse
			}
		}

		return c
	}
}
