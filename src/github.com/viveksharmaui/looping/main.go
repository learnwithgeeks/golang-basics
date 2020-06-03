package main

import "fmt"

func main() {

	// For Loop
	for i := 0; i < 5; i++ {
		fmt.Println("Output")
	}

	// Infinite For Loop
	// for {
	// 	fmt.Println("Infinite Output")
	// }

	// Nested Loop
Loop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}

	// For Loop with Collection
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}
}
