# Holiday List

A very simple holiday helper

## Getting Started

```
go get github.com/anheric/holidaylist
```

### Usage

```
t := time.Now()
year := t.Year()
location := t.Location()

// create a new holiday list
list := holidaylist.New(location)

// add holidays
list.Define(
	holidaylist.Holiday{
		Name: "New year's day",
		Date: time.Date(year, time.January, 1, 0, 0, 0, 0, location),
	},
	holidaylist.Holiday{
		Name: "Good Friday",
		Date: holidaylist.GetEaster(year, location).AddDate(0, 0, -2),
	},
	holidaylist.Holiday{
		Name: "Easter Sunday",
		Date: holidaylist.GetEaster(year, location),
	},
	holidaylist.Holiday{
		Name: "Christmas Day",
		Date: time.Date(year, time.December, 25, 0, 0, 0, 0, location),
	},
)

// get all holidays by year
byYear := list.GetYear(year)
fmt.Printf("By year:\n%+v\n\n", byYear)

// check if day is holiday
isHoliday, _ := list.IsHoliday(time.Date(year, time.December, 26, 0, 0, 0, 0, location))
if isHoliday {
	// your code
}

// find holidays in date range
res := list.FindHolidays(time.Date(year, time.December, 1, 0, 0, 0, 0, location), time.Date(year, time.December, 30, 0, 0, 0, 0, location))
for _, h := range res {
	fmt.Printf("%+v IS HOLIDAY \n", h.Date.Format("Mon Jan _2"))
}

```
