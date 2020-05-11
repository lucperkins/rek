package rek

import (
	"context"
	"net/http"
)

func makeClient(opts *options) *http.Client {
	var cl *http.Client

	if opts.client != nil {
		cl = opts.client
	} else {
		if opts.oauth2Cfg != nil {
			cfg, tok := opts.oauth2Cfg.config, opts.oauth2Cfg.token

			cl = cfg.Client(getCtx(opts), tok)
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

			cl = c
		}
	}

	return cl
}

func getCtx(opts *options) context.Context {
	if opts.ctx == nil {
		return context.Background()
	} else {
		return opts.ctx
	}
}
