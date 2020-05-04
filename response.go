package rek

import (
	"io/ioutil"
	"net/http"
)

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

func (r *Response) StatusCode() int {
	return r.statusCode
}

func (r *Response) Content() []byte {
	return r.content
}

func (r *Response) Headers() map[string]string {
	return r.headers
}

func (r *Response) Encoding() []string {
	return r.encoding
}

func (r *Response) Text() string {
	return string(r.content)
}

func (r *Response) Raw() *http.Response {
	return r.res
}

func (r *Response) Cookies() []*http.Cookie {
	return r.cookies
}

func (r *Response) ContentLength() int64 {
	return r.res.ContentLength
}

func (r *Response) Status() string {
	return r.res.Status
}
