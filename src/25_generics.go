package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Generics allow a function to accept different types for the same parameters.
// We have 2 similar functions - addInt and addFloat - that can be combined into one using Generics
func addInt(a int, b int) int {
	return a + b
}
func addFloat(a int, b int) int {
	return a + b
}

// T can be int or float64
func add[T int | float64](a T, b T) T {
	return a + b
}

// To simplify the verbosity of the list of types passed, an interface can be created
// with all the types needed and used as the Generics
type Num interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func addV2[T Num](a T, b T) T {
	return a + b
}

// We can go further and utilize the constraints package, which has many preset interfaces
func addV3[T constraints.Ordered](a T, b T) T {
	return a + b
}

// using a tilde (~) before the type allow us to pass any alias of the same type
type MyInt int

func addV4[T ~int | float64](a T, b T) T {
	return a + b
}

func generics() {
	// Testing the functions and printing the results
	resultInt := addInt(1, 1)
	resultFloat := addFloat(2.0, 2.0)
	resultGenerics := add(3, 3)
	resultIGenerics := addV2(4, 4)
	resultCGenerics := addV3(5, 5)

	a := MyInt(6)
	b := MyInt(6)
	resultAlias := addV4(a, b)
	fmt.Println(resultInt, resultFloat, resultGenerics, resultIGenerics, resultCGenerics, resultAlias)
}
