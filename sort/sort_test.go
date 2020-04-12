package sort

import (
	"reflect"
	"testing"
	"time"

	"github.com/irnak4t/leaderboards/models"
)

var sorted map[string]map[string][]models.Record

/*
	Sort out records slice
	[
		0: Record{
			Title: "title1"
			Category: "category%"
			...
		}
		1: Record{
			Title: "title2"
			Category: "category%"
			...
		}
	]

	to

	[
		"title1": [
			"category1": [
				0: Record{}
				1: Record{}
			]
			"category2": [
				0: Record{}
				1: Record{}
			]
		]
		"title2": [
			"category1": [
				0: Record{}
				1: Record{}
			]
			"category2": [
				0: Record{}
				1: Record{}
			]
		]
	]
*/
func TestSortByRecord(t *testing.T) {
	record1 := models.Record{
		ID:       1,
		Title:    "Title1",
		Category: "Any%",
		Runner:   "runner1",
		Time:     time.Date(2000, time.January, 01, 00, 31, 10, 123999999, time.UTC),
		Url:      "http://example.com",
	}
	record2 := models.Record{
		ID:       2,
		Title:    "Title1",
		Category: "100%",
		Runner:   "runner2",
		Time:     time.Date(2000, time.January, 01, 00, 31, 10, 123999999, time.UTC),
		Url:      "http://example.com2",
	}
	record3 := models.Record{
		ID:       3,
		Title:    "Title2",
		Category: "Any%",
		Runner:   "runner1",
		Time:     time.Date(2000, time.January, 01, 00, 31, 11, 123999999, time.UTC),
		Url:      "http://example.com3",
	}
	record4 := models.Record{
		ID:       4,
		Title:    "Title2",
		Category: "Any%",
		Runner:   "runner1",
		Time:     time.Date(2000, time.January, 01, 00, 31, 4, 123999999, time.UTC),
		Url:      "http://example.com4",
	}
	records := []models.Record{
		record1,
		record2,
		record3,
		record4,
	}
	expect := make(map[string]map[string][]models.Record)

	expect["Title1"] = map[string][]models.Record{
		"Any%": []models.Record{
			record1,
		},
		"100%": []models.Record{
			record2,
		},
	}
	expect["Title2"] = map[string][]models.Record{
		"Any%": []models.Record{
			record4,
			record3,
		},
	}

	SortByRecord(records, &sorted)

	if !reflect.DeepEqual(sorted, expect) {
		t.Fatalf("Unexpected sorting.\n sorted: %#v\n expect: %#v\n", sorted, expect)
	}
}
