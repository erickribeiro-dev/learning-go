package main

import "fmt"

func variables() {
	// Variables in Go are strong typed,
	// so if you don't initialize a value, you still have to define its type
	var z int
	z = 3

	// Syntax: <name> := <value>
	a := 4

	// Syntax: var <name> = <value>
	var b = 5

	// Syntax: var <name> <type> := <value>
	var c float32 = 3.141592

	// Declaring multiple variables
	// Syntax: <name1, name2, ..., nameN> := <value1, value2, ..., valueN>
	d, e, f := 6, 7, 8

	// Declaring variables in a block
	var (
		g int    = 9
		h string = "Variable h declared in a block"
	)

	// Go automatically initializes empty variables with falsy values
	var i string // ""
	var j int    // 0
	var k bool   // false

	fmt.Println(z)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
}
