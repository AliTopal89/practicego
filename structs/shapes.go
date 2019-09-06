//A struct is a sequence of named elements, 
//called fields, each of which has a name and a type.

package structs

// Defining a struct type 
type Rectangle struct {
    Width float64
    Height float64
}

func Perimeter(rectangle Rectangle) float64 {
   return 2 *(rectangle.Height + rectangle.Width)	
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Height * rectangle.Width
}
