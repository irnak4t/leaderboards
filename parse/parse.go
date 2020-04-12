package parse

import (
	"strconv"
	"strings"
	"time"

	"github.com/irnak4t/leaderboards/errors"
)

func ParseTime(t string) time.Time {
	var comma string
	var nano int
	var err error
	if strings.Contains(t, ".") {
		comma = t[strings.Index(t, ".")+1:]
		for i := len(comma); i < 9; i++ {
			comma += "0"
		}
		nano, err = strconv.Atoi(comma)
		errors.FailOnError(err)
		t = t[:strings.Index(t, ".")]
	}

	if strings.Count(t, ":") < 2 {
		for i := strings.Count(t, ":"); i < 2; i++ {
			t = "00:" + t
		}
	}

	ts := strings.Split(t, ":")
	hours, err := strconv.Atoi(ts[0])
	minutes, err := strconv.Atoi(ts[1])
	seconds, err := strconv.Atoi(ts[2])

	errors.FailOnError(err)

	return time.Date(2000, time.January, 01, hours, minutes, seconds, nano, time.UTC)
}
