package parse

import (
	"reflect"
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	post := "1:30:15"
	parsed := ParseTime(post)
	exp := time.Date(2000, time.January, 01, 1, 30, 15, 0, time.UTC)
	if !reflect.DeepEqual(parsed, exp) {
		t.Fatalf("Unexpected datetime.\n parsed: %#v\n exp: %#v\n", parsed, exp)
	}

	post = "30:15.382"
	parsed = ParseTime(post)
	exp = time.Date(2000, time.January, 01, 0, 30, 15, 382000000, time.UTC)
	if !reflect.DeepEqual(parsed, exp) {
		t.Fatalf("Unexpected datetime.\n parsed: %#v\n exp: %#v\n", parsed, exp)
	}

	post = "30:15.012"
	parsed = ParseTime(post)
	exp = time.Date(2000, time.January, 01, 0, 30, 15, 12000000, time.UTC)
	if !reflect.DeepEqual(parsed, exp) {
		t.Fatalf("Unexpected datetime.\n parsed: %#v\n exp: %#v\n", parsed, exp)
	}
}
