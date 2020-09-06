package periode

import "time"

var (
	Evaluates = evaluates
	Now       func() time.Time
)

func init() {
	now = func() time.Time {
		// NOTE: for testing only, now is set to be first day of the month in 2020 September
		return firstDayOfTheMonth(2020, time.September)
	}
	Now = now
}
