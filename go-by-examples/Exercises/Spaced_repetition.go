package main

import (
	"fmt"
	"time"
)
var suggested_interval []int = []int{1, 3, 7, 14, 28}

func main() {
	fmt.Println("Welcome to the spaced repetition calculator.")
	dt_today :=  time.Now().Local()
	
	fmt.Println("today's date: ", dt_today)
	// calculate days to recap given the specified date, for 5 iterations, using slices.
	var intervals = suggested_interval 
	days_to_recap := calculate_days_to_recap(dt_today, intervals)
	for i, day := range days_to_recap {
		fmt.Println("recap on: ", i+1, day.Format("2006-01-02"))
	}
}

func calculate_days_to_recap(start_date time.Time, intervals []int) []time.Time{
	// Calculate the days to recap given the specified date, for 5 iterations, using slices.
	var res []time.Time
	for _, interval := range intervals {
		day_to_recap := start_date.AddDate(0, 0, interval)
		// skip weekends
		if day_to_recap.Weekday() == time.Saturday || day_to_recap.Weekday() == time.Sunday {
			continue
		}
		res = append(res, day_to_recap)
	}
	return res
}