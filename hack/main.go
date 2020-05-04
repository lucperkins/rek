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

	_, err := rek.Post("https://httpbin.org/post", rek.Data(data), rek.Json(data))
	fmt.Println(err == rek.ErrRequestBodySetMultipleTimes)
}

func exitOnErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
