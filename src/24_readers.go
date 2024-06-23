package main

import (
	"fmt"
	"io"
	"strings"
)

func readers() {
	r := strings.NewReader("Hello, Reader!")
	fmt.Printf("%T", r)
	// strings.Reader will consume 8 bytes at a time
	b := make([]byte, 8) // [0 0 0 0 0 0 0 0]
	for {
		// Read populates the given byte slice with data and returns the number of bytes populated and an error value.
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF { // It returns an io.EOF error when the stream ends.
			break
		}
	}
}
