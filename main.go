package main

import (
	"fmt"
	"time"
)

func firstWeekDay(year int) int {
	weekday := time.Date(year, 0, 0, 0, 0, 0, 0, time.Local).Weekday()
	return int(weekday)
}


func GetDayOfWeek(date time.Time, weekday time.Weekday) time.Time {
	year, week := date.ISOWeek()
	thisWeekSunday := time.Date(year, 1, 1, 0, 0, 0, 0, time.Local).
		AddDate(0, 0, (week-1)*7).
		AddDate(0, 0, firstWeekDay(year))
	offsetDate := 0
	switch weekday {
	case time.Monday:
		offsetDate = -6
	case time.Tuesday:
		offsetDate = -5
	case time.Wednesday:
		offsetDate = -4
	case time.Thursday:
		offsetDate = -3
	case time.Friday:
		offsetDate = -2
	case time.Saturday:
		offsetDate = -1
	case time.Sunday:
		offsetDate = 0
	}

	return thisWeekSunday.AddDate(0, 0, offsetDate)
}

func main() {
	monday := GetDayOfWeek(time.Now(), time.Monday)
	friday := GetDayOfWeek(time.Now(), time.Friday)
	fmt.Printf("週報(%s ~ %s)\n"+
		"# やったこと\n"+
		"# 業務に関係ないやったこと\n",
		monday.Format("2006/01/02"),
		friday.Format("02"),
	)
}
