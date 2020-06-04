package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		sayMessage("Hello Go!", i)
	}
	hello := "Hello"
	name := "Vivek"
	sayGreeting(&hello, &name)
	fmt.Println(name)
	s := sum(1, 2, 3, 4, 5)
	fmt.Println("The Sum is :", *s)

	// Anonymous Functions
	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}

	var f func() = func() {
		fmt.Println("Hello Go!")
	}
	f()
}

func sayMessage(msg string, id int) {
	fmt.Println("Message is : ", msg)
	fmt.Println("Id is : ", id)
}

func sayGreeting(hello, name *string) {
	fmt.Println(*hello, *name)
	*name = "Fahad"
}

func sum(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}
