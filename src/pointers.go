package main

import "fmt"

func pointers() {
	// A pointer is a reference to the memory address of a variable
	// x := 31
	// pointerX := &x
	var x int = 31
	var pointerX *int = &x

	changeValue(pointerX)

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", pointerX, pointerX)

	// You can use a dereference operator (*) to obtain the value stored in a memory location
	fmt.Println(*pointerX)

	// Pointer arithmetic
	// You can't do pointer arithmetics by default in Go, unless by using the "unsafe" library
	// Notice how q and r are 8 bits (or 4 bits, depending on your system) apart
	p := [3]int{1, 2, 3}
	q := &p[0]
	r := &p[1] // - 8 equals q in other programming languages like C and C++
	fmt.Printf("%v %p %p\n", p, q, r)

	// Every pointer that isn't initialized with a value, will have the default value <nil>
	var futurePointer *string
	fmt.Println(futurePointer) // nil

	// Initialize a variable with a pointer to an object
	fmt.Println(&myStruct{})
	var ms *myStruct
	ms = &myStruct{foo: 42}
	fmt.Println(ms)

	// Using the new function
	var ms2 *myStruct
	ms2 = new(myStruct)
	ms2.foo = 55
	fmt.Println(ms2)
	fmt.Println(ms2.foo)
}

func changeValue(val *int) {
	// You can change the value on a memory location by dereferencing a variable
	*val = 42
}

type myStruct struct {
	foo int
}
