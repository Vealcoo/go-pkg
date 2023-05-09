package clock

import (
	"time"
)

// BeginningOfMonth beginning of month
func BeginningOfMonth(t time.Time) time.Time {
	y, m, _ := t.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
}

// EndOfMonth end of month
func EndOfMonth(t time.Time) time.Time {
	return BeginningOfMonth(t).AddDate(0, 1, 0).Add(-time.Millisecond)
}

// BeginningOfDay beginning of day
func BeginningOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

// EndOfDay end of day
func EndOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Microsecond), t.Location())
}

// ToISO8601Date YYYY-MM-DD
func ToISO8601Date(t time.Time) string {
	return t.Format("2006-01-02")
}

func ToMonth(t time.Time) string {
	return t.Format("2006-01")
}

// EndOfWeek end of month
func EndOfWeek(t time.Time) time.Time {
	weekDay := t.Weekday()
	if weekDay == time.Sunday {
		return t
	}
	return t.AddDate(0, 0, int(7-weekDay))
}

// BeginningOfWeek beginning of day
func BeginningOfWeek(t time.Time, weekStartDay time.Weekday) time.Time {
	t = BeginningOfDay(t)
	weekday := int(t.Weekday())
	if weekStartDay != time.Sunday {
		weekStartDayInt := int(weekStartDay)

		if weekday < weekStartDayInt {
			weekday = weekday + 7 - weekStartDayInt
		} else {
			weekday = weekday - weekStartDayInt
		}
	}
	return t.AddDate(0, 0, -weekday)
}

func ModifyMonthsWithoutOverflow(t time.Time, months int) time.Time {
	year, month, day := t.Date()
	hour, min, sec := t.Clock()

	firstDayOfMonthAfterAdd := time.Date(year, month+time.Month(months), 1, 0, 0, 0, 0, t.Location())

	lastDay := EndOfMonth(firstDayOfMonthAfterAdd).Day()
	if day > lastDay {
		day = lastDay
	}

	return time.Date(firstDayOfMonthAfterAdd.Year(), firstDayOfMonthAfterAdd.Month(), day, hour, min, sec, t.Nanosecond(), t.Location())
}
