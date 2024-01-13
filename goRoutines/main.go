package main

import (
	"fmt"
	"net/http"
	"sync"
)

var someRand = []string{"test"}
var wg sync.WaitGroup // Pointer
var mut sync.Mutex    // Pointer lock or unloc memory space

func main() {
	// go greeter("Hello")
	// greeter("World")

	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(someRand)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endPoint string) {
	defer wg.Done()

	res, err := http.Get(endPoint)

	if err != nil {
		fmt.Println("OOPS ENDPOINT")
	} else {
		mut.Lock()
		someRand = append(someRand, endPoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endPoint)

	}

}
