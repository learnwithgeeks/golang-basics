package main

import (
	"fmt"
	"log"
)

func main() {
	// Defer
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")

	// Panic
	fmt.Println("start")
	panic("Something bad happened")
	fmt.Println("end")

	// Defer and Panic
	fmt.Println("start")
	defer fmt.Println("defer happened")
	panic("panic happened")
	fmt.Println("end")

	// Defer , Panic and Recover
	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error", err)
		}
	}()
	panic("something bad happened")
	fmt.Println("end")
}
