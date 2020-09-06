package periode

import (
	"errors"
	"time"
)

var (
	// ErrUnknownPeriod occurs when creates period info with unknown period.
	ErrUnknownPeriod = errors.New("unknown period")
	// ErrPassedPeriod occurs when creates period info in the past.
	ErrPassedPeriod = errors.New("passed period")
)

// Period defines period in financial terms.
type Period string

// Available period.
const (
	Monthly   = Period("MONTHLY")
	Quarterly = Period("QUARTERLY")
	Annual    = Period("ANNUAL")
)

// Info describes the period, when it's started, and when it's evaluated based on the period.
type Info struct {
	period   Period
	start    time.Time
	evaluate time.Time
}

// stub out time.Now for testing friendly
// change func now implementation in the test
// https://stackoverflow.com/a/41661462/6878109
var now = time.Now

func firstDayOfTheMonth(year int, month time.Month) time.Time {
	return time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
}

var evaluates = map[Period]func(time.Time) time.Time{
	Monthly: func(start time.Time) time.Time {
		return start.AddDate(0, 1, 0).Add(time.Nanosecond * -1)
	},
	Quarterly: func(start time.Time) time.Time {
		return start.AddDate(0, 3, 0).Add(time.Nanosecond * -1)
	},
	Annual: func(start time.Time) time.Time {
		return start.AddDate(1, 0, 0).Add(time.Nanosecond * -1)
	},
}

// NewInfo creates new period info.
func NewInfo(period Period, year int, month time.Month) (Info, error) {
	start := firstDayOfTheMonth(year, month)
	nowYear, nowMonth, _ := now().Date()
	timeNow := firstDayOfTheMonth(nowYear, nowMonth)
	if start.Before(timeNow) {
		return Info{}, ErrPassedPeriod
	}

	f, ok := evaluates[period]
	if !ok {
		return Info{}, ErrUnknownPeriod
	}
	evaluate := f(start)

	return Info{
		period:   period,
		start:    start,
		evaluate: evaluate,
	}, nil
}

// Period returns period.
func (p *Info) Period() Period {
	return p.period
}

// Start returns start time.
func (p *Info) Start() time.Time {
	return p.start
}

// Evaluate returns evaluate time.
func (p *Info) Evaluate() time.Time {
	return p.evaluate
}
