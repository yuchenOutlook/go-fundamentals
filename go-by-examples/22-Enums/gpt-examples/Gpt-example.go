package enums_test

import "fmt"

// Define an enumeration of days of the week.
type Day int
type DriverType int

const (
    NormalDriver DriverType = iota << 1
    HosDriver
    ELDDriver
    AutopilotDriver
)

const (
    Sunday Day = iota  // iota starts at 0
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)

func (dt DriverType) String() string {
    names := map[DriverType]string{
        NormalDriver: "NormalDriver",
        HosDriver: "HosDriver",
        ELDDriver: "ELDDriver",
        AutopilotDriver: "AutopilotDriver",
    }
    return names[dt]
}

// A function that returns whether a day is a weekday or not.
func (d Day) IsWeekday() bool {
    switch d {
    case Saturday, Sunday:
        return false
    default:
        return true
    }
}

func (dt DriverType) IsAutopilot() bool {
    switch dt {
    case AutopilotDriver:
        return true
    default:
        return false
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

    new_driver := AutopilotDriver
    if new_driver.IsAutopilot() {
        fmt.Println(new_driver, "is an autopilot driver.")
    }

    normal_driver := NormalDriver
    if !normal_driver.IsAutopilot() {
        fmt.Println(normal_driver, "is not an autopilot driver.")
    }

    if today.IsWeekday() {
        fmt.Println(today, "is a weekday.")
    } else {
        fmt.Println(today, "is a weekend.")
    }
}