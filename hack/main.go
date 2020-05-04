package main

import (
	"fmt"
	"github.com/lucperkins/rek"
	"io/ioutil"
	"log"
)

func main() {
	headers := map[string]string{
		"Foo": "bar",
	}

	res, err := rek.Get("https://api.github.com", rek.WithHeaders(headers))
	if err != nil {
		log.Fatal(err)
	}

	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))
}
