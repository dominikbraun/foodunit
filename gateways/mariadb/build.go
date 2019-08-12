// package mariadb provides repository implementations as SQL gateways.
package mariadb

import (
	"bytes"
	"fmt"
)

type (
	// fieldMap represents a list of field names mapped against their
	// corresponding placeholder (INSERT) or data type (CREATE).
	fieldMap map[string]string
	// joinMap represents a list of JOIN types including the table name
	// mapped against an ON condition.
	joinMap map[string]string
	// conditionMap represents a list of field names mapped against their
	// corresponding operator and condition.
	conditionMap map[string]string
)

// buildInsert builds a MariaDB-compatible query string with placeholders.
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

// buildInsert builds a MariaDB-compatible query string with column types.
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

// buildSelect builds a MariaDB-compatible query string with multiple fields,
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
	}
	return buf.String()
}
