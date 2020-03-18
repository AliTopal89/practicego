package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"

/*
 Countdown function would not be responsible for how long
 the sleep is. This simplifies our code a little
 for now at least and means a user of our function
 can configure that sleepiness however they like.
*/

//Sleep pauses the current goroutine for at least the duration "d"
type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

// use general purpose interface instead of *bytes.Buffer
// Use a for loop counting backwards with i--
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
		fmt.Printf("Current Time: %v\n", time.Now())
		sleeper.Sleep()
		fmt.Printf("Current Time: %v\n", time.Now())
	}
	sleeper.Sleep()
	fmt.Fprintf(out, finalWord)

}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
