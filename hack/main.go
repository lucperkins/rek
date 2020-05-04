package main

import (
	"fmt"
	"github.com/lucperkins/rek"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var luc = Person{
	Name: "Luc",
	Age:  38,
}

func main() {
	form := map[string]string{
		"foo": "bar",
	}

	fmt.Println(form)

	res, err := rek.Get("https://httpbin.org/anything",
		//rek.UserAgent("This-Guy"),
		//rek.Headers(map[string]string{"Foo": "bar"}),
		//rek.File("file", "go.mod", form),
		//rek.FormData(form),
		//rek.BasicAuth("haxor", "opensesame"),
		//rek.Headers(map[string]string{"Location": "https://google.com"}),
		//rek.Struct(luc),
		rek.JWT("here-is-my-token"),
	)
	exitOnErr(err)

	fmt.Println(res.Text())
	fmt.Println(res.Status())

	loc, _ := res.Raw().Location()
	fmt.Println(loc)
}

func exitOnErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
