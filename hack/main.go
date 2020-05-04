package main

import (
	"fmt"
	"github.com/lucperkins/rek"
	"os"
)

type Comment struct {
	Body string `json:"body"`
}

func main() {
	comment := Comment{Body: "this is a cool Gist!"}

	res, err := rek.Post("https://httpbin.org/post",
		rek.Json(comment),
		rek.Headers(map[string]string{"Custom-Header": "foo,bar,baz"}))
	exitOnErr(err)

	fmt.Println(res.StatusCode())
	fmt.Println(res.Text())
}

func exitOnErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
