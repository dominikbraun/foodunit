// Copyright 2019 The FoodUnit Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package controllers provides various controllers for triggering the core logic.
package controllers

import (
	"github.com/pkg/errors"
	"time"
)

// parseDates can parse strings in RFC3339 - format to time.Time note that you have
// to provide either the timezone or Z (UTC) e.g. 2015-09-15T14:00:12-00:00 or
// 2015-09-15T14:00:13Z see https://stackoverflow.com/a/32592903
//
// Returns an error if any of the dates cannot be parsed.
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
