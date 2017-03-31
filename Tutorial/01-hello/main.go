package main

import (
	"fmt"
	"time"
	"math/rand"
)

// Function that add two value
// @params : integer, integer
// @return : integer - the two params summed
func add (x int, y int) int {
	return x + y
}

//Function that print out the value of two variable by passing the reference at the top
// @params : integer
// @return : integer, integer - one param added by itself and other part multiply by itself
func processValue (value int) (x, y int) {
	x = value + value
	y = value * value
	return
}

//Main function
func main() {
	fmt.Print("Hello Go!\n")
	fmt.Println("Right now is", time.Now())

	rand.Seed(time.Now().UTC().UnixNano())
	randomNumber := rand.Intn(30)
	fmt.Printf("This is your lucky number %d. \n", randomNumber)
	fmt.Print("Now add a 5 and the result will be: ")
	fmt.Println(add(randomNumber, 5))

	fmt.Print("Passing a 7 to adding/multiply : ")
	fmt.Println(processValue(7))
}