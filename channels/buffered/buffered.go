package main

import (
	"fmt"
	"time"
)

// done does the work of goroutine synchronizer
// emailChan <-chan (receive only channel)
// chan<- (send only channel)
func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func() { done <- true }()
	for email := range emailChan { // channel being received here | emailChan runs in infinite loop so needs to be closed
		fmt.Println("Sending email to:", email)
		time.Sleep(time.Second)
	}
}

func main() {
	// We can send around 100 values without blocking
	emailChan := make(chan string, 100) // This is a buffered channel, it wont block till 100 times are being processed and it needs to closed properly else gives deadlock
	done := make(chan bool)             // unbuffered channel

	go emailSender(emailChan, done) // goroutine data communication is established

	for i := range 5 {
		emailChan <- fmt.Sprintf("%d@gmail.com", i) // this is non blocking // data being sent
	}

	fmt.Println("Done Sending...") // This is printed first
	close(emailChan)               // This is important
	<-done

}
