package main

import "fmt"

func operators() {
	// Arithmetic
	a, b := 4, 3
	fmt.Println("Arithmetic:")
	fmt.Println(a + b) // Sum
	fmt.Println(a - b) // Subtraction
	fmt.Println(a * b) // Multiplication
	fmt.Println(a / b) // Division
	fmt.Println(a % b) // Module

	// Relational
	fmt.Println("Relational:")
	fmt.Println(a > b)  // Greater than
	fmt.Println(a >= b) // Greater or equal than
	fmt.Println(a < b)  // Less than
	fmt.Println(a <= b) // Less or equal than
	fmt.Println(a == b) // Equal to
	fmt.Println(a != b) // Different than

	// Logical
	fmt.Println("Logical:")
	fmt.Println(true && false)          // true if both are true
	fmt.Println(false || false || true) // true if all are true
	fmt.Println(!false)                 // negation

	// Assignment
	x := 6
	x += 3 // x = x + 3
	x -= 3 // x = x - 3
	x *= 3 // x = x * 3
	x /= 3 // x = x / 3

	// Bitwise
	fmt.Println("Bitwise:")
	fmt.Println(21 & 19)  // 10101 AND 10011 equals 10001 (17)
	fmt.Println(21 | 19)  // 10101 OR 10011 equals 10111 (23)
	fmt.Println(21 ^ 19)  // 10101 XOR 10011 equals 00110 (6)
	fmt.Println(21 &^ 19) // 10101 AND NOT 10011 equals 00100 (4)

	// Left shift and Right shift
	fmt.Println(13 << 2) // 001101 becomes 110100
	fmt.Println(13 >> 2) // 001101 becomes 000011. Note that bits are lost
}
