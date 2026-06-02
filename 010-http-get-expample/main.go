package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(resp.Status)
	// first approach to read the body, not recommended
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// second approach to read the body, recommended
	// io.Copy(os.Stdout, resp.Body)

	// third approach to read the body, using a custom writer
	lw := logWriter{}
	_, err = io.Copy(lw, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Printf("Just wrote this many bytes: %d\n", len(bs))
	return len(bs), nil
}
