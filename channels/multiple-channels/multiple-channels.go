package main

import "fmt"

func main() {
	channel1 := make(chan int)
	channel2 := make(chan string)

	go func() {
		channel1 <- 10
	}()

	go func() {
		channel2 <- "pong"
	}()

	for range 2 {
		select {
		case channel1Value := <-channel1:
			fmt.Println("Received data from channel 1", channel1Value)
		case channel2Value := <-channel2:
			fmt.Println("Received data from channel 2", channel2Value)
		}
	}
}
