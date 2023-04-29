package expenses

import (
	"fmt"
	"strings"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	filteredRecords := make([]Record, 0)

	for _, v := range in {
		if predicate(v) {
			filteredRecords = append(filteredRecords, v)
		}
	}

	return filteredRecords
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		return r.Day >= p.From && r.Day <= p.To
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return strings.EqualFold(r.Category, c)
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	filteredRecords := Filter(in, ByDaysPeriod(p))
	var result float64
	for _, v := range filteredRecords {
		result += v.Amount
	}

	return result
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	filteredRecordsByCategory := Filter(in, ByCategory(c))
	if len(filteredRecordsByCategory) == 0 {
		return 0, fmt.Errorf("no records for the given category: %s", c)
	}

	filteredRecordsByPeriod := Filter(filteredRecordsByCategory, ByDaysPeriod(p))

	var result float64
	for _, v := range filteredRecordsByPeriod {
		result += v.Amount
	}

	return result, nil
}
