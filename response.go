package rek

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

// A struct containing the relevant response information returned by a rek request.
type Response struct {
	statusCode int
	headers    map[string]string
	encoding   []string
	cookies    []*http.Cookie
	res        *http.Response
	body       io.ReadCloser
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
		resp.body = res.Body
	}

	if res.TransferEncoding != nil {
		resp.encoding = res.TransferEncoding
	}

	if res.Cookies() != nil {
		resp.cookies = res.Cookies()
	}

	return resp, nil
}

// The status code of the response (200, 404, etc.).
func (r *Response) StatusCode() int {
	return r.statusCode
}

// The response body as a io.ReadCloser. Bear in mind that the response body can only be read once.
func (r *Response) Body() io.ReadCloser {
	return r.body
}

// The response body as a byte slice. Bear in mind that the response body can only be read once.
func (r *Response) BodyAsBytes() ([]byte, error) {
	return bodyBytes(r.body)
}

// The response body as a string. Bear in mind that the response body can only be read once.
func (r *Response) BodyAsString() (string, error) {
	bs, err := bodyBytes(r.body)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func bodyBytes(r io.ReadCloser) ([]byte, error) {
	buf := bytes.Buffer{}
	if _, err := buf.ReadFrom(r); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// The headers associated with the response.
func (r *Response) Headers() map[string]string {
	return r.headers
}

// The response's encoding.
func (r *Response) Encoding() []string {
	return r.encoding
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
