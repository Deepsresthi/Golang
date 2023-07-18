package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Sleeper is an interface that defines the Sleep method.
type Sleeper interface {
	Sleep()
}

// DefaultSleeper is an implementation of the Sleeper interface using time.Sleep().
type DefaultSleeper struct{}

// Sleep pauses the execution for 1 second.
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const finalWord = "Go!"
const countdownStart = 3

// Countdown prints a countdown from countdownStart to 1, with a delay between each count.
// It takes an io.Writer to write the output and a Sleeper interface to handle the delay.
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i) // Print the current countdown number
		sleeper.Sleep()      // Pause the execution for the specified duration using the provided Sleeper implementation
	}
	fmt.Fprint(out, finalWord) // Print the final word "Go!"
}

func main() {
	sleeper := &DefaultSleeper{} // Create an instance of DefaultSleeper
	Countdown(os.Stdout, sleeper) // Call Countdown with os.Stdout as the output and the DefaultSleeper as the Sleeper
}
