package main

import (
	"flag"
	"fmt"
	"time"
)

const (
	timeFormat = "15:04"
	workHours  = 9.5
)

func main() {
	start := flag.String("since", "00:01", "Start time.")
	end := flag.String("until", "23:59", "End time.")

	flag.Parse()

	since, err := time.Parse(timeFormat, *start)
	if err != nil {
		fmt.Println(err)
		return
	}

	until, err := time.Parse(timeFormat, *end)
	if err != nil {
		fmt.Println(err)
		return
	}

	duration := until.Sub(since)
	fmt.Println(duration.Hours() - workHours)
}
