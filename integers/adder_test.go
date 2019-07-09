// using %d vs %s. That's because we want it 
// to print an integer rather than a string.

/*
Godoc examples are snippets of Go code 
that are displayed as package documentation 
and that are verified by running them as tests.
*/
package integers

import "testing"
import	"fmt"

func TestAdder(t *testing.T) {
    sum := Add(2, 2)
    expected := 4

    if sum != expected {
        t.Errorf("expected '%d' but got '%d'", expected, sum)
    }
}

func ExampleAdd() {
    sum := Add(1, 5)
    fmt.Println(sum)
    // Output: 6
}

/*
Please note that the example function will not 
be executed if you remove the comment "//Output: 6". 
Although the function will be compiled, it won't be executed.
*/
