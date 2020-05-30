package main

import (
	"fmt"
	"strconv"
)

var (
	n bool
)

func main() {
	var i int = 42
	fmt.Printf("%v,%T\n", i, i)
	var j string
	j = strconv.Itoa(i)
	fmt.Printf("%v,%T", j, j)
	fmt.Print("\n")
	n = true
	fmt.Printf("%v,%T", n, n)
}
