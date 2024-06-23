package main

import "fmt"

// Structs (Non-Reference Type) are a collection of fields that can have different types
// Note that the struct name starts with a capital letter, so it can be used by other packages

// Structs
type Employee struct {
	name   string
	age    int
	salary float64
	skills []string
}

func structs() {
	employee_one := Employee{
		name:   "Elizabeth",
		age:    34,
		salary: 1234.56,
		skills: []string{"data analysis", "financial analysis"},
	}

	fmt.Println("employee_one:", employee_one)

	// Access a property with . and the field name
	fmt.Println("name:", employee_one.name)
	fmt.Println("age:", employee_one.age)
	fmt.Println("salary:", employee_one.salary)
	fmt.Println("skill:", employee_one.skills[0])
	fmt.Println("skill:", employee_one.skills[1])

	// If you know the order of the fields, you don't need to specify the names and can simply
	// create a new instance with the values
	employee_two := Employee{"Albert", 26, 2345.67, []string{"Cloud Security", "Network Security"}}
	fmt.Println(employee_two)

	// Fields can be empty when instantiating a new struct
	employee_three := Employee{}
	fmt.Println(employee_three)

	// Anonymous structs
	person_one := struct{ name string }{name: "Nicole Nash"}
	fmt.Println(person_one)

	// Copying a struct will create a copy by value, not by reference
	employee_four := employee_two
	employee_four.name = "Dorothy"
	fmt.Println("employee_two:", employee_two.name)  // still has the original values
	fmt.Println("employee_four", employee_four.name) // copy of employee_two with a different "name" field

	// Create a pointer reference
	employee_five := &employee_four
	employee_five.name = "Raquel"
	fmt.Println("employee_four.name:", employee_four.name) // The value has changed because employee_five has changed and is a pointing to employee_four
	fmt.Println("employee_five.name:", employee_five.name)
}
