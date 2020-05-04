# rek

An easy HTTP client for [Go](https://golang.org), inspired by [Requests](https://requests.readthedocs.io/en/master/).

## Example

```go
res, err := rek.Get("https://api.github.com")
// handleErr

js = res.Json
headers := res.Headers
statusCode := res.StatusCode
```