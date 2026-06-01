package main

import "fmt"

func main() {
	// there is multiple ways to create a map
	//
	// -- option 1
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	fmt.Println(colors)

	// -- option 2
	var colors2 map[string]string
	// colors2["white"] = "#ffffff" // when map defined with var, it is nil and will cause panic when we try to assign value to it
	fmt.Println(colors2)

	// -- options 3
	colors3 := make(map[string]string)
	colors3["black"] = "#000000" // when map defined with make, it is initialized and we can assign value to it
	fmt.Println(colors3)

	colors4 := make(map[int]string)
	colors4[1] = "red"
	colors4[2] = "black"
	fmt.Println(colors4)

	// to delete a key from map, we can use delete function
	delete(colors4, 1) // delete key "red" from colors map
	fmt.Println(colors4)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Hex code for %s is %s\n", color, hex)
		fmt.Println("Hex code for", color, "is", hex)
	}
}
