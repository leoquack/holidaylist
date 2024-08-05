package holidaylist

import (
	"errors"
	"time"
)

// Holiday holds all information about the day of the holiday.
type Holiday struct {
	Name  string     `json:"name"`
	Year  int        `json:"year"`
	Month time.Month `json:"month"`
	Day   int        `json:"day"`
	Time  time.Time  `json:"date"`

	Calc Calculate
}

// Calculate helps calculate a variable holiday.
type Calculate func(year int, location *time.Location) time.Time

// Definitions hold the holiday definitions.
type Definitions struct {
	Location *time.Location
	Holidays []Holiday
}

// List used for further holiday checkups.
type List Definitions

// New returns an empty holiday definitions.
func New(location *time.Location) *Definitions {
	return &Definitions{
		Location: location,
	}
}

// NewHoliday creates and returns a new Holiday.
func NewHoliday(name string) *Holiday {
	return &Holiday{
		Name: name,
	}
}

// Add appends into Holiday definitions.
func (d *Definitions) Add(h ...*Holiday) {
	for _, p := range h {
		d.Holidays = append(d.Holidays, *p)
	}
}

// YearList returns List with holidays of requested year.
func (d *Definitions) YearList(year int) (*List, error) {
	if year < 326 {
		return nil, errors.New("year has to be greater than 325")
	}
	yl := &List{
		Location: d.Location,
	}
	for _, h := range d.Holidays {
		h.calcTime(year, yl.Location)
		yl.Holidays = append(yl.Holidays, h)
	}
	return yl, nil
}

// RangeList returns List with holidays of requested date range.
func (d *Definitions) RangeList(from, to time.Time) (*List, error) {
	if from.After(to) {
		return nil, errors.New(`"from date" should be an earlier date than "to date"`)
	}

	fromYear := from.Year()
	toYear := to.Year()
	yl := &List{
		Location: d.Location,
		Holidays: make([]Holiday, 0),
	}

	for i, r := 0, toYear-fromYear+1; i < r; i++ {
		for _, h := range d.Holidays {
			h.calcTime(fromYear+i, yl.Location)
			if (h.Time.After(from) && h.Time.Before(to)) || h.Time.Equal(from) || h.Time.Equal(to) {
				yl.Holidays = append(yl.Holidays, h)
			}
		}
	}
	return yl, nil
}

// IsHoliday checks if date is a holiday and returns it.
func (l *List) IsHoliday(t time.Time) (bool, *Holiday) {
	for _, h := range l.Holidays {
		if h.Time.Year() == t.Year() && h.Time.YearDay() == t.YearDay() {
			return true, &h
		}
	}
	return false, nil
}

// FindHolidays returns holiday days from date range.
func (l *List) FindHolidays(from, to time.Time) []Holiday {
	holidays := make([]Holiday, 0)
	for _, h := range l.Holidays {
		if h.Time.After(from) && h.Time.Before(to) {
			holidays = append(holidays, h)
		}
	}
	return holidays
}

// Date sets month and day to Holiday.
func (h *Holiday) Date(month time.Month, day int) *Holiday {
	h.Month = month
	h.Day = day
	return h
}

// Func sets the Calc function.
func (h *Holiday) Func(calc Calculate) *Holiday {
	h.Calc = calc
	return h
}

// calcTime uses either Calc function, if defined, or year, month and day
// to calculate the date of the holiday for requested year.
func (h *Holiday) calcTime(year int, location *time.Location) {
	if h.Calc != nil {
		h.Time = h.Calc(year, location)
	} else {
		h.Time = time.Date(year, h.Month, h.Day, 0, 0, 0, 0, location)
	}
}
