# rek

An easy HTTP client for [Go](https://golang.org), inspired by [Requests](https://requests.readthedocs.io/en/master/). Here's an example:

```go
type Comment struct {
    Body string `json:"body"`
}

comment := Comment{Body: ""}

headers := map[string]string{"My-Custom-Header", "foo,bar,baz"}

res, err := rek.Post("https://api.github.com/",
    rek.Json(comment),
    rek.Headers(headers),
    rek.BasicAuth("user", "pass"),
    rek.Timeout(5 * time.Second),
)

fmt.Println(res.StatusCode())
fmt.Println(res.Text())
```
