package main

import (
	"fmt"
)

func constants() {
	// Constants can't have their values changed
	// and the only types allowed for constants are numbers, strings and booleans
	const maxAllowedConnections = 128
	fmt.Println(maxAllowedConnections)

	// The following line will result in error because the value should be determined at compile time
	// and math.Sin is calculated at run time
	// const result float64 = math.Sin(1.57) <-- Will result in an error

	// you can do operations with variables and constants
	const x int = 12
	var y int = 18
	fmt.Println(x + y)

	// You can't do operations if you determine different types for the variables,
	// but you can do so if you don't set the type for the constant
	// const j int = 8
	// var k int16 = 12
	// fmt.Println(j + k) <-- Will result in an error
	const j = 8
	var k int16 = 12
	fmt.Printf("%v of type %T\n", j+k, j+k) // this works, and the result is of Type int16
	// the above line is the same as the following line because the compile replaces the variable with the literal value:
	fmt.Println(8 + k)
}
