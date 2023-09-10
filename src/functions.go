package main

import "fmt"

func functions() {
	// Calling a function with a parameter (argument),
	// similar to most proframming languages
	for i := 0; i < 5; i++ {
		sayMessage("Hello Go!", i)
	}

	// If all arguments have the same type,
	// you only need to specify the type once in the function declaration
	sayGreeting("Hello", "Peter")

	// You can manipulate the original parameter if you pass it as a pointer as argument
	// Remember that maps and slices are always passed in as pointers
	name := "John"
	sayHelloAndAlterName(&name)
	fmt.Println(name)

	// Variadic parameters (maximum 1 variadic parameter and must be the last listed in the arguments)
	// the type of the variadic parameter is a slice
	sumFunction("The sum is:", 1, 3, 5, 7, 9)

	// Return values
	fmt.Println(getLength("This is a 36 characters long message"))

	// Returninit a pointer
	fmt.Println(*isEven(88))

	// Named return
	fmt.Println(isOdd(77))

	// multiple return values are stored in multiple variables
	divideResult, err := divideFunction(4.0, 1.5)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(divideResult)

	// Anonymous functions
	func() {
		fmt.Println("This text was printed inside an anonymous function")
	}()

	// Functions as types
	// var f func() = func() {
	f := func() {
		fmt.Println("You called a function declared as type")
	}
	f()
	fmt.Printf("%T\n", f) // f is of type func()

	// Methods
	cat1 := Cat{
		name:     "Nova",
		greeting: "Meow!",
		age:      3,
	}
	cat2 := Cat{
		name:     "Sagan",
		greeting: "Meoow!",
		age:      2,
	}
	cat1.speak()
	cat2.speak()
}

// Basic function with arguments
func sayMessage(msg string, index int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is:", index)
}

// All arguments of the same type
func sayGreeting(msg, name string) {
	fmt.Println(msg, name)
}

// Pointers as arguments
func sayHelloAndAlterName(name *string) {
	fmt.Println(*name)
	*name = "Scott"
	fmt.Println("Your name is now", *name)
}

// Variadic parameters
// the type of the variadic parameter is a slice
func sumFunction(msg string, args ...int) {
	fmt.Printf("%v %T\n", args, args)
	result := 0
	for _, val := range args {
		result += val
	}
	fmt.Println(msg, result)
}

// Return values
func getLength(msg string) int {
	return len(msg)
}

// You can return a pointer to a variable declared inside a function.
// Unlike most programming languages, Go will recognize that you're returning a pointer and
// keep the variable in the heap (shared) memory instead of freeing the memory.
func isEven(num int) *bool {
	result := num%2 == 0
	// return the address of the result
	return &result
}

// Named return will return the function with the same name as defined in the declaration
func isOdd(num int) (result bool) {
	result = num%2 != 0
	return
}

// Multiple return values
func divideFunction(numA, numB float64) (float64, error) {
	if numB == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return numA / numB, nil
}

// Methods
type Cat struct {
	name     string
	greeting string
	age      int
}

func (greeter Cat) speak() {
	fmt.Println(greeter.greeting)
}
