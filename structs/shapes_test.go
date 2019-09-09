package structs

import (
   "testing"
)
func TestPerimeter(t *testing.T) {
   rectangle := Rectangle{10.0, 10.0}
   got := rectangle.Perimeter()
   want := 40.0

   if got!= want {
	   t.Errorf("got %.2f want %.2f", got, want)
   }
}

// Methods are called by invoking them on an instance of a particular type

 //test cases that can be tested in the same manner.
func TestArea(t *testing.T) {

   areaTests := []struct {
       shape Shape
       want  float64
   }{
      // assertions of truth about shapes and their areas
      {shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
      {shape: Circle{Radius: 10}, want: 314.1592653589793},
      {shape: Triangle{Base: 12, Height: 6}, want: 36.0},
   }
// The blank identifier "_," provides a way to ignore 
// left-hand side values in an assignment.

   for _, tt := range areaTests {
       got := tt.shape.Area()
       if got != tt.want {
           t.Errorf("got %.2f want %.2f", got, tt.want)
       }
   }

}

/*
created an interface with "checkArea"
interfaces alllow to make functions 
that can be used with different types and 
create highly-independent code while still 
maintaining the same type
*/