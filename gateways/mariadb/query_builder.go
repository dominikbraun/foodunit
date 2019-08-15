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

// Package mariadb provides repository implementations as SQL gateways.
package mariadb

import (
	"bytes"
	"fmt"
)

type (
	// fieldMap represents a list of field names mapped against their
	// corresponding placeholder (INSERT) or data type (CREATE TABLE).
	fieldMap map[string]string
	// joinMap represents a list of JOIN types including the table name
	// mapped against an ON condition.
	joinMap map[string]string
	// conditionMap represents a list of field names mapped against their
	// corresponding operator and condition.
	conditionMap map[string]string
)

// buildInsert builds a MariaDB-compatible SQL statement with placeholders.
func buildInsert(table string, fields fieldMap) string {
	var buf bytes.Buffer

	i := 0
	buf.WriteString(fmt.Sprintf("INSERT INTO %s (", table))

	for f, _ := range fields {
		buf.WriteString(f)
		if i < len(fields)-1 {
			buf.WriteString(", ")
		}
		i++
	}

	i = 0
	buf.WriteString(fmt.Sprintf(") VALUES ("))

	for _, n := range fields {
		buf.WriteString(n)
		if i < len(fields)-1 {
			buf.WriteString(", ")
		}
		i++
	}

	buf.WriteString(")")
	return buf.String()
}

// buildInsert builds a MariaDB-compatible SQL statement with column types.
func buildCreate(table string, fields fieldMap) string {
	var buf bytes.Buffer

	i := 0
	buf.WriteString(fmt.Sprintf("CREATE TABLE %s (", table))

	for f, t := range fields {
		buf.WriteString(fmt.Sprintf("%s %s", f, t))
		if i < len(fields)-1 {
			buf.WriteString(", ")
		}
		i++
	}

	buf.WriteString(")")
	return buf.String()
}

// buildSelect builds a MariaDB-compatible SQL statement with multiple fields,
// joined tables and WHERE conditions.
func buildSelect(table string, fields []string, joins joinMap, where conditionMap) string {
	var buf bytes.Buffer

	i := 0
	buf.WriteString("SELECT ")

	for _, f := range fields {
		buf.WriteString(f)
		if i < len(fields)-1 {
			buf.WriteString(", ")
		}
		i++
	}

	i = 0
	buf.WriteString(fmt.Sprintf(" FROM %s ", table))

	for t, on := range joins {
		buf.WriteString(fmt.Sprintf("%s ON %s ", t, on))
	}

	i = 0
	buf.WriteString("WHERE ")

	for f, c := range where {
		buf.WriteString(fmt.Sprintf("%s %s", f, c))
		if i < len(fields)-1 {
			buf.WriteString(" ")
		}
		i++
	}
	return buf.String()
}

// buildUpdate builds a MariaDB-compatible SQL statement for updating rows.
func buildUpdate(table string, set fieldMap, where conditionMap) string {
	var buf bytes.Buffer

	i := 0
	buf.WriteString(fmt.Sprintf("UPDATE %s SET ", table))

	for f, v := range set {
		buf.WriteString(fmt.Sprintf("%s = %s", f, v))
		if i < len(set)-1 {
			buf.WriteString(", ")
		}
		i++
	}

	i = 0
	buf.WriteString(" WHERE ")

	for f, c := range where {
		buf.WriteString(fmt.Sprintf("%s %s", f, c))
		if i < len(where)-1 {
			buf.WriteString(" ")
		}
		i++
	}
	return buf.String()
}

// buildDelete builds a MariaDB-compatible SQL statement for deleting rows.
func buildDelete(table string, where conditionMap) string {
	var buf bytes.Buffer

	i := 0
	buf.WriteString(fmt.Sprintf("DELETE FROM %s WHERE ", table))

	for f, c := range where {
		buf.WriteString(fmt.Sprintf("%s %s", f, c))
		if i < len(where)-1 {
			buf.WriteString(" ")
		}
		i++
	}
	return buf.String()
}
