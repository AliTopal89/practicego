
/*
created an interface with "checkArea"
interfaces alllow to make functions 
that can be used with different types and 
create highly-independent code while still 
maintaining the same type
*/
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

 // Table driven tests are useful when you want to build a list of
 //test cases that can be tested in the same manner.
 func TestArea(t *testing.T) {

   areaTests := []struct {
       name    string
       shape   Shape
       hasArea float64
   }{
       /*
       What we’d like to to set up all the inputs and expected outputs 
       and feel them to a single test harness - "tests := []test{name rectangle..."
       This is a great time to introduce table driven testing.
       
       we can use an anonymous struct literal to reduce the boilerplate

       you can name the fields of instances of structs
       */ 
       {name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
       {name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
       {name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
   }

   for _, tt := range areaTests {
       // using tt.name from the case to use it as the `t.Run` test name
       t.Run(tt.name, func(t *testing.T) {
           got := tt.shape.Area()
           if got != tt.hasArea {
               t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasArea)
           }
       })

   }

}


