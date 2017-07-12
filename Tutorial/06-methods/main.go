package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func methodType() float64 {
	v := Vertex{5, 6}
	return v.Abs()
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func methodPointer() {
	v := Vertex{2, 4}
	v.Scale(2)
	println(v.X, v.Y)
}

func main() {
	println(methodType())
	fmt.Println()
	fmt.Println()
	methodPointer()

}

//INTERFACE PAGE 9!!!!
