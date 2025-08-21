package main

import (
	"fmt"
	"math/rand"
	"time"
)

func processNum(numChan chan int) {
	for num := range numChan { // while using range <- is not required
		fmt.Println("Processing numbers from channel: ", num)
		time.Sleep(time.Second * 1)
	}
}

// sending
func main() {
	numChan := make(chan int)
	go processNum(numChan) // sending channel from main goroutine to processNum
	numChan <- 5

	for {
		numChan <- rand.Intn(100)
	}

}
