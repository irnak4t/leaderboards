package sort

import (
	"sort"

	"github.com/irnak4t/leaderboards/models"
)

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
