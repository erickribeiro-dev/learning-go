package main

import "fmt"

func ifElseDecisionMaking() {
	// The most basic if test
	if true {
		fmt.Println("The test is true")
	}

	// Verify by comparison with relational operators: <, <=, >, >=, ==, !=
	// Can also be used in combination with logical operators: &&, ||, !
	age := 20
	wantsToDrink := true
	if age >= 18 && wantsToDrink {
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
	// Notice that the returnTrue() function will not execute
	x := -10
	if x < 0 || returnTrue() || x > 100 {
		fmt.Println("x is not between 0 and 100")
	}
}

func returnTrue() bool {
	fmt.Println("the returnTrue() function has run")
	return true
}
