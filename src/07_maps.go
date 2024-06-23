package main

import "fmt"

func maps() {
	// Maps (Reference Type) are a collection of key-values pairs of the same type.
	// Slices, maps and functions can't be used as keys
	devices_sold := map[string]int{
		"PC":           287159,
		"Tablet":       136938,
		"Mobile Phone": 1395257,
	}
	fmt.Println(devices_sold["Mobile Phone"])

	// Maps can be created with the make function
	state_population := make(map[string]int)
	state_population["São Paulo"] = 44040000
	state_population["Rio de Janeiro"] = 16460000
	state_population["Rio Grande do Sul"] = 11290000
	fmt.Println(state_population["Rio de Janeiro"])

	// Delete keys from a map
	delete(state_population, "Rio de Janeiro")
	fmt.Println(state_population)

	// Check if a key exists in the map with "ok"
	population, ok := state_population["Rio de Janeiro"]
	fmt.Println(population, ok)

	// Check the length of a map
	fmt.Println("Length:", len(state_population))

	// Changing a copy of a map also changes the original map
	sp := state_population
	sp["Paraná"] = 11080000
	fmt.Println(state_population) // Paraná has been added to state_population

	// The zero value of an empty map is nil created with the make() is nil
	empty_map := make(map[string]string)
	fmt.Printf("empty_map value: %v ; empty_map is nil? %v\n", empty_map, empty_map == nil)

	// Avoid creating an empty map with this syntax, as it can cause "assignment to entry in nil map" errors
	var second_empty_map map[string]string
	fmt.Printf("second_empty_map value: %v ; second_empty_map is nil? %v\n", second_empty_map, second_empty_map == nil)
}
