package models

import (
	"time"

	"github.com/irnak4t/leaderboards/db/mysql"
)

type Record struct {
	ID       int    `gorm:primary_key`
	Title    string `gorm:"not null"`
	Category string `gorm:"not null"`
	Runner   string
	Time     time.Time `gorm:"type:datetime(3)"`
	Url      string    `gorm:"unique"`
}

func (r Record) GetAll() []Record {
	db := mysql.Open()
	defer db.Close()

	var records []Record
	db.Find(&records)
	return records
}

func (r Record) Create() {
	db := mysql.Open()
	defer db.Close()
	db.Create(&r)
}

func (r Record) GetRuns(runner string) []Record {
	db := mysql.Open()
	defer db.Close()

	var records []Record
	db.Where("runner = ?", runner).Find(&records)
	return records
}

func (r Record) Get(id int) Record {
	db := mysql.Open()
	defer db.Close()

	var record Record
	db.Where("id = ?", id).First(&record)
	return record
}

func (r Record) Delete() {
	db := mysql.Open()
	defer db.Close()
	db.Delete(&r)
}
