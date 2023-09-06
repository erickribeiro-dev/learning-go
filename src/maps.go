package main

import "fmt"

func maps() {
	// Maps are a collection of key-values pairs of the same type
	// Slices, maps and functions can't be used as keys
	devicesSold := map[string]int{
		"PC":           287159,
		"Tablet":       136938,
		"Mobile Phone": 1395257,
	}
	fmt.Println(devicesSold)

	// Maps can be created with the make function
	statePopulations := make(map[string]int)
	statePopulations["São Paulo"] = 44040000
	statePopulations["Rio de Janeiro"] = 16460000
	statePopulations["Rio Grande do Sul"] = 11290000

	fmt.Println(statePopulations)

	// Delete keys from a map
	delete(statePopulations, "Rio de Janeiro")
	fmt.Println(statePopulations)

	// Check if a key exists in the map
	population, ok := statePopulations["Sãããão Paulo!"]
	fmt.Println(population, ok)

	// Check the length of the map
	fmt.Println(len(statePopulations))

	// Changing a copy of a map also changes the original map
	sp := statePopulations
	sp["Paraná"] = 11080000

	fmt.Println(statePopulations)
}
