package main

import (
	"fmt"
	"math"
)


type Powerer interface {
	Power() int
}

type Vertex struct {
	X, Y int
}

type Myint int


func interfaceAssignment() {
	var a Powerer

	f := Myint(2)
	v := Vertex{3, 4}

	a = f
	fmt.Println(a.Power())

	a = &v
	fmt.Println(a.Power())
}


func (f Myint) Power() int {
	return int(math.Pow(float64(f), 2))
}

func (v *Vertex) Power() int {
	return int(math.Pow(float64(v.X), float64(v.Y)))
}


/////////////////////////////////////////////////////////////

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("this is a int: %v\n", v)
	case string:
		fmt.Printf("this is a string: %q\n", v)
	default:
		fmt.Printf("Still not implemented type %T\n", v)
	}
}
/////////////////////////////////////////////////////////////

type IPAddr [4]byte

func (p IPAddr) String() string{
	return fmt.Sprintf("%d.%d.%d.%d",p[0], p[1], p[2], p[3])
}

func ipAddress() {

	hosts := map[string]IPAddr{
		"loopback": {127, 0, 0, 1},
		"googleDns": {8, 8, 8 ,8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

func main() {
	interfaceAssignment()
	fmt.Println(); fmt.Println()
	do(123); do("ciao"); do(false)
	fmt.Println(); fmt.Println()
	ipAddress()
    	//fmt.Println(); fmt.Println();
	//typeAssertion();
	//fmt.Println(); fmt.Println();
	//typeDo(1); typeDo("ciao"); typeDo(1.2);

}
