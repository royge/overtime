package main

import (
	"flag"
	"fmt"
	"time"
)

const (
	timeFormat = "15:04"
)

func main() {
	start := flag.String("since", "00:01", "Start time.")
	end := flag.String("until", "23:59", "End time.")
	workHours := flag.String("work", "9h30m", "Working hours including lunch break.")

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

	if *workHours != "0" {
		w, _ := time.ParseDuration(*workHours)
		duration -= w
	}

	fmt.Println(duration.String())
}
