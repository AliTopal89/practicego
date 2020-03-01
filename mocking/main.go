package main

import (
	"fmt"
	"io"
	"os"
)

// use general purpose interface instead of *bytes.Buffer
// Use a for loop counting backwards with i--
func Countdown(out io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprintf(out, "Go!")

}

func main() {
	Countdown(os.Stdout)
}
