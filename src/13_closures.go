package main

import "fmt"

// Example #1
// Functions are a data type. They can be passed as parameters, stored in variables, be returned, etc.
// In this example, the function intSequence returns "a function that returns an int"
func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// Example #2
// Using recursion to calculate the factorial of a number
func factorial(i int) int {
	if i == 0 {
		return 1
	}
	return i * factorial(i-1)
}

func closures() {
	// increment is a function, and it captures its own value everytime it's called, updating the value
	increment := counter()
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
	fmt.Println(increment()) // 3

	// Factorial
	fmt.Printf("Factorial of %v: %v\n", 8, factorial(8))

	// fibonacci
	var fibonacci func(n int) int
	fibonacci = func(n int) int {
		if n < 2 {
			return n
		}
		return fibonacci(n-1) + fibonacci(n-2)
	}

	fmt.Printf("Fibonacci of %v: %v\n", 8, fibonacci(8))
}
