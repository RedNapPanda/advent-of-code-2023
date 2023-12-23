package aoc_util

import (
	"fmt"
	"time"
)

// Timer credits: https://stackoverflow.com/questions/45766572/is-there-an-efficient-way-to-calculate-execution-time-in-golang
//
// Learned more about defer and how you could use/abuse it like this
//
// timer returns a function that prints the name argument and
// the elapsed time between the call to timer and the call to
// the returned function. The returned function is intended to
// be used in a defer statement:
//
// defer timer("sum")()
func Timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
