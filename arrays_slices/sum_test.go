// %v	the value in a default format
// format strings: https://golang.org/pkg/fmt/

/* 
For example, the type [4]int represents 
an array of four integer values layed out sequentially.

Slicing does not copy the slice's data. 
It creates a new slice value that points to the original array. 
This makes slice operations as efficient as manipulating array indices. 
Therefore, modifying the elements (not the slice itself) of a re-slice 
modifies the elements of the original slice:
*/
package arrays

import (
	"testing"
	"reflect"
)

func TestSum(t *testing.T) {

    t.Run("collection of 5 numbers", func(t *testing.T) {
        numbers := []int{1, 2, 3, 4, 5}

        got := Sum(numbers)
        want := 15

        if got != want {
            t.Errorf("got %d want %d given, %v", got, want, numbers)
        }
    })

    t.Run("collection of any size", func(t *testing.T) {
        numbers := []int{1, 2, 3}

        got := Sum(numbers)
        want := 6

        if got != want {
            t.Errorf("got %d want %d given, %v", got, want, numbers)
        }
    })

}

// The tail of a collection is all the items apart from the first one (the "head")
func TestSumAllTails(t *testing.T) {
    checkSums := func(t *testing.T, got, want []int) {
        t.Helper()
        if !reflect.DeepEqual(got, want) {
            t.Errorf("got %v want %v", got, want)
        }
    }

    t.Run("make the sums of some slices", func(t *testing.T){
        got := SumAllTails([]int{1,2}, []int{0,9})
        want := []int{2, 9}
        checkSums(t, got, want)
    })

    t.Run("safely sum empty slices", func(t *testing.T){
        got := SumAllTails([]int{}, []int{3, 4, 5})
        want := []int{0, 9}
        checkSums(t, got, want)
    })
    
	t.Run("makes sums of some other slices", func(t *testing.T) {
		got := SumAllTails([]int{1,2}, []int{1, 2, 3})
		want := []int{2, 5}
		checkSums(t, got, want)
	})
}