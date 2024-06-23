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
	// such as initializing a specialist_type without a value
	// (the value will automatically be set to 0, which would be cat_specialist).
	// Alternatively, the initial value name for iota can be error_specialist
	_                = iota // 0
	cat_specialist          // 1
	dog_specialist          // 2
	snake_specialist        // 3
)

// Bit shifting operators are very useful when working with iota
const (
	_  = iota
	KB = 1 << (10 * iota) // 1 * 2^10
	MB                    // 1 * 2^100
	GB                    // 1 * 2^1000
	TB                    // 1 * 2^10000
	PB                    // 1 * 2^100000
)

// Using bit shifting to set boolean flags in a single byte
const (
	is_admin              = 1 << iota // 00000001
	is_headquarters                   // 00000010
	can_see_financials                // 00000100
	can_see_africa                    // 00001000
	can_see_asia                      // 00010000
	can_see_europe                    // 00100000
	can_see_north_america             // 01000000
	can_see_south_america             // 10000000
)

func enumerated_constants() {
	fmt.Printf("%v %T\n", z, z) // iota is 0 and of type integer
	fmt.Println(b, c, d, e)     // 0 1 2 3

	var specialist_type int = dog_specialist
	fmt.Printf("Is dog specialist? %v\n", specialist_type == dog_specialist) // true

	var fileSize float64 = 4000000000
	fmt.Printf("%.2fGB\n", fileSize/GB) // 3.73GB

	var roles byte = is_admin | can_see_financials | can_see_europe
	fmt.Printf("%b\n", roles)                                           // 100101
	fmt.Printf("Is admin? %v\n", is_admin&roles == is_admin)            // true
	fmt.Printf("Is HQ? %v\n", is_headquarters&roles == is_headquarters) // false
}
