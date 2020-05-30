package main

import "fmt"

// constants shadow
const a int = 10

// iota
const (
	f = iota
	g
	h
)

func main() {
	fmt.Printf("%v\n", f)
	fmt.Printf("%v\n", g)
	fmt.Printf("%v\n", h)
	const a int = 5
	var e int = 10
	const b float32 = 2.4
	const c string = "string constant"
	const d bool = true
	fmt.Printf("%v,%T", a+e, a+e)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	const f = 10
	fmt.Println(f)
}
