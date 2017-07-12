package main

import (
	"fmt"
	"math/cmplx"
)

var (
	aBool  bool       = true
	MaxInt uint64     = 1<<64 - 1 //Shifting left of 63 bit
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	//zero value variable (variable not declared)
	var i int
	var f float32
	var b bool
	var s string

	//Long variable declaration
	var longVarDeclaration int = 2

	//Short variable declaration
	shortVarDeclaration := 2

	fmt.Println("---- Play with Variables ----")
	fmt.Println("Zero Value Variable: ")
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	fmt.Println("Long Variable: ")
	fmt.Println(longVarDeclaration)

	fmt.Println("Short Variable: ")
	fmt.Println(shortVarDeclaration)

	fmt.Println("Play with cool variables and constant")
	const format = "%T(%v)\n"
	//where f is the format
	fmt.Printf(format, aBool, aBool)
	fmt.Printf(format, MaxInt, MaxInt)
	fmt.Printf(format, z, z)

}
