package main

import "fmt"

func main() {
	// if statement
	if true {
		fmt.Println("The test is true")
	}

	statePopulation := map[string]int{
		"Pakistan": 1000,
	}

	if pop, ok := statePopulation["Pakistan"]; ok {
		fmt.Println(pop)
	}

	number := 50
	guess := 50
	if guess == number || returnTrue() {
		fmt.Println("To Equal")
	}
	if guess < number {
		fmt.Println("Less Than")
	}
	if guess > number {
		fmt.Println("Greater Than")
	}

	// if else if statement
	guessB := 40
	if guessB < 30 {
		fmt.Println("Too Low")
	} else if guessB > 30 {
		fmt.Println("Too High")
	} else {
		fmt.Println("To Equal")
	}

	// Switch Statement
	switch 2 {
	case 1:
		fmt.Println("Case 1")
	case 2:
		fmt.Println("Case 2")
	default:
		fmt.Println("Case Default")
	}

	// Multiple cases in switch statement

	switch 7 {
	case 1, 3, 5:
		fmt.Println("one three five")
	case 7, 9, 10:
		fmt.Println("seven nine ten")
		fallthrough
	default:
		fmt.Println("default case")
	}

	switch i := 2 + 3; i {
	case 2:
		fmt.Println("2nd Case")
	case 5:
		fmt.Println("5th Case")
	default:
		fmt.Println("Default Case")
	}
}

func returnTrue() bool {
	fmt.Println("return true")
	return true
}
