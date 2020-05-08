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
			var ctx context.Context

			cfg, tok := opts.oauth2Cfg.config, opts.oauth2Cfg.token

			if opts.ctx == nil {
				ctx = context.Background()
			} else {
				ctx = opts.ctx
			}

			cl = cfg.Client(ctx, tok)
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
