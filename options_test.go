package rek

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

var (
	headers  = map[string]string{"Foo": "bar"}
	timeout  = 11 * time.Hour
	username = "user"
	password = "pass"
	data     = struct {
		Name string
		Age  int
	}{"Methuselah", 965}
	userAgent   = "Feldman"
	callback    = func(res *Response) { fmt.Println(res.StatusCode()) }
	cookies     = []*http.Cookie{{Path: "/foo"}}
	f           = &file{FieldName: "file", Filepath: "go.mod", Params: nil}
	form        = map[string]string{"foo": "bar", "baq": "baz"}
	bearer      = "a1b2c3d4"
	reqModifier = func(req *http.Request) { fmt.Println(req.URL.String()) }
	apiKey      = bearer
	ctx         = context.WithValue(context.Background(), "value", "some-val")
	cl          = &http.Client{}
	oauthCfg    = &oauth2.Config{Scopes: []string{"admin", "superuser"}}
	oauthTok    = &oauth2.Token{AccessToken: "a1b2c3d4"}
)

func TestOptionsBuilder(t *testing.T) {
	is := assert.New(t)

	opts, err := buildOptions()
	is.NoError(err)
	is.NotNil(opts)

	is.Equal(opts.headers, map[string]string(nil))
	is.Zero(opts.timeout)
	is.Empty(opts.username, opts.password, opts.userAgent, opts.bearer, opts.apiKey, opts.ctx, opts.client,
		opts.oauth2Cfg)
	is.Nil(opts.data, opts.jsonObj, opts.callback, opts.cookies, opts.cookieJar, opts.file, opts.formData,
		opts.reqModifier)
	is.False(opts.disallowRedirects)

	// Headers
	opts, err = buildOptions(Headers(headers))
	is.NoError(err)
	is.Equal(opts.headers, headers)

	// Timeout
	opts, err = buildOptions(Timeout(timeout))
	is.NoError(err)
	is.Equal(opts.timeout, timeout)

	// Basic auth
	opts, err = buildOptions(BasicAuth(username, password))
	is.NoError(err)
	is.Equal(opts.username, username)
	is.Equal(opts.password, password)

	// Data
	opts, err = buildOptions(Data(data))
	is.NoError(err)
	is.Equal(opts.data, data)

	// User agent
	opts, err = buildOptions(UserAgent(userAgent))
	is.NoError(err)
	is.Equal(opts.userAgent, userAgent)

	// JSON
	opts, err = buildOptions(Json(data))
	is.NoError(err)
	is.Equal(opts.jsonObj, data)

	// Callback
	opts, err = buildOptions(Callback(callback))
	is.NoError(err)
	// We can't compare functions for equality, so NotNil will have to do
	is.NotNil(opts.callback)

	// Cookies
	opts, err = buildOptions(Cookies(cookies))
	is.NoError(err)
	is.Equal(opts.cookies, cookies)

	// Cookie jar
	// TODO

	// File
	opts, err = buildOptions(File(f.FieldName, f.Filepath, f.Params))
	is.NoError(err)
	is.Equal(opts.file, f)

	// Form data
	opts, err = buildOptions(FormData(form))
	is.NoError(err)
	is.Equal(opts.formData, form)

	// Bearer
	opts, err = buildOptions(Bearer(bearer))
	is.NoError(err)
	is.Equal(opts.bearer, bearer)

	// Disallow redirects
	opts, err = buildOptions(DisallowRedirects())
	is.NoError(err)
	is.True(opts.disallowRedirects)

	// Request modifier
	opts, err = buildOptions(RequestModifier(reqModifier))
	is.NoError(err)
	// We can't compare functions for equality, so NotNil will have to do
	is.NotNil(opts.reqModifier)

	// API key
	opts, err = buildOptions(ApiKey(apiKey))
	is.NoError(err)
	is.Equal(opts.apiKey, apiKey)

	// Context
	opts, err = buildOptions(Context(ctx))
	is.NoError(err)
	is.Equal(opts.ctx, ctx)

	// Client
	opts, err = buildOptions(Client(cl))
	is.NoError(err)
	is.Equal(opts.client, cl)

	// OAuth2
	opts, err = buildOptions(OAuth2(oauthCfg, oauthTok))
	is.NoError(err)
	is.Equal(opts.oauth2Cfg.config, oauthCfg)
	is.Equal(opts.oauth2Cfg.token, oauthTok)
}
