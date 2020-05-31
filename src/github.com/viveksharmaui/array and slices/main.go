package main

import "fmt"

func main() {

	// One Dimentional Arrays (Pass By Value By Default)
	grades := [3]int{1, 2, 3}
	fmt.Println(grades)
	var students [3]string
	fmt.Printf("Students = %v", students)
	students[0] = "Vivek"
	students[1] = "Fahad"
	students[2] = "Darshan"
	fmt.Printf("\nStudents = %v", students)
	fmt.Printf("\nStudent # 1 = %v", students[0])
	fmt.Printf("\nNo. of Students = %v\n", len(students))
	a := [...]string{"hello", "world"}
	fmt.Println(a)

	// Multi-Dimentional Arrays
	var matrix1 = [3][3]int{[3]int{1, 2, 3}, [3]int{1, 2, 3}, [3]int{1, 2, 3}}
	fmt.Println(matrix1)
	var matrix2 [3][3]int
	matrix2[0] = [3]int{1, 2, 3}
	matrix2[1] = [3]int{1, 2, 3}
	matrix2[2] = [3]int{1, 2, 3}
	fmt.Println(matrix2)

	// Slices (Pass By Reference By Default)
	slicea := []int{1, 2, 3}
	sliceb := slicea
	sliceb[1] = 3
	fmt.Println(slicea)
	fmt.Println(sliceb)
	fmt.Println(len(slicea))
	fmt.Println(cap(slicea))
	slicec := make([]int, 3, 100)
	fmt.Println(slicec)
	fmt.Println(len(slicec))
	fmt.Println(cap(slicec))

	// Append to slice
	appendSlice := []int{1, 2, 3}
	appendSlice = append(appendSlice, 4)
	fmt.Println(appendSlice)

	// Shift slice
	shiftSlice := []int{1, 2, 3}
	shiftSlice = shiftSlice[:len(shiftSlice)-1]
	fmt.Println(shiftSlice)
}
