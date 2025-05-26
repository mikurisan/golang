package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()                        // get current local time.
	tick := time.Tick(100 * time.Millisecond)  // send a tick at the specifed time duration frequency.
	boom := time.After(500 * time.Millisecond) // send the current time once after the specified time duration.
	elapsed := func() time.Duration {
		return time.Since(start).Round(time.Millisecond) // calculates the total time from `start` time point to the current time.
	}
	for {
		select {
		case <-tick:
			fmt.Printf("[%6s] tick.\n", elapsed())
		case <-boom:
			fmt.Printf("[%6s] BOOM!\n", elapsed())
			return
		default:
			fmt.Printf("[%6s]     .\n", elapsed())
			time.Sleep(50 * time.Millisecond)
		}
	}
}
