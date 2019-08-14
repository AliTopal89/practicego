package arrays

import (
	"testing"
    "reflect"
)

func TestCities(t *testing.T) {

    t.Run("slice string tests", func(t *testing.T) {
        cities := []string{"New York", "London", "Chicago", "Beijing", "Delhi", "Mumbai", "Bangalore", "Hyderabad", "Hong Kong"}

		got := []string{"Delhi", "Mumbai", "Bangalore", "Hyderabad", "Hong Kong"}
		want := cities[4:]

        if !reflect.DeepEqual(got, want) {
            t.Errorf("got %v want %v", got, want)
        }
	})
}