package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Printf("Make web request")

	// Call
	// preformGetRequest()

	// Post
	// performPostJsonRequest()

	// Post form data
	performPostFormDataRequest()
}

func preformGetRequest() {
	const myUrl = "http://localhost:3000"

	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length is:", response.ContentLength)

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))

}

func performPostJsonRequest() {
	const myUrl = "http://localhost:3000/iamGoReq"

	// Json Payload
	requestBody := strings.NewReader(`
		{
			"name":"surendra",
			"mail":"kanigiri"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}

func performPostFormDataRequest() {
	const myUrl = "http://localhost:3000/iamGoReq"

	// form data
	data := url.Values{}
	data.Add("firstName", "Surendra")
	data.Add("lastName", "Kanigiri")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}
