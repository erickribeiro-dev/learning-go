package main

import "fmt"

func arrays() {
	// Arrays (Non-Reference Type) have fixed size
	// Syntax: <name> := [size]type{<value 1>, <value 2>, <value N>}

	// Create an array of size 3
	first_arr := [3]string{"Amanda", "Bob", "Carlos"}
	fmt.Println("first_arr:", first_arr)

	// Create an array of the same size as the amount of initialized values
	second_arr := [...]string{"Europe", "North America", "Asia"}
	fmt.Printf("second_arr: %v of size: %v\n", second_arr, len(second_arr))

	// Create an array of size 5 but with no values.
	// The values will be filled with zero values for that type (0 for type int)
	third_arr := [5]int{}
	fmt.Printf("third_arr: %v of size: %v\n", third_arr, len(third_arr))

	// Add or change values of an array
	third_arr[0] = 2
	third_arr[1] = 3
	third_arr[2] = 5
	third_arr[3] = 7
	third_arr[4] = 11
	fmt.Printf("third_arr: %v of size: %v\n", third_arr, len(third_arr))

	// Unlike slices, arrays are considered values.
	// When the copy of an array is changed, the original array doesn't change
	planets := [...]string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	copy_planets := planets
	copy_planets[2] = "Terra"
	fmt.Println(planets) // still has the original values
	fmt.Println(copy_planets)

	// Arrays of arrays (two-dimensional arrays)
	// var identity_matrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	// var identity_matrix [3][3]int = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}

	// For easier visualization
	var identity_matrix [3][3]int
	identity_matrix[0] = [3]int{1, 0, 0}
	identity_matrix[1] = [3]int{0, 1, 0}
	identity_matrix[2] = [3]int{0, 0, 1}

	fmt.Println(identity_matrix)

	// Pointers can be used to change the value of the original array through another variable
	// another_copy_planets will point to the same memory location as planets
	another_copy_planets := &planets
	another_copy_planets[2] = "Terra"
	fmt.Println(planets) // Earth changed to Terra
}
