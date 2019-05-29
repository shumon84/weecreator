package main

import (
	"fmt"
	"time"
)

func FirstWeekDay(year int) int {
	weekday := time.Date(year, 0, 0, 0, 0, 0, 0, time.Local).Weekday()
	return int(weekday)
}

func main() {
	year, week := time.Now().ISOWeek()
	thisWeekSunday := time.Date(year, 1, 1, 0, 0, 0, 0, time.Local).
		AddDate(0, 0, (week-1)*7).
		AddDate(0, 0, FirstWeekDay(year))
	thisWeekMonday := thisWeekSunday.AddDate(0, 0, -6)
	thisWeekFriday := thisWeekSunday.AddDate(0, 0, -2)
	fmt.Printf("週報(%s ~ %s)\n"+
		"# やったこと\n"+
		"# 業務に関係ないやったこと\n",
		thisWeekMonday.Format("2006/01/02"),
		thisWeekFriday.Format("02"),
	)
}
