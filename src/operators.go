package main

import "fmt"

func operators() {
	// Arithmetic
	fmt.Println("Arithmetic:")
	fmt.Println(4 + 3)
	fmt.Println(4 - 3)
	fmt.Println(4 * 3)
	fmt.Println(4 / 3)
	fmt.Println(4 % 3)

	// Relational
	fmt.Println("Relational:")
	fmt.Println(4 > 3)
	fmt.Println(4 >= 3)
	fmt.Println(4 < 3)
	fmt.Println(4 <= 3)
	fmt.Println(4 == 3)
	fmt.Println(4 != 3)

	// Logical
	fmt.Println("Logical:")
	fmt.Println(true && false)          // true if both are true
	fmt.Println(false || false || true) // true if all are true
	fmt.Println(!false)                 // negation

	// Assignment
	x := 5
	x += 3 // x = x + 3
	x -= 3 // x = x - 3
	x *= 2 // x = x * 3
	x /= 2 // x = x / 3

	// Bitwise
	fmt.Println("Bitwise:")
	fmt.Println(21 & 19) // 10101 & 10011 equals 10001 (17)
	fmt.Println(21 | 19) // 10101 & 10011 equals 10111 (23)
	fmt.Println(21 ^ 19) // 10101 & 10011 equals 00110 (6)
	// Left shift and Right shift
	fmt.Println(13 << 2) // 001101 becomes 110100
	fmt.Println(13 >> 2) // 001101 becomes 000011. Note that bits are lost
}
