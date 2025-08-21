package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("files/example.txt")
	if err != nil {
		// log the error
		panic(err)
	}

	defer f.Close()

	// get file info
	fileInfo, err := f.Stat()
	if err != nil {
		// log the error
		panic(err)
	}

	// read file
	buf := make([]byte, fileInfo.Size())

	d, err := f.Read(buf)
	if err != nil {
		// log the error
		panic(err)
	}
	for i := range buf {
		fmt.Println("data", d, string(buf[i]))
	}

	fmt.Println("data", d, string(buf))
	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	// log the error
	// 	panic(err)
	// }

	// fmt.Println("file name:", fileInfo.Name())
	// fmt.Println("file or folder:", fileInfo.IsDir())
	// fmt.Println("file size:", fileInfo.Size())
	// fmt.Println("file permission:", fileInfo.Mode())
	// fmt.Println("file modified at:", fileInfo.ModTime())

}
