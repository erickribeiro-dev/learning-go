package main

import "fmt"

// iota is an enumerated constant and the initial value is always 0
const z = iota

// It increments by 1 everytime it's used in the same block of constants declarations
const (
	b = iota
	c // = iota
	d // = iota
	e // = iota
)

// When declaring a new constants block the value of iota resets to 0
const (
	// It's recommended to skip the first value to avoid confusing errors
	// such as initializing a specialistType without a value
	// (the value will automatically be set to 0, which would be catSpecialist).
	// Alternatively, the initial value name for iota can be errorSpecialist
	_               = iota // 0
	catSpecialist          // 1
	dogSpecialist          // 2
	snakeSpecialist        // 3
)

// Using bit shifting operators is very useful when working with iota
const (
	_  = iota
	KB = 1 << (10 * iota) // 1 * 2^10
	MB                    // 1 * 2^100
	GB                    // 1 * 2^1000
	TB                    // 1 * 2^10000
	PB                    // 1 * 2^100000
)

// You can use bit shifting to set boolean flags in a single byte
const (
	isAdmin            = 1 << iota // 00000001
	isHeadquarters                 // 00000010
	canSeeFinancials               // 00000100
	canSeeAfrica                   // 00001000
	canSeeAsia                     // 00010000
	canSeeEurope                   // 00100000
	canSeeNorthAmerica             // 01000000
	canSeeSouthAmerica             // 10000000
)

func enumConstants() {
	fmt.Printf("%v %T\n", z, z) // iota is 0 and of type integer
	fmt.Println(b, c, d, e)     // 0 1 2 3

	var specialistType int = dogSpecialist
	fmt.Printf("Is dog specialist? %v\n", specialistType == dogSpecialist) // returns true

	var fileSize float64 = 4000000000
	fmt.Printf("%.2fGB\n", fileSize/GB) // 3.73GB

	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)                                         // 100101
	fmt.Printf("Is admin? %v\n", isAdmin&roles == isAdmin)            // true
	fmt.Printf("Is HQ? %v\n", isHeadquarters&roles == isHeadquarters) // false
}
