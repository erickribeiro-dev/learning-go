package main

import "fmt"

// A struct is a collection of fields that can have different types
// Note that the struct name starts with a capital D, so it can be used by other packages
type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func structs() {
	// You can construct a new struct with the following syntax
	aDoctor := Doctor{
		number:    3,
		actorName: "John Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah",
			"Jane Smith",
		},
	}
	fmt.Println(aDoctor)
	// You can access a property with . and the field name
	fmt.Println(aDoctor.companions[1])

	// If you know the order of the fields, you don't need to specify the names and cam simply
	// create a new instance with the values
	bDoctor := Doctor{4, "abc", []string{"def", "ghi"}}
	fmt.Println(bDoctor)

	// You don't need to provide the fields when instantiating a new struct
	cDoctor := Doctor{}
	fmt.Println(cDoctor)

	// Anonymous structs
	dDoctor := struct{ name string }{name: "John Pertwee"}
	fmt.Println(dDoctor)

	// Copying a struct will create a copy by value, not by reference
	eDoctor := dDoctor
	eDoctor.name = "Tom Baker"
	fmt.Println(dDoctor)
	fmt.Println(eDoctor)

	// You can create a pointer reference
	fDoctor := &dDoctor
	fDoctor.name = "Tom Baker"
	fmt.Println(fDoctor)
	fmt.Println(dDoctor.name)
}
