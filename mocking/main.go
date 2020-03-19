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

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// use general purpose interface instead of *bytes.Buffer
// Use a for loop counting backwards with i--
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Printf("Current Time: %v\n", time.Now())
		sleeper.Sleep()
		fmt.Fprintln(out, i)
		fmt.Printf("Current Time: %v\n", time.Now())
	}
	sleeper.Sleep()
	fmt.Fprintf(out, finalWord)

}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
