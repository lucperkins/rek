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
	i := 0

	if o.jsonObj != nil {
		i++
	}

	if o.data != nil {
		i++
	}

	if o.formData != nil {
		i++
	}

	// Throw an error if the request body has been set more than once.
	if i > 1 {
		return ErrRequestBodySetMultipleTimes
	}

	return nil
}

type option func(*options)

// Add headers to the request.
func Headers(headers map[string]string) option {
	return func(opts *options) {
		opts.headers = headers
	}
}

// Add a timeout to the request.
func Timeout(timeout time.Duration) option {
	return func(opts *options) {
		opts.timeout = timeout
	}
}

// Add a basic auth username and password to the request.
func BasicAuth(username, password string) option {
	return func(opts *options) {
		opts.username = username
		opts.password = password
	}
}

// Add any interface{} that can be serialized to a []byte and apply a "Content-Type: application/octet-stream" header.
func Data(data interface{}) option {
	return func(opts *options) {
		opts.data = data
	}
}

// Add a User-Agent header to the request.
func UserAgent(agent string) option {
	return func(opts *options) {
		opts.userAgent = agent
	}
}

// Add any interface{} that can be marshaled as JSON to the request body and apply a "Content-Type:
// application/json;charset=utf-8" header.
func Json(v interface{}) option {
	return func(opts *options) {
		opts.jsonObj = v
	}
}

// Add a callback function for handling the Response.
func Callback(cb func(*Response)) option {
	return func(opts *options) {
		opts.callback = cb
	}
}

// Add cookies to the request.
func Cookies(cookies []*http.Cookie) option {
	return func(opts *options) {
		opts.cookies = cookies
	}
}

// Add a cookie jar to the request.
func CookieJar(jar http.CookieJar) option {
	return func(opts *options) {
		opts.cookieJar = &jar
	}
}

// Create a multipart file upload request.
func File(fieldName, filepath string, params map[string]string) option {
	return func(opts *options) {
		opts.file = &file{
			FieldName: fieldName,
			Filepath:  filepath,
			Params:    params,
		}
	}
}

// Add key/value form data to the request body and apply a "Content-Type: application/x-www-form-urlencoded" header.
func FormData(form map[string]string) option {
	return func(opts *options) {
		opts.formData = form
	}
}

// Add a bearer header of the form "Authorization: Bearer ..."
func Bearer(bearer string) option {
	return func(opts *options) {
		opts.bearer = bearer
	}
}

// Turn redirects off.
func DisallowRedirects() option {
	return func(opts *options) {
		opts.disallowRedirects = true
	}
}

func buildOptions(opts ...option) (*options, error) {
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
