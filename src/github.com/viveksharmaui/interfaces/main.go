package main

import "fmt"

// Interface => Define Behaviour
type Shape interface {
	Area() float64
}

// Struct
type Rect struct {
	width  float64
	height float64
}

// Area Method
func (r Rect) Area() float64 {
	return r.width * r.height
}


// Main Function
func main() {
	var s Shape
	s = Rect{5.0, 4.0}
	fmt.Println("Area of Rectangle Is : ", s.Area())
}
