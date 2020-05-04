package main

import (
	"fmt"
	"github.com/lucperkins/rek"
	"log"
	"time"
)

func main() {
	headers := map[string]string{
		"Foo": "bar",
	}

	res, err := rek.Get(
		"https://api.github.com",
		rek.WithHeaders(headers),
		rek.WithTimeout(5*time.Second),
	)

	if err != nil {
		log.Fatal(err)
	}

	loc, _ := res.Raw().Location()

	fmt.Println(loc)
}
