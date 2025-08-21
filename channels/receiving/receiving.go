package main

import "fmt"

func sum(result chan int, num1 int, num2 int) {
	sum := num1 + num2
	result <- sum
}

func main() {
	result := make(chan int)
	go sum(result, 4, 5)
	res := <-result // receiving // blocking so no need of wait group or sleep (partially true, doesn't apply for buffer channel)
	fmt.Println("Result: ", res)
}
