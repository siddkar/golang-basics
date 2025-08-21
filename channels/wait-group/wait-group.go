package main

import "fmt"

// goroutine synchronizer
func task(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("Processing...")

}

func main() {
	// unbuffered channels - we can send only one value, until that's received we can't progress
	done := make(chan bool)

	go task(done)

	// To block a single goroutine preferred to use channel where as of multiple preffered to use wait group
	<-done // blocks the code until task returns
}
