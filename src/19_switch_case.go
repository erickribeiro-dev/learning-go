package main

import "fmt"

func switch_case() {
	// Compare a value with multiple other values
	year := 2024
	switch year {
	case 2022:
		fmt.Println("The year is 2022.")
	case 2023:
		fmt.Println("The year is 2023.")
	case 2024:
		fmt.Println("The year is 2024.")
	// Comparing multiple values at once
	case 2019, 2020, 2021:
		fmt.Println("That was some years ago.")
	case 2025, 2026, 2027:
		fmt.Println("That's a few years from now.")
	default:
		fmt.Println("The year is far from the current year.")
	}

	// Using an initializer.
	switch i := 2 + 3; i {
	case 4:
		fmt.Println("The result is: Four")
	case 5:
		fmt.Println("The result is: Five")
	case 6:
		fmt.Println("The result is: Six")
	}

	// Tagless switches
	// Notice that only the first case will execute if you ommit the fallthrough keyword.
	// When using fallthrough, the next "case" will run regardless if the test passes or not
	j := 10
	switch {
	case j <= 10:
		fmt.Println("j is less or equal to 10.")
		fallthrough
	case j <= 20:
		fmt.Println("j is less or equal to 20.")
	default:
		fmt.Println("j is greater than 20.")

	}

	// Type switches and the use of the break keyword
	var k interface{} = 1
	switch k.(type) {
	case int:
		fmt.Println("k is int")
		fmt.Println("This line will be printed too")
	case float64:
		fmt.Println("k is float64")
	case string:
		fmt.Println("k is string")
	default:
		fmt.Println("k is another type")
	}
}
