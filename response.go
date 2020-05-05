package rek

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// A struct containing the relevant response information returned by a rek request.
type Response struct {
	statusCode int
	content    []byte
	headers    map[string]string
	encoding   []string
	cookies    []*http.Cookie
	res        *http.Response
}

func makeResponse(res *http.Response) (*Response, error) {
	resp := &Response{
		statusCode: res.StatusCode,
		res:        res,
	}

	if res.Header != nil {
		headers := make(map[string]string)

		for k, v := range res.Header {
			headers[k] = v[0]
		}

		resp.headers = headers
	}

	if res.Body != nil {
		defer res.Body.Close()

		bs, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		resp.content = bs
	}

	if res.TransferEncoding != nil {
		resp.encoding = res.TransferEncoding
	}

	if res.Cookies() != nil {
		resp.cookies = res.Cookies()
	}

	return resp, nil
}

// The status code of the response (200, 404, etc.)
func (r *Response) StatusCode() int {
	return r.statusCode
}

// The response body as raw bytes.
func (r *Response) Content() []byte {
	return r.content
}

// The headers associated with the response.
func (r *Response) Headers() map[string]string {
	return r.headers
}

// The response's encoding.
func (r *Response) Encoding() []string {
	return r.encoding
}

// The response body as a string.
func (r *Response) Text() string {
	return string(r.content)
}

// Marshal a JSON response body.
func (r *Response) Json(v interface{}) error {
	return json.Unmarshal(r.content, v)
}

// The Content-Type header for the request (if any).
func (r *Response) ContentType() string {
	return r.Headers()["Content-Type"]
}

// The raw *http.Response returned by the operation.
func (r *Response) Raw() *http.Response {
	return r.res
}

// The cookies associated with the response.
func (r *Response) Cookies() []*http.Cookie {
	return r.cookies
}

// The content length of the response body.
func (r *Response) ContentLength() int64 {
	return r.res.ContentLength
}

// The response's status, e.g. "200 OK."
func (r *Response) Status() string {
	return r.res.Status
}
