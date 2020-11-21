package oo_approach

import (
	"fmt"
	"math"
)

/*
	Pointer Receiver vs Value Receiver
	you can pass pointer to functions and receive them only with one determine
	if you are going to use this, your inputs in func start with (*) ,
	and your calling functions passed variable with (&)
*/

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Pointer_And_Value_Receiver() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
