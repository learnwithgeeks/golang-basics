package main

import "fmt"

func main() {

	// boolean types
	n := 1 == 1
	m := 2 == 3
	fmt.Printf("%v,%T", n, n)
	fmt.Printf("\n%v,%T", m, m)
	var a bool
	fmt.Printf("\n%v,%T", a, a)

	// integer types
	c := 1
	d := 2
	fmt.Println(c + d)
	fmt.Println(c * d)
	fmt.Println(c / d)
	fmt.Println(c - d)
	fmt.Println(c % d)

	// floating types
	e := 1.4543
	f := 2.345
	fmt.Println(e + f)
	fmt.Println(e * f)
	fmt.Println(e / f)
	fmt.Println(e - f)

	// string
	g := "hello"
	h := "world"
	fmt.Println(g, string(h[0]))

	// complex types
	var i complex128 = 1 + 2i
	fmt.Printf("%v,%T", real(i), real(i))
	fmt.Printf("\n%v,%T", imag(i), imag(i))

	var j complex128 = complex(1, 2)
	fmt.Printf("\n%v,%T", j, j)
}
