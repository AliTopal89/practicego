/*
An interface type consists of a set of method signatures.
A variable of interface type can hold any value that
implements these methods.
*/

package main

import (
	"fmt"
	"strconv"
)

// // MyStringer is the interface that wraps the basic String method.
type MyStringer interface {
	String() string
}

type Temp int

func (t Temp) String() string {
	return strconv.Itoa(int(t)) + " C"
}

type Point struct {
	x, y int
}

/*
 Both Temp and *Point implement the MyStringer interface
 because both are using String() method.
*/

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

var x MyStringer

/*
  when you call a method on an interface value (this is "x"),
  the method of its underlying type string is executed
*/
func main() {
	x = Temp(24)
	fmt.Println(x.String())
	// & - variable's memory address.
	x = &Point{1, 2}
	fmt.Println(x.String())
}
