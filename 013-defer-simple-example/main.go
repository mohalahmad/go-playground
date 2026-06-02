package main

import "fmt"

func main() {
	fmt.Println("1")
	defer fmt.Println("2") // this will be executed after the main function returns, but before the program exits
	fmt.Println("3")
}
