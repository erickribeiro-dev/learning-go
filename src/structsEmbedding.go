package main

import "fmt"

// Go doesn't have inheritance but you can use structs embedding
type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	// embeds the Animal struct in the Bird struct
	Animal
	Speed  float32
	CanFly bool
}

func structsEmbedding() {
	// instantiate a new Bird with the Animal and Bird fields
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.Speed = 48
	b.CanFly = false
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// another way to instantiate a new Bird
	c := Bird{
		Animal: Animal{Name: "Ostrich", Origin: "Africa"},
		Speed:  70,
		CanFly: false,
	}
	fmt.Println(c)
}
