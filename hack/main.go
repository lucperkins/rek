package main

import (
	"fmt"
	"github.com/lucperkins/rek"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

var luc = Person{
	Name: "Luc",
	Age: 38,
}

func main() {
	res, err := rek.Get("https://httpbin.org/anything",
		rek.UserAgent("This-Guy"),
		rek.Headers(map[string]string{"Foo": "bar"}),
		rek.Data(map[string]interface{}{"foo": 1}),
		rek.Struct(luc),
	)
	exitOnErr(err)

	fmt.Println(res.Text())
}

func exitOnErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
