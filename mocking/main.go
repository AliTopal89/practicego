package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"

// use general purpose interface instead of *bytes.Buffer
// Use a for loop counting backwards with i--
func Countdown(out io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
		fmt.Printf("Current Time: %v\n", time.Now())
		time.Sleep(1 * time.Second)
		fmt.Printf("Current Time: %v\n", time.Now())
	}

	fmt.Fprintf(out, finalWord)

}

func main() {
	Countdown(os.Stdout)
}
