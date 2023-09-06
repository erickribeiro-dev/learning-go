package main

import "fmt"

func arrays() {
	// Arrays have fixed size and, like most other programming languages, start at index 0.
	brands := [3]string{"BMW", "Mercedes"}
	fmt.Println(brands[1]) // Mercedes
	// the len function returns the size of the array
	fmt.Printf("Array length: %v\n", len(brands))

	// Automatically initialize the array size
	fib := [...]int{1, 1, 2, 3, 5, 8}
	fmt.Println(fib)

	// Declare a new array without values
	var students [5]string
	students[0] = "Amanda"
	students[1] = "Bob"
	students[4] = "Elena"

	fmt.Println(students)

	// Arrays of arrays (two-dimensional arrays)
	// var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	// var identityMatrix [3][3]int = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}

	// For easier visualization
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}

	fmt.Println(identityMatrix)

	// Unlike slices, arrays are considered values.
	// When you change the copy of an array, you don't change the original array.
	planets := [...]string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	copyPlanets := planets
	copyPlanets[2] = "Terra"

	fmt.Println(planets) // still has the same values
	fmt.Println(copyPlanets)

	// Pointers can be used to change the value of the original array through another variable
	// anotherCopyPlanets will point to the same memory location as planets
	anotherCopyPlanets := &planets
	anotherCopyPlanets[2] = "Terra"
	fmt.Println(planets) // Earth changed to Terra
}
