package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {

	// The layout here has to be the format of the string that we're passing in
	const layout = "1/2/2006 15:04:05"

	t, e := time.Parse(layout, date)
	if e != nil {
		fmt.Println("*** ERROR: " + e.Error())
	}

	return t

}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {

	const layout = "January 2, 2006 15:04:05"

	t, e := time.Parse(layout, date)
	if e != nil {
		fmt.Println("*** ERROR: " + e.Error())
	}
	fmt.Println("*** Given time: " + t.GoString())
	fmt.Println("*** Current time: " + time.Now().GoString())

	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {

	const layout = "Monday, January 2, 2006 15:04:05"

	t, e := time.Parse(layout, date)
	if e != nil {
		fmt.Println("*** ERROR: " + e.Error())
	}

	return t.Hour() >= 12 && t.Hour() < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {

	const layout = "1/2/2006 15:04:05"

	t, e := time.Parse(layout, date)
	if e != nil {
		fmt.Println("*** ERROR: " + e.Error())
	}

	return fmt.Sprintf("You have an appointment on %s.", t.Format("Monday, January 2, 2006, at 15:04"))
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {

	// Opening Date
	openingDate := time.Date(2012, time.September, 15, 0, 0, 0, 0, time.UTC)

	// Anniversary Date
	return time.Date(time.Now().Year(), openingDate.Month(), openingDate.Day(), openingDate.Hour(), openingDate.Minute(), openingDate.Second(), openingDate.Nanosecond(), openingDate.Location())

}
