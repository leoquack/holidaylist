package main

import (
	"fmt"
	"time"

	"github.com/leoquack/holidaylist"
)

func easterDays(diff int) holidaylist.Calculate {
	return func(year int, location *time.Location) time.Time {
		return holidaylist.GetOrthodoxEaster(year, location).AddDate(0, 0, diff)
	}
}

func main() {
	location := time.Local

	list := holidaylist.New(location)
	// add holidays
	list.Add(
		holidaylist.NewHoliday("New year's day").Date(time.January, 1),
		holidaylist.NewHoliday("Epiphany").Date(time.January, 6),
		holidaylist.NewHoliday("Green Monday").Func(easterDays(-48)),
		holidaylist.NewHoliday("Good Friday").Func(easterDays(-2)),
		holidaylist.NewHoliday("Holy Saturday").Func(easterDays(-1)),
		holidaylist.NewHoliday("Easter Sunday").Func(easterDays(0)),
		holidaylist.NewHoliday("Easter Monday").Func(easterDays(1)),
		holidaylist.NewHoliday("Whit Monday").Func(easterDays(50)),
		holidaylist.NewHoliday("Assumption Day").Date(time.August, 15),
		holidaylist.NewHoliday("Christmas Day").Date(time.December, 25),
		holidaylist.NewHoliday("Boxing Day").Date(time.December, 26),
	)

	year := time.Now().Year()

	// get List of holidays by year
	yearList, err := list.YearList(year)
	if err != nil {
		fmt.Println("error getting holidays by year:", err)
	}
	fmt.Printf("By year:\n%+v\n\n", yearList.Holidays)

	// get List of holidays by date range
	from := time.Date(year, time.December, 1, 0, 0, 0, 0, location)
	to := time.Date(year, time.December, 26, 0, 0, 0, 0, location)
	rangeList, err := list.RangeList(from, to)
	if err != nil {
		fmt.Println("error getting holidays by date range:", err)
	}
	fmt.Printf("By date range:\n%+v\n\n", rangeList)

	// check if day is holiday
	isHoliday, _ := yearList.IsHoliday(time.Date(year, time.December, 26, 0, 0, 0, 0, location))
	if isHoliday {
		// your code
	}

	// find holidays in date range
	res := yearList.FindHolidays(time.Date(year, time.December, 1, 0, 0, 0, 0, location), time.Date(year, time.December, 30, 0, 0, 0, 0, location))
	for _, h := range res {
		fmt.Printf("%+v IS HOLIDAY \n", h.Time.Format("Mon Jan _2"))
	}
}
