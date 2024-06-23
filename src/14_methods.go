package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Methods are functions with receivers. The receiver is a type that will use said methods.
// All methods on a given type should have either value or pointer receivers, but not a mixture of both.
// So even though a value is not being modified on this method, it still needs to be defined with a pointer argument
func (v *Vertex) Hypot() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Use pointer receivers to change the value properly
func (v *Vertex) SetToZeros() {
	v.X = 0
	v.Y = 0
}

func methods() {
	vert := Vertex{4, 3}
	fmt.Println(vert, vert.Hypot())
	vert.SetToZeros()
	fmt.Println(vert, vert.Hypot())
}
