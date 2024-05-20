package enums_test

import "fmt"

// Define an enumeration of days of the week.
type Day int

const (
    Sunday Day = iota  // iota starts at 0
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)

// A function that returns whether a day is a weekday or not.
func (d Day) IsWeekday() bool {
    switch d {
    case Saturday, Sunday:
        return false
    default:
        return true
    }
}

// A method to convert a day to its string representation.
func (d Day) String() string {
    names := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
    return names[d]
}

// EnumExample is a function to demonstrate enum usage in Go.
func EnumExample() {
    today := Tuesday
    fmt.Println("Today is:", today)

    if today.IsWeekday() {
        fmt.Println(today, "is a weekday.")
    } else {
        fmt.Println(today, "is a weekend.")
    }
}