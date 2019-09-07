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

func TestArea(t *testing.T) {
   checkArea := func(t *testing.T, shape Shape, want float64) {
      t.Helper()
      got := shape.Area()
      if got != want {
         t.Errorf("got %.2f want %.2f", got, want)
      }
   }
   t.Run("rectangles", func(t *testing.T) {
       rectangle := Rectangle{12, 6}
       checkArea(t, rectangle, 72.0)
   })

   t.Run("circles", func(t *testing.T) {
       circle := Circle{10}
       checkArea(t, circle, 314.1592653589793)
   })

}

/*
created an interface with "checkArea"
interfaces alllow to make functions 
that can be used with different types and 
create highly-independent code while still 
maintaining the same type
*/