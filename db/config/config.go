package config

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/irnak4t/leaderboards/errors"
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

func LoadToml() {
	exe, _ := os.Executable()

	dir := filepath.Dir(exe)
	if filepath.Base(exe) == "leaderboards" {
		dir = dir + "/db"
	}
	_, err := toml.DecodeFile(dir+"/db.toml", &config)
	errors.FailOnError(err)
}

func Get() Config {
	LoadToml()
	return config
}
