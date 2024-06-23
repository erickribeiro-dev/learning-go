package main

import "fmt"

func slices() {
	// Slices (Reference Type) sizes can change during the lifetime of the program
	bodies_of_water := []string{"Lake", "Ocean", "River", "Sea"}
	fmt.Println("bodies_of_water:", bodies_of_water)

	// The len method returns the size of the slice
	// and the cap method returns the capacity of the slice
	fmt.Println("length and capacity of bodies_of_water:", len(bodies_of_water), cap(bodies_of_water))

	// Unlike arrays, slices are reference types
	// so when we change the copy of a slice, the original slice changes as well
	copy_bodies_of_water := bodies_of_water
	copy_bodies_of_water[0] = "Creek"
	fmt.Println("bodies_of_water:", bodies_of_water)
	fmt.Println("copy_bodies_of_water:", copy_bodies_of_water)

	// There are several ways to create a slice.
	// Create an slice from an array:
	array_prime_numbers := [...]int{2, 3, 5, 7, 11, 13, 17}
	slice_one := array_prime_numbers[:]
	fmt.Printf("Type of prime_numbers: %T\n", array_prime_numbers)
	fmt.Printf("Type of slice_one: %T\n", slice_one)

	// Create a slice from a range of indexes
	slice_two := array_prime_numbers[3:]
	slice_three := array_prime_numbers[:6]
	slice_four := array_prime_numbers[3:6]
	fmt.Println("slice_two:", slice_two)
	fmt.Println("slice_three:", slice_three)
	fmt.Println("slice_four:", slice_four)

	// Create a slice with the make function
	biomes := make([]string, 7, 10)
	biomes[0] = "Savannah"
	biomes[1] = "Desert"
	biomes[2] = "Grassland"
	biomes[3] = "Temperate forest"
	biomes[4] = "Tundra"
	fmt.Println("biomes:", biomes)
	fmt.Println("length and capacity of biomes:", len(biomes), cap(biomes))

	// Add elements to a slice
	biomes = append(biomes, "Tropical rainforest", "Boreal forest")
	fmt.Println("biomes:", biomes)

	// Concatenate slices using spread operator
	i := []string{"Taiga"}
	biomes = append(biomes, i...)
	fmt.Println("biomes:", biomes)

	// Stack operators
	// Shift operation
	j := []int{1, 2, 3, 4, 5}
	k := j[1:]
	l := j[:len(j)-1]
	fmt.Println("j:", j)
	fmt.Println("k and l:", k, l)

	// Remove an element from the middle
	m := append(j[:2], j[3:]...)
	fmt.Println("m:", m) // [1 2 4 5]

	// Note that removing elements can modify the original slice
	fmt.Println("j:", j) // [1 2 4 5 5]

	// Copy one slice into another
	n := make([]int, len(m)) // creating a new slice n
	copy(n, m)               // Copying m to n
	fmt.Println("m and n:", m, n)

	// Iterate over a slcie
	array_prime_numbers = [...]int{2, 3, 5, 7, 11, 13, 17}
	for index, val := range array_prime_numbers {
		fmt.Println(index, val)
	}
}
