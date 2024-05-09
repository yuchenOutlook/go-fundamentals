package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)
var suggested_interval []int = []int{1, 3, 7, 14, 28}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the spaced repetition calculator.")
	dt_today :=  time.Now().Local()
	
	fmt.Println("today's date: ", dt_today)
	// calculate days to recap given the specified date, for 5 iterations, using slices.
	
	// 1. get the interval user wants to use
	fmt.Println("Enter how many months you want to keep your memory to last...(1, 3, 7, 14, 28)")
	fmt.Print("Enter your choice: ")
	memory_length_in_month, _ := reader.ReadString('\n')
	memory_length_in_month_int, err := strconv.Atoi(memory_length_in_month)
	if err != nil {
		fmt.Println("Error in your input: ", err)
	}

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