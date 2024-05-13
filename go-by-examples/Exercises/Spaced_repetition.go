package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/bastengao/chinese-holidays-go/holidays"
)
var suggested_interval []int = []int{1, 6, 14, 30, 66, 150}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the spaced repetition calculator.")
	dt_today :=  time.Now().Local()
	var intervals = suggested_interval 	
	fmt.Println("today's date: ", dt_today)
	fmt.Println("Getting the holidays data...")
	queryer, err := holidays.BundleQueryer()
	if err != nil {
		panic(err)
	}
	fmt.Println("Holidays data loaded.")
	// calculate days to recap given the specified date, for 5 iterations, using slices.
for {
		// 1. get the interval user wants to use
		fmt.Println("Enter the date you first learned the material, in the format YYYY-MM-DD: ")
		start_date_string, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error in your input: ", err)
		}
		// Trim the newline character from input depending on your OS
		start_date_string= start_date_string[:len(start_date_string)-1]  // For Unix-based systems
		// Parse the date from string
		start_date, err := time.Parse("2006-01-02", start_date_string)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			return
		}

		days_to_recap := calculate_days_to_recap(start_date, intervals, queryer)
		for i, day := range days_to_recap {
			fmt.Println("recap on: ", i+1, day.Format("2006-01-02"))
		}
}
}	

func calculate_days_to_recap(start_date time.Time, intervals []int, queryer holidays.Queryer) []time.Time{
	// Calculate the days to recap given the specified date, for 5 iterations, using slices.
	var res []time.Time
	for _, interval := range intervals {
		day_to_recap := start_date.AddDate(0, 0, interval)
		// skip weekends and chinese holidays
		
		for day_to_recap.Weekday() == time.Saturday || day_to_recap.Weekday() == time.Sunday || is_Holiday(queryer, day_to_recap) {
			day_to_recap = day_to_recap.AddDate(0, 0, 1)
			continue
		}
		res = append(res, day_to_recap)
	}
	return res
}

func is_Holiday(queryer holidays.Queryer, day_to_recap time.Time) bool {
	is_Holiday, err := queryer.IsHoliday(day_to_recap)
	if err != nil {
		fmt.Println("Error checking holiday:", err, "Defaulting to not a holiday.")
		is_Holiday = false
	}
	return is_Holiday
}