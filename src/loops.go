package main

import "fmt"

func loops() {
	// Basic for loop. i is scoped to the loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Multiple initializers can be used
	for i, j := 0, 0; i < 4; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	// without an initializer. l is scoped to the main function
	l := 20
	for ; l < 25; l++ {
		fmt.Println(l)
	}

	// without the increment (similar to a do-while loop)
	m := 0
	for m < 5 {
		fmt.Println(m)
		m++
	}

	// Breaking out of an infinite loop
	n := 10
	for {
		fmt.Println("This could run forever.")
		n++
		// Logical test to break out of the loop
		if n == 20 {
			break
		}
	}

	// Continue from the beginning of the loop
	for i := 0; i < 10; i++ { // <--+
		if i%2 == 0 { //            |
			continue // ------------+
		}
		fmt.Println(i)
	}

	// Nested loops and labels
	fmt.Println("Nested loops:")
Loop: // Loop is the name of the label
	for i := 1; i < 4; i++ {
		fmt.Println(i)
		for j := 1; j < 4; j++ {
			fmt.Println(i * j)
			if i*j > 3 {
				break Loop
			}
		}
	}

	// Loop through collections () with for range
	s := []int{10, 20, 30}
	for key, value := range s {
		fmt.Println(key, value)
	}
}
