package middleware

import (
	"fmt"
	"strconv"
	"time"
)

func Inc(i int) int {
	return i + 1
}

func ParseRecord(t *time.Time) string {
	var str string
	hour, minutes, seconds := t.Clock()

	if hour != 0 {
		str = strconv.Itoa(hour) + ":"
	}

	if minutes != 0 {
		str = str + strconv.Itoa(minutes) + ":"
	}

	s := fmt.Sprintf("%02d", seconds)
	str = str + s

	if t.Nanosecond() != 0 {
		n := fmt.Sprintf("%09d", t.Nanosecond())
		str = str + "." + n[:3]
	}

	return str
}
