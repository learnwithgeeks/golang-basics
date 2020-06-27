package main

// Synchronous package
import (
	"fmt"
	"sync"
)

// Wait Group
var wg = sync.WaitGroup{}

// Main Function
func main() {
	wg.Add(1)
	go sayHello()
	fmt.Println("Hello World 2")
	wg.Wait()
}

// SayHello Function
func sayHello() {
	fmt.Println("Hello World 1")
	wg.Done()
}
