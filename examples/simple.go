package main

import (
	"fmt"
	"time"

	"github.com/leoquack/holidaylist"
)

func main() {
	location := time.Local

	list := holidaylist.New(location)

	list.Add(
		holidaylist.NewHoliday("New year's day").Date(time.January, 1),
		holidaylist.NewHoliday("Epiphany").Date(time.January, 6),
		holidaylist.NewHoliday("Assumption Day").Date(time.August, 15),
		holidaylist.NewHoliday("Christmas Day").Date(time.December, 25),
		holidaylist.NewHoliday("Boxing Day").Date(time.December, 26),
	)

	addOrthodoxEasterHolidays(list)

	year := time.Now().Year()

	// Get all holidays of a year.
	yearList, err := list.YearList(year)
	if err != nil {
		fmt.Println("couldn't get holidays by year:", err)
	}
	fmt.Printf("By year:\n%+v\n\n", yearList.Holidays)

	// Get holidays by date range.
	from := time.Date(year, time.December, 1, 0, 0, 0, 0, location)
	to := time.Date(year, time.December, 26, 0, 0, 0, 0, location)
	rangeList, err := list.RangeList(from, to)
	if err != nil {
		fmt.Println("couldn't get holidays by date range:", err)
	}
	fmt.Printf("By date range:\n%+v\n\n", rangeList)

	checkDay := time.Date(year, time.December, 26, 0, 0, 0, 0, location)
	isHoliday, _ := yearList.IsHoliday(checkDay)
	if isHoliday {
		fmt.Printf("Date '%s' is a holiday\n\n", checkDay.Format("_2 Jan 2006"))
	}

	// Find holidays in date range.
	res := yearList.FindHolidays(time.Date(year, time.December, 1, 0, 0, 0, 0, location), time.Date(year, time.December, 30, 0, 0, 0, 0, location))
	for _, h := range res {
		fmt.Printf("%+v IS HOLIDAY \n", h.Time.Format("_2 Jan 2006"))
	}
}

// Orthodox Easter example.
const (
	EasterDiffGreenMonday  = -48
	EasterDiffGoodFriday   = -2
	EasterDiffHolySaturday = -1
	EasterDiffSunday       = 0
	EasterDiffMonday       = 1
	EasterDiffWhitMonday   = 50
)

func easterDays(diff int) holidaylist.Calculate {
	return func(year int, location *time.Location) time.Time {
		return holidaylist.GetOrthodoxEaster(year, location).AddDate(0, 0, diff)
	}
}

func addOrthodoxEasterHolidays(list *holidaylist.Definitions) {
	list.Add(
		holidaylist.NewHoliday("Green Monday").Func(easterDays(EasterDiffGreenMonday)),
		holidaylist.NewHoliday("Good Friday").Func(easterDays(EasterDiffGoodFriday)),
		holidaylist.NewHoliday("Holy Saturday").Func(easterDays(EasterDiffHolySaturday)),
		holidaylist.NewHoliday("Easter Sunday").Func(easterDays(EasterDiffSunday)),
		holidaylist.NewHoliday("Easter Monday").Func(easterDays(EasterDiffMonday)),
		holidaylist.NewHoliday("Whit Monday").Func(easterDays(EasterDiffWhitMonday)),
	)
}
