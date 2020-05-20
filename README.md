# rek

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=gostyle=flat)](https://pkg.go.dev/mod/github.com/lucperkins/rek)

An easy HTTP client for [Go](https://golang.org) inspired by [Requests](https://requests.readthedocs.io/en/master/), plus all the Go-specific goodies you'd hope for in a client. Here's an example:

```go
// GET request
res, _ := rek.Get("https://httpbin.org/get")

fmt.Println(res.StatusCode())

body, _ := rek.BodyAsString(res.Body())

fmt.Println(body)

// POST request
type Comment struct {
    Username string `json:"username"`
    Body     string `json:"body"`
}

res, _ := rek.Post("https://httpbin.org/post",
    rek.Json(Comment{Username: "genesiskel", Body: "This movie sucked. Two thumbs down."}),
    rek.Headers(map[string]string{"My-Custom-Header": "foo,bar,baz"}),
    rek.BasicAuth("user", "pass"),
    rek.Timeout(5 * time.Second),
)

fmt.Println(res.StatusCode())

body, _ := rek.BodyAsString(res.Body())
fmt.Println(body)
```

## Responses

The `Response` struct has the following methods:


Method | Description | Return type
:------|:------------|:-----------
`StatusCode()` | HTTP status code, e.g. 200, 400 | `int`
`Body()` | The HTTP response body as a reader | [`io.ReadCloser`](https://pkg.go.dev/io?tab=doc#ReadCloser)
`Headers()` | The response headers | `map[string]string`
`Encoding()` | The content encoding of the response body, e.g. `gzip` | `[]string`
`ContentType()` | The value of the `Content-Type` header | `string`
`Raw()` | The unmodified [`*http.Response`](https://pkg.go.dev/net/http?tab=doc#Response) | [`*http.Response`](https://pkg.go.dev/net/http?tab=doc#Response)
`Cookies()` | The cookies attached to the response | `[]*http.Cookie`
`ContentLength()` | The length of the response | `int64`
`Status()` | The status of the response, e.g. `200 OK` | `string`

### HTTP response body

Keep in mind that the HTTP response body **can only be read once**. This is one area in which rek does *not* directly correspond to the [Requests](https://requests.readthedocs.io/en/master/) API. And so this, for example, won't work the way you might want:

```go
res, _ := rek.Get("https://httpbin.org/get")

bs1, _ := ioutil.ReadAll(res.Body()) // Non-empty byte slice

bs2, _ := ioutil.ReadAll(res.Body()) // Empty byte slice
```

If you'd like to use the response body more than once, store it in a variable rather than re-accessing the body.

### Helper methods

There are two simple helper methods for working with the response body:

Function | Return types
:--------|:------------
`BodyAsString(io.ReadCloser)` | `(string, error)`
`BodyAsBytes(io.ReadCloser)` | `([]byte, error)`

Bear in mind the caveat mentioned above, that the request body can only be read once, still holds. Here are some examples:

```go
res, _ := rek.Get("https://httpbin.org/get")

s1, _ := rek.BodyAsString(res.Body()) // body is read here

s2, _ := rek.BodyAsString(res.Body()) // s2 is an empty string
```

## Options

### Headers

```go
headers := map[string]string{
    "My-Custom-Header": "foo,bar,baz",
}

res, err := rek.Post("https://httpbin.org/post", rek.Headers(headers))
```

### JSON

Pass in any struct:

```go
type Comment struct {
    ID        int64     `json:"id"`
    Body      string    `json:"body"`
    Timestamp time.Time `json:"timestamp"`
}

comment := Comment{ID:47, Body:"Cool movie!", Timestamp: time.Now()}

res, err := rek.Post("https://httpbin.org/post", rek.Json(comment))
```

> Request headers are automatically updated to include `Content-Type` as `application/json;charset=utf-8`.

### Request timeout

```go
res, err := rek.Get("https://httpbin.org/get", rek.Timeout(5 * time.Second))
```

### Form data

```go
form := map[string]string{
    "foo": "bar",
    "baq": "baz",
}

res, err := rek.Put("https://httpbin.org/put", rek.FormData(form))
```

> Request headers are automatically updated to include `Content-Type` as `application/x-www-form-urlencoded`.

### File upload

```go
fieldName := "file"
filepath := "docs/README.md"
params := nil

res, err := rek.Post("https:/httpbin.org/post", rek.File(fieldName, filepath, params))
```

> Request headers are automatically updated to include `Content-Type` as `multipart/form-data; boundary=...`.

### Basic auth

```go
username, password := "user", "pass"

res, _ := rek.Get(fmt.Sprintf("https://httpbin.org/basic-auth/%s/%s", username, password),
    rek.BasicAuth(username, password))

fmt.Println(res.StatusCode()) // 200

res, _ := rek.Get(fmt.Sprintf("https://httpbin.org/basic-auth/%s/other", username, password),
    rek.BasicAuth(username, password))

fmt.Println(res.StatusCode()) // 401
```

### Data

Takes any input and serializes it to a `[]byte`:

```go
data := map[string]interface{
    "age": 38,
    "name": "Luc",
}

res, err := rek.Post("https://httpbin.org/post", rek.Data(data))
```

> Request headers are automatically updated to include `Content-Type` as `application/octet-stream`.

### User agent

```go
res, err := rek.Post("https://httpbin.org/post", rek.UserAgent("ThisGuy"))
```

### Bearer (useful for JSON Web Tokens)

```go
token := "... token ..."

res, err := rek.Post("https://httpbin.org/post", rek.Bearer(token))
```

### Request modifier

Supply a function that modifies the [`*http.Request`](https://pkg.go.dev/net/http?tab=doc#Request) (after all other supplied options have been applied to the request):

```go
modifier := func(r *http.Request) {
   // Do whatever you want with the request
}

res, err := rek.Get("https://httpbin.org/get", rek.RequestModifier(modifier))
```

### Accept

Apply an `Accept` header to the request:

```go
res, err := rek.Get("https://httpbin.org/get", rek.Accept("application/tar+gzip"))
```

### API key

Add an API key to the request as an `Authorization` header (where the value is `Basic ${KEY}`):

```go
res, err := rek.Get("https://some-secure-api.io", rek.ApiKey("a1b2c3..."))
```

### Context

Supply a [`Context`](https://pkg.go.dev/context?tab=doc#Context) to the request:

```go
ctx, cancel := context.WithCancel(context.Background())

res, err := rek.Get("https://long-winded-api.io", rek.Context(ctx))

// Ugh, I don't want this request to happen anymore

cancel()
```

### OAuth2

You can add an OAuth2 configuration and token to your request:

```go
conf := &oauth2.Config{
	ClientID:     "YOUR_CLIENT_ID",
	ClientSecret: "YOUR_CLIENT_SECRET",
	Scopes:       []string{"SCOPE1", "SCOPE2"},
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://provider.com/o/oauth2/auth",
		TokenURL: "https://provider.com/o/oauth2/token",
	},
}

tok, err := conf.Exchange(ctx, code)
// handle error

res, err := rek.Get("https://oauth2-protected-site.com", rek.OAuth2(conf, tok))
```

## Custom client

You can pass in your own `*http.Client` as the basis for the request:

```go
client := &http.Client{
    // Custom attributes
}

res, err := rek.Get("https://httpbin.org/get", rek.Client(client))
```


## Validation

It's important to bear in mind that rek provides *no validation* for the options that you provide on a specific request and doesn't provide any constraints on which options can be used with which request method. Some options may not make sense for some methods, e.g. request JSON on a `HEAD` request, but I leave it up to the end user to supply their own constraints. One exception is that the request body can only be set once. If you attempt to set it more than once you'll get a `ErrRequestBodySetMultipleTimes` error. This, for example, will throw that error:

```go
comment := Comment{Body: "This movie sucked"}

_, err := rek.Post("https://httpbin.org/post",
    rek.Json(comment),
    rek.FormData(map[string]string{"foo": "bar"}))

fmt.Println(errors.Is(err, rek.ErrRequestBodySetMultipleTimes)) // true
```
