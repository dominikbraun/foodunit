package controllers

import (
	"github.com/pkg/errors"
	"time"
)

// can parse strings in RFC3339 - format to time.Time
// note that you have to provide either the timezone or Z (UTC)
// e.g. 2015-09-15T14:00:12-00:00 or 2015-09-15T14:00:13Z
// see https://stackoverflow.com/a/32592903
//
// error if any of the dates cannot be parsed
func parseDates(dates ...string) ([]time.Time, error) {
	result := make([]time.Time, 0)

	for i := range dates {
		date, err := time.Parse(time.RFC3339, dates[i])
		if err != nil {
			return nil, errors.Wrapf(err, "while parsing %v", dates[i])
		}
		result = append(result, date)
	}

	return result, nil
}
