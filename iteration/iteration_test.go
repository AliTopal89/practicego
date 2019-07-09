// To do stuff repeatedly in Go, you'll need for. 
// In Go there are no "while", "do", "until"
package iteration

import "testing"

func TestRepeat(t *testing.T) {
    repeated := Repeat("a")
    expected := "aaaaa"

    if repeated != expected {
        t.Errorf("expected '%s' but got '%s'", expected, repeated)
    }
}
