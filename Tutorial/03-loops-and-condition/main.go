package main

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
	"time"
)

func loops() {
	var sumLoop int
	sumLoopContinued := 1
	sumWhile := 1

	//for loop
	for i := 0; i < 10; i++ {
		sumLoop += i
	}
	fmt.Println(sumLoop)

	//For Continued
	for sumLoopContinued < 10 {
		sumLoopContinued += sumLoopContinued
	}
	fmt.Println(sumLoopContinued)

	//While
	for sumWhile < 1000 {
		sumWhile += sumWhile
	}
	fmt.Println(sumWhile)
}

func switches() {
	fmt.Println("You are running on : ")
	//Normal switch
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("You Rock")
	default:
		fmt.Println("%s. Switch to linux", os)
	}

	//"True" switch
	fmt.Println("It's time to: ")
	t := time.Now()
	switch {
	case t.Hour() < 10:
		fmt.Println("Coffee Time")
	case t.Hour() < 17:
		fmt.Println("Lunch Time")
	case t.Hour() < 23:
		fmt.Println("Beer O'clock")
	default:
		fmt.Println("Boring Time")
	}
}

func deferrer() {
	defer fmt.Println("I'm executing after the defer() has finished is execution")
	fmt.Println("I'm executing during the defeer() scope")

	//LIFO cycle of defer
	//Defer call are stacked (Last In First Out)
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Uhh the countdown will be launched soon!!")
}

func sliceLoops() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2^%d = %d = %s \n", i, v, strconv.FormatInt(int64(v), 2))
	}

	fmt.Println()
	fmt.Println()

	powTwo := make([]int, 10)
	for i := range pow {
		powTwo[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range powTwo {
		fmt.Printf("%d\n", value)
	}
}

func sqrt(x float64) string {
	//if statement
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	//if short statement (v is alive just in the IF scope)
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func ifes() {
	sqrt(3)
	pow(2, 2, 10)
}

func main() {
	loops()
	fmt.Println()
	fmt.Println()
	switches()
	fmt.Println()
	fmt.Println()
	deferrer()
	fmt.Println()
	fmt.Println()
	sliceLoops()
	fmt.Println()
	fmt.Println()
	ifes()
}
