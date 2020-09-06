// Package periode_test contains init function in the export_test.go
// to mock time.Now function. It's mock time.Now to be set
// in the first day of the month in 2020 September.
package periode_test

import (
	"testing"
	"time"

	"github.com/billyzaelani/is"
	"github.com/billyzaelani/richman/richman/periode"
)

func TestInfo(t *testing.T) {
	tests := []struct {
		period     periode.Period
		startYear  int
		startMonth time.Month

		evaluateYear  int
		evaluateMonth time.Month
	}{
		{periode.Monthly, 2020, time.September,
			2020, time.September},
		{periode.Quarterly, 2020, time.September,
			2020, time.November},
		{periode.Annual, 2020, time.September,
			2021, time.August},
	}

	for _, tt := range tests {
		t.Run(string(tt.period), func(t *testing.T) {
			is := is.New(t)
			p, err := periode.NewInfo(tt.period, tt.startYear, tt.startMonth)
			if err != nil {
				is.NoError(err)
			}

			firstDayOfTheMonth := func(year int, month time.Month) time.Time {
				return time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
			}
			lastDayOfTheMonth := func(year int, month time.Month) time.Time {
				start := firstDayOfTheMonth(year, month)
				return start.AddDate(0, 1, 0).Add(time.Nanosecond * -1)
			}

			is.Equal(p.Period(), tt.period)
			is.Equal(p.Start(), firstDayOfTheMonth(tt.startYear, tt.startMonth))
			is.Equal(p.Evaluate(), lastDayOfTheMonth(tt.evaluateYear, tt.evaluateMonth))
		})
	}
}

func TestInfoWithError(t *testing.T) {
	tests := []struct {
		period periode.Period
		year   int
		month  time.Month

		err error
	}{
		{periode.Period("UNKNOWN"), 2020, time.September,
			periode.ErrUnknownPeriod},
		{periode.Monthly, 2020, time.August,
			periode.ErrPassedPeriod},
	}

	for _, tt := range tests {
		t.Run(tt.err.Error(), func(t *testing.T) {
			is := is.New(t)
			_, err := periode.NewInfo(tt.period, tt.year, tt.month)
			is.Error(err, tt.err)
		})
	}
}
