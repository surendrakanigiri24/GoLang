package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "http://www.google.com"

func main() {
	fmt.Printf("Web request\n")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Type of web response is %T\n", response)
	defer response.Body.Close() // Caller responsibility to close connection

	dataBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(dataBytes)
	fmt.Printf(content)
}
