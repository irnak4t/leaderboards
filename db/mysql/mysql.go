package mysql

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/irnak4t/leaderboards/errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Config struct {
	Db DbConfig `toml:"mysql"`
}

type DbConfig struct {
	Database string `toml:"database"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

var config Config

func Open() *gorm.DB {
	LoadToml()

	args := config.Db.User + ":" + config.Db.Password + "@/" + config.Db.Database + "?parseTime=true"
	db, err := gorm.Open("mysql", args)
	errors.FailOnError(err)
	return db
}

func LoadToml() {
	exe, err := os.Executable()
	errors.FailOnError(err)

	dir := filepath.Dir(exe)
	if filepath.Base(exe) == "leaderboards" {
		dir = dir + "/db"
	}
	_, err = toml.DecodeFile(dir+"/db.toml", &config)
	errors.FailOnError(err)
}
