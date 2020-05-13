package rek

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestClientBuilder(t *testing.T) {
	var (
		timeout = 13 * time.Second
	)

	is := assert.New(t)

	opts := &options{}

	// No options
	cl := buildClient(opts)
	is.Equal(cl, &http.Client{})

	// Custom client
	client := &http.Client{
		Timeout: timeout,
	}
	opts.client = client
	cl = buildClient(opts)
	is.Equal(cl.Timeout, timeout)

	// OAuth
	opts = &options{}
	opts.oauth2Cfg = &oauth2Config{
		config: oauthCfg,
		token: oauthTok,
	}
	cl = buildClient(opts)
	is.Equal(cl, oauthCfg.Client(context.Background(), oauthTok))
}
