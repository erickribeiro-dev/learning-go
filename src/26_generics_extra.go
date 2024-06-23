package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Map function receives a slice of integers, a function and return a resulting slice of integers
func MapValues(values []int, mapFunc func(int) int) []int {
	var result []int
	// For each value in values, call the mapFunc() on it and add the result to the resulting slice
	for _, val := range values {
		newValue := mapFunc(val)
		result = append(result, newValue)
	}
	return result
}

// The problem is that the function only accept integers.
// Let's modify it so it accepts all ints and floats
func MapValuesV2[T constraints.Ordered](values []T, mapFunc func(T) T) []T {
	var result []T
	// For each value in values, call the mapFunc() on it and add the result to the resulting slice
	for _, val := range values {
		newValue := mapFunc(val)
		result = append(result, newValue)
	}
	return result
}

// Generics can be used on structs
type CustomData interface {
	constraints.Ordered | []byte
}

// Data T is any type defined by the interface CustomData
type UserWithData[T CustomData] struct {
	ID   int
	Name string
	Data T
}

// Generics can be used on maps. Let's use any comparable types as the map key,
// which excludes the types: slices, maps and functions.
// Values (V) can be int or string
type customMap[T comparable, V int | string] map[T]V

func generics_extra() {
	// Testing the map functions
	myIntSlice := []int{1, 2, 3}
	myFloatSlice := []float64{1.0, 2.0, 3.0}
	resultInt := MapValues(myIntSlice, func(i int) int {
		return i * 2
	})
	resultFloat := MapValuesV2(myFloatSlice, func(i float64) float64 {
		return i * 2
	})
	fmt.Println(resultInt)
	fmt.Println(resultFloat)

	// Testing the custom data types
	usr := UserWithData[string]{
		ID:   0,
		Name: "Peter",
		Data: "Any defined data",
	}
	fmt.Println(usr)

	// Testing with a custom map
	myMap := make(customMap[int, string])
	myMap[3] = "Hello"
	fmt.Println(myMap)
}
