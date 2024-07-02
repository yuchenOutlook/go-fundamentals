package main

import (
	"fmt"
	"time"
)

func main() {
    p := fmt.Println

    now := time.Now()
    p(now)

    then := time.Date(
        2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
    p(then)

    p(then.Year())
    p(then.Month())
    p(then.Day())
    p(then.Hour())
    p(then.Minute())
    p(then.Second())
    p(then.Nanosecond())
	p("Location if not specified when creating time, it will be UTC: ")
    p(then.Location())

    p(then.Weekday())

    p(then.Before(now))
    p(then.After(now))
    p(then.Equal(now))

    diff := now.Sub(then)
	p("The diff between now and then:")
    p(diff)

    p("which is ", diff.Hours(), " hours")
    p(diff.Minutes(), " minutes")
    p(diff.Seconds(), " Seconds")
    p(diff.Nanoseconds(), "Nano seconds")

	p("then added the diff: ")
    p(then.Add(diff))
	p("then minused the diff: ")
    p(then.Add(-diff))
}