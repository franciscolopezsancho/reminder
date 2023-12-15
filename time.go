package main

import (
	"fmt"
	"math"
	"time"

	"google.golang.org/api/calendar/v3"
)

func dates(t time.Time) []calendar.EventDateTime {
	var events []calendar.EventDateTime

	for i := 0; i < 20; i++ {
		numberDays := int(math.Pow(2, float64(i)))
		if (numberDays > 90) numberDays = 90
		events = append(events, calendar.EventDateTime{
			Date: toString(t.Add(time.Duration(int(time.Hour) * 24 * numberDays))),
		})
	}
	return events
}

func toString(t time.Time) string {
	return fmt.Sprintf("%d-%d-%d", t.Year(), t.Month(), t.Day())
}
