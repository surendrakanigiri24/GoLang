package main

import "fmt"

func main() {
	fmt.Println("Channels are used to communicate with each wait group")

	myCh := make(chan int)

	myCh <- 5
	fmt.Println(<-myCh)
}
