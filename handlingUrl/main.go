package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://pkg.go.dev/net/http"

func main() {
	fmt.Printf("Welcome handling url in golang")
	fmt.Println(myUrl)

	_result, _ := url.Parse(myUrl)

	fmt.Println(_result.Scheme)
	fmt.Println(_result.Host)
	fmt.Println(_result.Path)
	// fmt.Println(_result.Port)
	fmt.Println(_result.RawQuery)

	qparams := _result.Query()

	for _, val := range qparams {
		fmt.Println(val)
	}

}
