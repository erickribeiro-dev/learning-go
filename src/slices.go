package main

import "fmt"

func slices() {
	// Slices size can change over the lifetime of the program
	a := []int{1, 3, 5}
	fmt.Println(a)

	// the len method returns the size of the slice
	// and the cap method returns the capacity of the slice
	fmt.Printf("slice length and capacity: %v, %v\n", len(a), cap(a))

	// Unlike arrays, slices are reference types
	// so when we change the copy of a slice, the original slice changes as well
	b := a
	b[0] = 2

	fmt.Println(a)
	fmt.Println(b)

	// There are several ways to create a slice
	// You can create a slice from an array
	c := [...]int{10, 15, 20, 25, 30, 35, 40, 45}
	d := c[:]
	fmt.Printf("Type of c: %T\n", c)
	fmt.Printf("Type of d: %T\n", d)

	// You can create slices from a range of indexes
	e := c[3:]
	f := c[:6]
	g := c[3:6]

	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	// Using the make function to create a slice
	h := make([]string, 4, 100)
	fmt.Printf("values: %v - length: %v - capacity: %v", h, len(h), cap(h))

	// Add elements to a slice
	h = append(h, "River", "Ocean", "Lake")
	fmt.Println(h)

	// Concatenate slices using spread operator
	i := []string{"Sea"}
	h = append(h, i...)
	fmt.Println(h)

	// Stack operators
	// Shift operation
	j := []int{1, 2, 3, 4, 5}
	k := j[1:]
	l := j[:len(j)-1]
	fmt.Println(k, l)

	// Remove an element from the middle
	m := append(j[:2], j[3:]...)
	fmt.Println(m) // [1 2 4 5]

	// Note that removing elements can modify the original slice
	fmt.Println(j) // [1 2 4 5 5]

}
