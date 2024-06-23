package main

import "fmt"

// Go doesn't have inheritance but you can use structs embedding (composition)
type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	// Embeds the Animal struct in the Bird struct
	Animal
	Speed   float32
	Can_fly bool
}

func structs_embedding() {
	// Instantiate a new Bird with the Animal and Bird fields
	bird_one := Bird{}
	bird_one.Name = "Emu"
	bird_one.Origin = "Australia"
	bird_one.Speed = 48
	bird_one.Can_fly = false
	fmt.Println("bird_one:", bird_one)
	fmt.Printf("type of bird_one: %T\n", bird_one)

	// Alternative way to instantiate a new Bird
	bird_two := Bird{
		Animal:  Animal{Name: "Ostrich", Origin: "Africa"},
		Speed:   70,
		Can_fly: false,
	}
	fmt.Println("bird_two:", bird_two)
}
