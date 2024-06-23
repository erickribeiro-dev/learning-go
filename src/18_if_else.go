package main

import "fmt"

func if_else() {
	// The most basic if test
	if true {
		fmt.Println("The test is true")
	}

	// Verify by comparison with relational operators: <, <=, >, >=, ==, !=
	// Can also be used in combination with logical operators: &&, ||, !
	age := 20
	wants_to_drink := true
	if age >= 18 && wants_to_drink {
		fmt.Println("You can drink.")
	} else {
		fmt.Println("You cannot drink.")
	}

	// Decision making with maps and the "ok" check
	capitals := map[string]string{
		"Australia":        "Canberra",
		"New Zealand":      "Wellington",
		"Papua New Guinea": "Port Moresby",
	}

	// These variables can only be used in the scope of that if statement
	if pop, ok := capitals["Australia"]; ok {
		fmt.Println(pop)
	} else {
		fmt.Println("Country not found.")
	}

	// Short circuiting:
	// Go doesn't evaluate the comparison tests if it already knows the result is false
	// Notice that the return_true() function will not execute because the statement already
	// knows that x is less than 0
	x := -10
	if x < 0 || return_true() || x > 100 {
		fmt.Println("x is not between 0 and 100")
	}
}

func return_true() bool {
	fmt.Println("the return_true() function has run")
	return true
}
