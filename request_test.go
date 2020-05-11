package rek

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestMaker(t *testing.T) {
	is := assert.New(t)

	method, url := http.MethodGet, "https://httpbin.org/get"

	basis, err := http.NewRequest(method, url, nil)
	is.NoError(err)
	is.NotNil(basis)

	req, err := buildRequest(method, url, &options{})
	is.NoError(err)
	is.NotNil(req)

	headers := map[string]string{"Foo": "bar"}
	opts := &options{headers: headers}
	req, err = buildRequest(method, url, opts)
	is.NoError(err)
	is.Equal(req.Header.Get("Foo"), "bar")
}
