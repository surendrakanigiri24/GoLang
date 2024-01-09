package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // if we keep - then it won't shown
	Tags     []string `json:"tags,omitempty"` // if we keep ,omitempty if tags are nil then it won't show
}

func main() {
	fmt.Println("Welcome to Json")

	// Call
	// EncodeJson()
	decodeJson()

}

func EncodeJson() {

	lcoCourses := []course{
		{"React", 299, "online", "1234", []string{"web dev", "js"}},
		{"Mern", 499, "online", "1234", []string{"full stack", "js"}},
		{"Angular", 299, "online", "1234", nil},
	}

	// Package  this data  as JSON data
	// finalJson, err := json.Marshal(lcoCourses)
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func decodeJson() {
	jsonData := []byte(`
		{
			"coursename": "Mern",
			"Price": 499,
			"website": "online",
			"tags": ["full stack","js"]
		}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("Valid Json")
		json.Unmarshal(jsonData, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// Cases where we want to add data to key value
	var myData map[string]interface{}
	json.Unmarshal(jsonData, &myData)
	fmt.Printf("%#v\n", myData)

	for k, v := range myData {
		fmt.Printf("Key is %v and value is %v and Type is: %T\n", k, v, v)
	}

}
