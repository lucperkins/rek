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
	data := Comment{Body: "foo"}

	res, err := rek.Post("https://httpbin.org/post", rek.Data(data))
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
