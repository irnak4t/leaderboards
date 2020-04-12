package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	mysql "github.com/irnak4t/leaderboards/db/mysql"
	"github.com/irnak4t/leaderboards/models"
)

func main() {
	db := mysql.Open()
	defer db.Close()

	if !db.HasTable("records") {
		db.CreateTable(&models.Record{})
		fmt.Printf("\x1b[32m%s\x1b[0m %s\n", "[success]", "Creating table successfully")
	} else {
		fmt.Printf("\x1b[36m%s\x1b[0m %s\n", "[skipped]", "Table already exists.")
	}
}
