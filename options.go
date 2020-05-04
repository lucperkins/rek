package rek

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type options struct {
	headers           map[string]string
	timeout           time.Duration
	username          string
	password          string
	data              interface{}
	userAgent         string
	jsonObj           interface{}
	callback          func(*Response)
	cookies           []*http.Cookie
	file              *file
	formData          map[string]string
	cookieJar         *http.CookieJar
	bearer            string
	disallowRedirects bool
}

func (o *options) validate() error {
	if o.jsonObj != nil && o.data != nil {
		return ErrMultipleBodies
	}

	return nil
}

type Option func(*options)

func Headers(headers map[string]string) Option {
	return func(opts *options) {
		opts.headers = headers
	}
}

func Timeout(timeout time.Duration) Option {
	return func(opts *options) {
		opts.timeout = timeout
	}
}

func BasicAuth(username, password string) Option {
	return func(opts *options) {
		opts.username = username
		opts.password = password
	}
}

func Data(data interface{}) Option {
	return func(opts *options) {
		opts.data = data
	}
}

func UserAgent(agent string) Option {
	return func(opts *options) {
		opts.userAgent = agent
	}
}

func Json(v interface{}) Option {
	return func(opts *options) {
		opts.jsonObj = v
	}
}

func Callback(cb func(*Response)) Option {
	return func(opts *options) {
		opts.callback = cb
	}
}

func Cookies(cookies []*http.Cookie) Option {
	return func(opts *options) {
		opts.cookies = cookies
	}
}

func CookieJar(jar http.CookieJar) Option {
	return func(opts *options) {
		opts.cookieJar = &jar
	}
}

func File(fieldName, filepath string, params map[string]string) Option {
	return func(opts *options) {
		opts.file = &file{
			FieldName: fieldName,
			Filepath:  filepath,
			Params:    params,
		}
	}
}

func FormData(form map[string]string) Option {
	return func(opts *options) {
		opts.formData = form
	}
}

func Bearer(bearer string) Option {
	return func(opts *options) {
		opts.bearer = bearer
	}
}

func DisallowRedirects() Option {
	return func(opts *options) {
		opts.disallowRedirects = true
	}
}

func buildOptions(opts ...Option) (*options, error) {
	os := &options{
		headers: nil,
		timeout: 10 * time.Second,
	}

	for _, opt := range opts {
		opt(os)
	}

	if err := os.validate(); err != nil {
		return nil, err
	}

	return os, nil
}

func setHeaders(req *http.Request, opts *options) *http.Request {
	if opts.headers != nil {
		for k, v := range opts.headers {
			req.Header.Set(k, v)
		}
	}

	if opts.userAgent != "" {
		req.Header.Set("User-Agent", opts.userAgent)
	}

	if opts.jsonObj != nil {
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	}

	if opts.data != nil {
		req.Header.Set("Content-Type", "application/octet-stream")
	}

	if opts.jsonObj != nil {
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	}

	if opts.formData != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	return req
}

func setBasicAuth(req *http.Request, opts *options) {
	if opts.username != "" && opts.password != "" {
		req.SetBasicAuth(opts.username, opts.password)
	}
}

func setCookies(req *http.Request, opts *options) {
	if opts.cookies != nil {
		for _, c := range opts.cookies {
			req.AddCookie(c)
		}
	}
}

func getData(opts *options) (io.Reader, error) {
	var buf bytes.Buffer

	if err := gob.NewEncoder(&buf).Encode(opts.data); err != nil {
		return nil, err
	}

	return &buf, nil
}

func getJson(opts *options) (io.Reader, error) {
	js, err := json.Marshal(opts.jsonObj)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(js), nil
}
