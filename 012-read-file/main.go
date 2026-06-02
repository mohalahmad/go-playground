package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// first approach to read the file, not recommended
	// because it reads the entire file into memory, which can be inefficient for large files
	// but can be useful for small files or when you need to process the entire file at once
	//
	// 	bs, err := os.ReadFile(os.Args[1])
	// 	if err != nil {
	// 		fmt.Println("Error reading file:", err)
	// 		os.Exit(1)
	// 	}
	// 	fmt.Println(string(bs))

	// second approach to read the file, recommended
	// because it reads the file in chunks and writes it to stdout, which is more efficient for large files
	// but can be less convenient if you need to process the entire file at once
	//

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	fmt.Println("File opened successfully:", f.Name())

	// defer f.Close() // this will ensure that the file is closed when the function returns, even if an error occurs
	// below is an alternative way to handle the error when closing the file, but it is not necessary in this case because we are using defer, which will automatically close the file when the function returns
	// or ingnore the error when closing the file, because it is not critical and we are already using defer to ensure that the file is closed
	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("Error closing file: %v", err)
		}
	}()

	_, err = io.Copy(os.Stdout, f)
	if err != nil {
		log.Fatal("Error copying file contents to stdout:", err)
	}
}
