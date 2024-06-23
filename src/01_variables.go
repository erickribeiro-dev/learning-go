package main

import "fmt"

func variables() {
	// Declare a variable and assign a value to it.
	// Syntax: var <name> <type> = <value>
	var message string = "This is the value of the variable message"
	fmt.Println("message:", message)

	// Go will automatically infer the variable type if you omit it.
	// Syntax: var <name> = <value>
	var year = 2024
	fmt.Println("year:", year)

	// Shorthand syntax
	// Syntax: <name> := <value>
	short_pi := 3.14
	fmt.Println("short_pi:", short_pi)

	// Variables initialized without a value will have a "zero" value (0, "", false, [], nil, etc.)
	// Syntax: var <name> <type>
	var is_admin bool
	fmt.Println("is_admin:", is_admin)

	// Assign or change the value of a variable
	// Syntax: <name> = <value>
	is_admin = true
	fmt.Println("is_admin:", is_admin)

	// Declare multiple variables in the same line
	// Syntax: <name 1>, <name 2>, <name N> := <value 1>, <value 2>, <value N>
	name, age, gender, salary := "Alice", 32, "F", 1234.56
	fmt.Println("name, age, gender and salary:", name, age, gender, salary)

	// Declare multiple variables of the same type in the same line
	// Syntax: var <name 1>, <name 2>, ..., <name N> <type> = <value 1>, <value 2>, ..., <value N>
	var first_name, last_name string = "Alice", "Smith"
	fmt.Println("first_name and last_name:", first_name, last_name)

	// Declare variables in a block
	// Syntax: var ( <variable declarations> )
	var (
		latitude  float64 = 16.019092
		longitude float64 = -25.308471
	)
	fmt.Println("latitude and longitude:", latitude, longitude)
}
