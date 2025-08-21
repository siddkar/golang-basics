package main

import (
	"fmt"
	"os"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Method 1: Read File Using Buffer
	f, err := os.Open("files/read-file/example.txt")
	handleError(err)
	defer func() {
		fmt.Println("Closing file 1...")
		f.Close()
	}()

	fileInfo, err := f.Stat()
	handleError(err)

	buffer := make([]byte, fileInfo.Size())
	bytes, err := f.Read(buffer)

	for i := range bytes {
		fmt.Println(string(buffer[i]))
	}
	// or (Might not be recommended as it would load the entire file in one go)
	fmt.Println("Method 1: Data:", string(buffer))

	// Method 2: Read file direct (Not recommended for huge files as it load the entire file at once)
	data, err := os.ReadFile("files/read-file/example.txt")
	handleError(err)

	fmt.Println("Method 2: Data:", string(data))

}
