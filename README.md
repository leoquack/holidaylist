# Holiday List

A very simple holiday helper

```
go get github.com/leoquack/holidaylist
```

Documentation: [godoc](https://godoc.org/github.com/leoquack/holidaylist)

### Simple usage

```go
list := holidaylist.New(time.Local)

// Add holidays.
list.Add(
	holidaylist.NewHoliday("New year's day").Date(time.January, 1),
	holidaylist.NewHoliday("Easter Sunday").Func(easterDays(0)),
	holidaylist.NewHoliday("Easter Monday").Func(easterDays(1)),
	holidaylist.NewHoliday("Christmas Day").Date(time.December, 25),
)

// Get List of holidays this year.
yearList, err := list.YearList(time.Now().Year())
if err != nil {
	// Handle error.
}
fmt.Printf("%+v\n\n", yearList)

// Check if day is holiday.
checkDay := time.Date(year, time.December, 26, 0, 0, 0, 0, location)
isHoliday, _ := yearList.IsHoliday(checkDay)
if isHoliday {
	fmt.Printf("Date '%s' is a holiday\n\n", checkDay.Format("_2 Jan 2006"))
}

// Get all holidays in date range.
from := time.Date(year, time.December, 1, 0, 0, 0, 0, location)
to := time.Date(year, time.December, 30, 0, 0, 0, 0, location)
res := yearList.FindHolidays(from, to)
for _, h := range res {
	fmt.Printf("%+v IS HOLIDAY \n", h.Time.Format("_2 Jan 2006"))
}
```

```go
// implementation of holidaylist.Calculate for easter days calculation
func easterDays(diff int) holidaylist.Calculate {
	return func(year int, location *time.Location) time.Time {
		return holidaylist.GetOrthodoxEaster(year, location).AddDate(0, 0, diff)
	}
}
```

For more examples see the [examples](https://github.com/leoquack/holidaylist/tree/master/examples) folder
