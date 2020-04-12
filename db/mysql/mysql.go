package mysql

import (
	"github.com/irnak4t/leaderboards/db/config"
	"github.com/irnak4t/leaderboards/errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Open() *gorm.DB {
	cfg := config.Get()

	args := cfg.Db.User + ":" + cfg.Db.Password + "@/" + cfg.Db.Database + "?parseTime=true"
	db, err := gorm.Open("mysql", args)
	errors.FailOnError(err)
	return db
}
