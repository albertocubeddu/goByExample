package main

import (
	"fmt"
	"strings"
)

func arrays() {
	//Array has a FIXED size.
	var array [2]string
	array[0] = "Hi"
	array[1] = "Golang"

	fmt.Println(array)
	fmt.Println(array[1], array[0])

	//Slices are reference to the original array so if you change the slice you change the original array
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func slices() {
	//Slices are reference to the original array so if you change the slice you change the original array
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var slice []int = primes[1:3]
	fmt.Println(slice)

	slice[0] = 1
	fmt.Println(slice)
	fmt.Println(primes)

	//Slice literals
	sliceLiterals := []int{1, 2, 3, 4, 5}
	fmt.Println(sliceLiterals)

	//Slice literals with structs

	vertexStruct := []struct {
		X int
		Y bool
	}{
		{0, true},
		{1, false},
		{2, true},
		{3, true},
	}

	fmt.Println(vertexStruct)

	//panic: runtime error: slice bounds out of range
	//vertexStruct = vertexStruct[1:5]
	//fmt.Println(vertexStruct)

	vertexStruct = vertexStruct[1:4]
	fmt.Println(vertexStruct)

	//left omited -> default min limit
	vertexStruct = vertexStruct[:2]
	fmt.Println(vertexStruct)

	//right omitted -> default max limit
	vertexStruct = vertexStruct[1:]
	fmt.Println(vertexStruct)
}

func slicesMake() {
	//Length of the slice and Capacity of the slice
	slice := make([]int, 5)
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice))

	sliceTwo := make([]int, 2, 5)
	fmt.Println(sliceTwo)
	fmt.Println(len(sliceTwo), cap(sliceTwo))
}

func sliceOfSlice() {
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[2][0] = "O"
	board[1][1] = "X"
	board[1][0] = "X"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func sliceAppending() {
	var slice []int
	fmt.Println(slice)

	slice = append(slice, 123)
	fmt.Println(slice)

	slice = append(slice, 1, 2, 43)
	fmt.Println(slice)

	//cannot append different type! cannot use "asD" (type string) as type int in append
	//slice = append(slice, "asD")
	//fmt.Println(slice)

}

func sliceExample(dx, dy int) [][]uint8 {
	//return slice of lenght dy
	fakeImage := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		fakeImage[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			fakeImage[y][x] = uint8((x + y) / 2)
		}
	}
	fmt.Println(fakeImage)

	return fakeImage
}

func maps() {
	type Vertex struct {
		x, y int
	}

	var m map[string]Vertex

	//First way to create a map
	m = make(map[string]Vertex)
	m["test"] = Vertex{1, 2}
	m["test2"] = Vertex{3, 4}

	//or short way
	m2 := map[string]Vertex{
		"test":  Vertex{1, 2},
		"test2": Vertex{1, 2},
	}
	fmt.Println(m2)

	fmt.Println(m)
	fmt.Println(m["test"])

	v, ok := m["test3"]
	fmt.Print("The Value:", v, "Is here? ", ok, "\n")
	v, ok = m["test2"]
	fmt.Print("The Value:", v, "Is here? ", ok, "\n")

	delete(m, "test2")
	fmt.Print("The Value:", v, "Is here? ", ok, "\n")
}

func mapsExample(s string) map[string]int {

	ss := strings.Fields(s)
	num := len(ss)
	ret := make(map[string]int)
	for i := 0; i < num; i++ {
		ret[ss[i]]++
	}
	fmt.Println(ret)
	return ret
}

func fibonacci() func() int {
	counter, first, second := 0, 0, 1
	//0, 0, 1 -> 0, 1, 1
	//0, 1, 1 -> 1, 1, 2
	//1, 1, 2 -> 1, 2, 3
	//1, 2, 3 -> 2, 3, 5
	//2, 3, 5 -> 3, 5, 8
	//etc.
	return func() int {
		counter, first, second = first, second, first+second
		return counter
	}
}

func runFibonacci() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func main() {
	arrays()
	fmt.Println()
	fmt.Println()
	slices()
	fmt.Println()
	fmt.Println()
	slicesMake()
	fmt.Println()
	fmt.Println()
	sliceOfSlice()
	fmt.Println()
	fmt.Println()
	sliceAppending()
	fmt.Println()
	fmt.Println()
	sliceExample(4, 4)
	fmt.Println()
	fmt.Println()
	maps()
	fmt.Println()
	fmt.Println()
	mapsExample("ciao io sono ciao ")
	fmt.Println()
	fmt.Println()
	runFibonacci()
}
