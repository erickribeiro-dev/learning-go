package main

import "fmt"

func functions() {
	// Functions (Reference Type)
	// Calling a function with a parameter (argument) is similar to most programming languages
	for i := 0; i < 5; i++ {
		say_message("Hello Go!", i)
	}

	// If all arguments have the same type,
	// you only need to specify the type once in the function declaration
	say_greeting("Hello", "Peter")

	// You can manipulate the original parameter if you pass it as a pointer as argument
	// Remember that maps and slices are always passed in as pointers
	name := "John"
	say_hello_and_alter_name(&name)
	fmt.Println(name)

	// Variadic parameters (maximum 1 variadic parameter and must be the last listed in the arguments)
	// the type of the variadic parameter is a slice
	sum_function("The sum is:", 1, 3, 5, 7, 9)

	// Return values
	fmt.Println(get_length("This is a 36 characters long message"))

	// Returning a pointer
	fmt.Println(*is_even(88))

	// Named return
	fmt.Println(is_odd(77))

	// multiple return values are stored in multiple variables
	divide_result, err := divide_function(4.0, 1.5)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(divide_result)

	// Anonymous functions
	message := "This text was printed inside an anonymous function"
	func(msg string) {
		fmt.Println(msg)
	}(message)

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
func say_message(msg string, index int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is:", index)
}

// All arguments of the same type
func say_greeting(msg, name string) {
	fmt.Println(msg, name)
}

// Pointers as arguments
func say_hello_and_alter_name(name *string) {
	fmt.Println(*name)
	*name = "Scott"
	fmt.Println("Your name is now", *name)
}

// Variadic parameters
// the type of the variadic parameter is a slice
func sum_function(msg string, args ...int) {
	fmt.Printf("%v %T\n", args, args)
	result := 0
	for _, val := range args {
		result += val
	}
	fmt.Println(msg, result)
}

// Return values
func get_length(msg string) int {
	return len(msg)
}

// You can return a pointer to a variable declared inside a function.
// Unlike most programming languages, Go will recognize that you're returning a pointer and
// keep the variable in the heap (shared) memory instead of freeing the memory.
func is_even(num int) *bool {
	result := num%2 == 0
	// return the address of the result
	return &result
}

// Named return will return the function with the same name as defined in the declaration
func is_odd(num int) (result bool) {
	result = num%2 != 0
	return
}

// Multiple return values
func divide_function(numA, numB float64) (float64, error) {
	if numB == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return numA / numB, nil
}

// Methods
type Cat struct {
	name     string
	greeting string
	age      int
}

// greeter is a value receiver. Pointer receivers can also be used
func (greeter Cat) speak() {
	fmt.Println(greeter.greeting)
}
