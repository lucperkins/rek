# rek

An easy HTTP client for [Go](https://golang.org), inspired by [Requests](https://requests.readthedocs.io/en/master/). Here's an example:

```go
type Comment struct {
    Body string `json:"body"`
}

comment := Comment{Body: "This movie sucked"}

headers := map[string]string{"My-Custom-Header", "foo,bar,baz"}

res, err := rek.Post("https://httpbin.org/post",
    rek.Json(comment),
    rek.Headers(headers),
    rek.BasicAuth("user", "pass"),
    rek.Timeout(5 * time.Second),
)

fmt.Println(res.StatusCode())
fmt.Println(res.Text())
```

## Responses

The `Response` struct has the following methods:

```go
StatusCode()    int
Content()       []byte // The raw content of the response body
Headers()       map[string]string
Encoding()      []string
Text()          string // string version of the response body
Raw()           *http.Response // The unmodified net/http HTTP response
Cookies()       []*http.Cookie
ContentLength() int64
Status()        string
```

## Options

### Headers

```go
headers := map[string]string{
    "My-Custom-Header": "foo,bar,baz",
}

res, err := rek.Post("https://httpbin.org/post", rek.Headers(headers))
```

### Timeout

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

### File upload

```go
fieldName := "file"
filepath := "docs/README.md"
params := nil

res, err := rek.Post("https:/httpbin.org/post", rek.File(fieldName, filepath, params))
```

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

```go
data := map[string]interface{
    "age": 38,
    "name": "Luc",
}

res, err := rek.Post("https://httpbin.org/post", rek.Data(data))
```

### User agent

```go
res, err := rek.Post("https://httpbin.org/post", rek.UserAgent("ThisGuy"))
```

### Bearer (useful for JSON Web Tokens)

```go
token := "... token ..."

res, err := rek.Post("https://httpbin.org/post", rek.Bearer(token))
```