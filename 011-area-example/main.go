package main

import "fmt"

type (
	triangle struct {
		height float64
		base   float64
	}
	square struct {
		sidelength float64
	}
	shape interface {
		getArea() float64
	}
)

func main() {
	t := triangle{height: 10, base: 5}
	s := square{sidelength: 4}
	printArea(t)
	printArea(s)
}

func (t triangle) getArea() float64 {
	area := (t.base * t.height) / 2
	return area
}

func (s square) getArea() float64 {
	area := s.sidelength * s.sidelength
	return area
}

func printArea(s shape) {
	fmt.Println("The area of the shape is ", int(s.getArea()))
	fmt.Println()
}
