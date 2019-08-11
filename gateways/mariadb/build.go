// package mariadb provides repository implementations as SQL gateways.
package mariadb

import (
	"bytes"
	"fmt"
)

// namedFields represents multiple field names (keys) mapped against their
// corresponding placeholder (values).
type namedFields map[string]string

// buildInsert builds a MariaDB-compatible query string with placeholders.
func buildInsert(table string, fields namedFields) string {
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

func buildCreate(table string, fields namedFields) string {
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
