package main

import "fmt"

func pointers() {
	// Pointers (Reference Type) are a reference to the memory address of a variable
	var x int = 31
	var pointer_x *int = &x
	change_value(pointer_x)
	fmt.Printf("x value: %v of type: %T\n", x, x)
	fmt.Printf("pointer_x value: %v of type: %T\n", pointer_x, pointer_x)

	// You can use a dereference operator (*) to obtain the value stored in a memory location
	fmt.Println(*pointer_x)

	// Pointer arithmetic
	// You can't do pointer arithmetics by default in Go, unless using the "unsafe" library
	p := [3]int{1, 2, 3}
	q := &p[0] // points to the memory location of p[0]
	r := &p[1] // points to the memory location of p[1]
	// Notice how q and r are 8 bits (or 4 bits, depending on your system) apart
	fmt.Printf("%v %p %p\n", p, q, r)

	// Every pointer that isn't initialized with a value, will have the default value <nil>
	var future_pointer *string
	fmt.Println(future_pointer) // nil

	// Initialize a variable with a pointer to an object
	fmt.Println(&my_struct{})
	var ms *my_struct
	ms = &my_struct{foo: 42}
	fmt.Println("ms:", ms)

	// Initialize a variable using the new function
	var ms2 *my_struct
	ms2 = new(my_struct)
	ms2.foo = 55
	fmt.Println("ms2 and ms2.foo:", ms2, ms2.foo)
}

func change_value(val *int) {
	// You can change the value on a memory location by dereferencing a variable,
	// that is, obtaining the value of the corresponding memory address
	*val = 42
}

type my_struct struct {
	foo int
}
