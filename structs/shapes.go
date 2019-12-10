package structs
import (
	"math"
)

/* If we try to call an interface in this case 
with something that isn't a shape, 
then it will not compile
*/
type Shape interface {
	Area() float64
}

// Defining a struct type 
type Rectangle struct {
    Width float64
    Height float64
}

// func (receiverName RecieverType) MethodName(args)
// you get your referene to methods data via the 'receiverName' variable.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func(r Rectangle) Perimeter() float64 {
	return 2*(r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
    Base   float64
    Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height)/2
}

func Perimeter(rectangle Rectangle) float64 {
   return 2 *(rectangle.Height + rectangle.Width)	
}
