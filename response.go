package rek

import (
	"io/ioutil"
	"net/http"
)

type Response struct {
	statusCode int
	body       []byte
	headers    map[string]string
	encoding   []string
}

func makeResponse(res *http.Response) (*Response, error) {
	resp := &Response{
		statusCode: res.StatusCode,
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

		resp.body = bs
	}

	if res.TransferEncoding != nil {
		resp.encoding = res.TransferEncoding
	}

	return resp, nil
}

func (r *Response) StatusCode() int {
	return r.statusCode
}

func (r *Response) Body() []byte {
	return r.body
}

func (r *Response) Headers() map[string]string {
	return r.headers
}

func (r *Response) Encoding() []string {
	return r.encoding
}

func (r *Response) Text() string {
	return string(r.body)
}