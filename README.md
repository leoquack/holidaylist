# Holiday List

A very simple holiday helper

```
go get github.com/anheric/holidaylist
```

Documentation: [godoc](https://godoc.org/github.com/anheric/holidaylist)

### Simple usage

```go
list := holidaylist.New(time.Local)
// add holidays
list.Add(
	list.New("New year's day").Date(time.January, 1),
	list.New("Easter Sunday").Func(easterDays(0)),
	list.New("Easter Monday").Func(easterDays(1)),
	list.New("Christmas Day").Date(time.December, 25),
)
// get List of holidays this year
yearList, err := list.YearList(time.Now().Year())
if err != nil {
	// handle error
}
fmt.Printf("%+v\n\n", yearList)

// check if day is holiday
isHoliday, err := yearList.IsHoliday(time.Date(year, time.December, 26, 0, 0, 0, 0, location))
if err != nil {
	// handle error
}
if isHoliday {
	// your code
}

// find holidays in date range
res := yearList.FindHolidays(time.Date(year, time.December, 1, 0, 0, 0, 0, location), time.Date(year, time.December, 30, 0, 0, 0, 0, location))
for _, h := range res {
	fmt.Printf("%+v IS HOLIDAY \n", h.Time.Format("Mon Jan _2"))
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

For more examples see the [examples](https://github.com/anheric/holidaylist/tree/master/examples) folder
