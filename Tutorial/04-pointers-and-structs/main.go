package main

import (
	"fmt"
)

func pointers() {
	x, y := 1, 2
	var pointX, pointY *int

	//assign to pointN the area of memory of the two variables (generates a pointer to its operand)
	pointX = &x
	pointY = &y

	//Print the area of memory that we are pointing at (for x)
	fmt.Println(pointX)
	//Print the value in that memory area (for y)
	fmt.Println(*pointY)

	//Arithmetic on the pointed value is OK
	*pointX = *pointX + 10
	fmt.Println(*pointX)
	fmt.Println(x)

	//Pointer Arithmetic is not possible!
	//pointX = pointX+1
	//it gonna return (mismatched types *int and int)
}

func structs() {
	type Vertex struct {
		X int
		Y int
	}

	fmt.Println()
	fmt.Println()
	fmt.Println("Playing with structs")
	fmt.Println(Vertex{1, 2})
	v := Vertex{3, 5}
	fmt.Println("Display X of the vertex")
	fmt.Println(v.X)
	fmt.Println("Display Y of the vertex")
	fmt.Println(v.Y)

	fmt.Println()
	fmt.Println()

	fmt.Println("Using pointer on struct???")
	var pointStruct *Vertex
	pointStruct = &v
	fmt.Println("Using pointer notation (*pointerName).attribute")
	fmt.Println((*pointStruct).X, (*pointStruct).Y)
	fmt.Println("Using pointer special notation for struct pointerName.attribute")
	fmt.Println(pointStruct.X, pointStruct.Y)
}

func main() {
	pointers()
	structs()
}
