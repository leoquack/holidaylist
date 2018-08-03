package holidaylist

import (
	"math"
	"time"
)

// Holiday holds all information about the day of the holiday
type Holiday struct {
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

// List holds the list of holidays
type List struct {
	Location *time.Location
	Holidays []Holiday
}

// New returns an empty Holidays list
func New(location *time.Location) *List {
	return &List{
		Location: location,
	}
}

// Define inserts holidays into the list
func (l *List) Define(h ...Holiday) {
	for _, p := range h {
		l.Holidays = append(l.Holidays, p)
	}
}

// GetYear returns holidays of year
func (l *List) GetYear(year int) []Holiday {
	yearList := make([]Holiday, 0)
	for _, h := range l.Holidays {
		if h.Date.Year() == year {
			yearList = append(yearList, h)
		}
	}
	return yearList
}

// IsHoliday checks if date is a holiday and returns it
func (l *List) IsHoliday(t time.Time) (bool, *Holiday) {
	for _, h := range l.Holidays {
		if h.Date.Year() == t.Year() && h.Date.YearDay() == t.YearDay() {
			return true, &h
		}
	}
	return false, nil
}

// FindHolidays returns holiday days from date range
func (l *List) FindHolidays(from, to time.Time) []Holiday {
	holidays := make([]Holiday, 0)
	for _, h := range l.Holidays {
		if h.Date.After(from) && h.Date.Before(to) {
			holidays = append(holidays, h)
		}
	}
	return holidays
}

// GetOrthodoxEaster returns Orthodox Easter Sunday
// slightly modified version of "github.com/vjeantet/eastertime"
func GetOrthodoxEaster(year int, location *time.Location) time.Time {
	var a, b, c, d, e int
	var month time.Month
	var day float64

	a = year % 4
	b = year % 7
	c = year % 19
	d = (19*c + 15) % 30
	e = (2*a + 4*b - d + 34) % 7
	month = time.Month((d + e + 114) / 31)
	day = math.Floor(float64((d+e+114)%31 + 1))
	day = day + 13

	return time.Date(year, month, int(day), 0, 0, 0, 0, location)
}
