package sort

import (
	"sort"

	"github.com/irnak4t/leaderboards/models"
)

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
func SortByRecord(records []models.Record, sorted *map[string]map[string][]models.Record) {
	*sorted = make(map[string]map[string][]models.Record)

	for i := 0; i < len(records); i++ {
		if _, exist := (*sorted)[records[i].Title]; !exist {
			(*sorted)[records[i].Title] = make(map[string][]models.Record)
		}
		(*sorted)[records[i].Title][records[i].Category] = append((*sorted)[records[i].Title][records[i].Category], records[i])
	}

	for title, categories := range *sorted {
		for category, _ := range categories {
			sort.SliceStable(
				(*sorted)[title][category],
				func(i, j int) bool {
					ti := (*sorted)[title][category][i].Time.Unix()
					tj := (*sorted)[title][category][j].Time.Unix()
					return ti < tj
				})
		}
	}
}
