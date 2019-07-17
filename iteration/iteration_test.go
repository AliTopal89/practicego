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

/* 
When the benchmark code is executed, 
it runs b.N times and measures how long it takes.
*/

func BenchmarkRepeat(b *testing.B) {
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        Repeat("aaaaaaaaaaaa")  
    }
}
