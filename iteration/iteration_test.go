package iteration

import (
    "testing"
    "fmt"
)

func TestRepeat(t *testing.T) {
    repeated := Repeat("a", 5)
    expected := "aaaaa"

    if repeated != expected {
        t.Errorf("expected '%s' but got '%s'", expected, repeated)
    }
}

func BenchmarkRepeat(b *testing.B) {
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        Repeat("a", 20)  
    }
}

func ExampleRepeat(t *testing.T){
    repeated := Repeat("a", 4)
    fmt.Println(repeated)


}