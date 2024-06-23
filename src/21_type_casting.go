package main

import (
	"fmt"
	"log"
	"strconv"
)

func type_casting() {
	// Type casting means converting a variable from one type to another type.

	// Float to integer
	f := 1000.55
	i := int(f)
	fmt.Println(f, i)

	// Integer to float
	fmt.Println(float64(2000))

	// String to integer
	converted_to_integer, err := strconv.Atoi("3000")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(converted_to_integer)

	// Integer to string
	fmt.Println(strconv.Itoa(4000))

	// String to byte
	b := []byte("Hello Zoë!")
	fmt.Println(b)

	// String to rune
	c := []rune("Hello Zoë!")
	fmt.Println(c)

	// Bytes to string
	st := string(b)
	fmt.Println(st)

	// During division, types are converted implicitly.
	// If either the dividend or divisor are of type float, the result will be a float number
	fmt.Printf("Division between integers results in an integer: %v of type %T\n", 5/2, 5/2)
	fmt.Printf("Division containing a float results in a float: %v of type %T\n", 5.0/2, 5.0/2)
}
