package main

import (
	"fmt"
	"log"
	"mongoAPI/router"
	"net/http"
)

func main() {
	fmt.Println("Working on API's with MONGO DB")
	r := router.Router()
	fmt.Println("Server is getting started")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening on port 4000 ..")
}
