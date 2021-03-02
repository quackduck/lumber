package lumber

import (
	"fmt"
	"strings"
	"time"
)

// If the log should have an extra new line at the bottom
var Padding = true

// Raw formatting for lumber log
func format(stat status, t time.Time, message string) string {
	out := fmt.Sprintf("%v | %v\n%v\n",
		applyColor(stat, string(stat)),
		t.Format(time.UnixDate),
		message)

	if !Padding {
		out = strings.TrimSuffix(out, "\n")
	}

	return out
}

// Join all the items in an interface together with spaces
func separateWithSpaces(items ...interface{}) string {
	return strings.Join(items, " ")
}
