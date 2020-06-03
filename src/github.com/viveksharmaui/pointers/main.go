package main

import "fmt"

func main() {
	// Pointers
	var a int = 45
	var b *int = &a
	a = 25
	fmt.Println(a, *b)

	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1]
	fmt.Println(a, b, c)
}
