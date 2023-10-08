package main

import (
	"fmt"
	"reflect"
)

func dataTypes() {
	// String
	var a string = "This is a string"

	// Getting the type of the variable
	fmt.Printf("%T\n", a)
	fmt.Println(reflect.TypeOf(a))

	// Boolean
	var b bool = true
	var c bool = false

	// Signed Integer
	var d int = -94                    // 32-bit or 64-bit depending on the system
	var e int8 = 127                   // 8 bit
	var f int16 = -32768               // 16 bit
	var g int32 = 2147483647           // 32 bit
	var h int64 = -9223372036854775808 // 64 bit

	// Unigned Integer
	var i uint = 0                      // 32-bit or 64-bit depending on the system
	var j uint8 = 255                   // 8 bit
	var k uint16 = 65535                // 16 bit
	var l uint32 = 4294967295           // 32 bit
	var m uint64 = 18446744073709551615 // 64 bit

	// Float
	var n float32 = 3.14159
	var o float64 = 2.71828

	fmt.Println(a, b, c)
	fmt.Println(d, e, f, g, h)
	fmt.Println(i, j, k, l, m)
	fmt.Println(n, o)
}
