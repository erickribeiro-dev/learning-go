package main

import (
	"fmt"
	"math"
)

func basic_data_types() {
	// Signed integers
	var a int = 0                      // 32 bits or 64 bits depending on the system
	var b int8 = -128                  // -128 to 127
	var c int16 = -32768               // -32768 to 32767
	var d int32 = -2147483648          // -2147483648 to 2147483647
	var e int64 = -9223372036854775808 // -9223372036854775808 to 9223372036854775807
	fmt.Println("a, b, c, d and e:", a, b, c, d, e)

	// Unsigned integers
	var f uint = 0                      // 32 bits or 64 bits depending on the system
	var g uint8 = 255                   // 0 to 255
	var h uint16 = 65535                // 0 to 65535
	var i uint32 = 4294967295           // 0 to 4294967295
	var j uint64 = 18446744073709551615 // 0 to 18446744073709551615
	fmt.Println("f, g, h, i and j:", f, g, h, i, j)

	// Floats
	var k float32 = math.Pi // 32 bits (4 bytes)
	var l float64 = math.Pi // 64 bits (8 bytes)
	fmt.Println("k and l:", k, l)

	// Complex numbers
	var m complex64 = 1 + 2i   // Both real and imaginary part are float32
	var n complex128 = -3 + 4i // Both real and imaginary part are float64
	o := complex(5, 6)
	fmt.Println("m, n and o:", m, n, o)

	// Byte (alias for uint8). Used to represent ASCII characters
	// Rune (alias for int32). Used to represent Unicode characters
	var p byte = 'p'
	var q rune = 'q'
	var r = 'r' // infered as Rune (int32)
	fmt.Println("p, q and r:", p, q, r)
	fmt.Printf("type of p: %T, q: %T and r: %T\n", p, q, r)

	// Strings
	var s string = "Hello"
	var t string = `World`
	fmt.Println("s and t:", s, t)

	// Booleans
	var u bool = true
	var v bool = false
	fmt.Println("u and v:", u, v)
	// Using logical or/and relational operators results in a boolean value
	var w = 10 < 20       // true (10 is less than 20)
	var x = true && false // false (is true if all values are true)
	var y = true || false // true (is true if one of the values is true)
	var z = !true         // false (inverts the value)
	fmt.Println("w, x, y and z:", w, x, y, z)
}
