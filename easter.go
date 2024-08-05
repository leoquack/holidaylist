package holidaylist

import (
	"errors"
	"math"
	"time"
)

type easterSunday struct {
	month time.Month
	day   int
}

var easterCache = make(map[int]easterSunday, 0)

// calculateOrthodoxEaster is a slightly modified version of
// github.com/vjeantet/eastertime. Added location parameter and year cache.
func calculateOrthodoxEaster(year int, location *time.Location) (time.Time, error) {
	if year < 326 {
		return time.Time{}, errors.New("year has to be greater than 325")
	}

	if t, ok := easterCache[year]; ok {
		return time.Date(year, t.month, t.day, 0, 0, 0, 0, location), nil
	}

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

	// add to the cache
	easterCache[year] = easterSunday{month: month, day: int(day)}

	return time.Date(year, month, int(day), 0, 0, 0, 0, location), nil
}

// GetOrthodoxEaster returns Orthodox Easter Sunday
func GetOrthodoxEaster(year int, location *time.Location) time.Time {
	t, _ := calculateOrthodoxEaster(year, location)
	return t
}
