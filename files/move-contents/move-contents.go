package main

import (
	"bufio"
	"fmt"
	"os"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	sourceFile, err := os.Open("files/move-contents/source.txt")
	handleError(err)

	defer func() {
		fmt.Println("Closing Source File...")
		sourceFile.Close()
	}()

	destinationFile, err := os.Create("files/move-contents/dest.txt")
	handleError(err)

	defer func() {
		fmt.Println("Closing Destination File...")
		destinationFile.Close()
	}()

	// Using streams
	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destinationFile)

	for {
		data, err := reader.ReadByte()

		if err != nil && err.Error() == "EOF" {
			break
		}
		handleError(err)

		e := writer.WriteByte(data)
		handleError(e)

	}
	writer.Flush()
	fmt.Println("File written to destination successfully!")

}
