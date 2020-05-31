package main

import "fmt"

type Doctor struct {
	firstName  string
	middleName string
	lastName   string
}

func main() {

	// Map (Key Value) -> Pass By Reference
	a := map[string]string{
		"firstName":  "Vivek",
		"middleName": "Anand",
		"lastName":   "Sharma",
	}
	delete(a, "firstName")
	_, ok := a["firstName"]
	fmt.Println(a, ok)
	fmt.Println(len(a))

	//Structs
	aDoctor := Doctor{
		firstName:  "Vivek",
		middleName: "Anand",
		lastName:   "Sharma",
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.firstName)

	bDoctor := struct{ name string }{name: "vivek"}
	fmt.Println(bDoctor)
}
