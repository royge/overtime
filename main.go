package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

const (
	timeFormat = "15:04"
)

func calculate(start, end, workHours string) (time.Duration, error) {
	since, err := time.Parse(timeFormat, start)
	if err != nil {
		return 0, err
	}

	until, err := time.Parse(timeFormat, end)
	if err != nil {
		return 0, err
	}

	duration := until.Sub(since)

	if workHours != "0" {
		w, _ := time.ParseDuration(workHours)
		duration -= w
	}

	return duration, nil
}

func main() {
	start := flag.String("since", "00:01", "Start time.")
	end := flag.String("until", "23:59", "End time.")
	workHours := flag.String("work", "9h30m", "Working hours including lunch break.")

	flag.Parse()

	duration, err := calculate(*start, *end, *workHours)
	if err != nil {
		log.Fatalf("can determine overtime: %v", err)
		return
	}

	fmt.Printf("Overtime: %s", duration.String())
}
