package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Working on gomod")

	// call
	greeter()

	// route call
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	// Server
	log.Fatal(http.ListenAndServe(":4000", r)) // Here log fatal halps to throw error
}

func greeter() {
	fmt.Println("Hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to GoLang </h1>"))
}
