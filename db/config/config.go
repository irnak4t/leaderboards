package config

import (
	"os"

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
	home, err := os.UserConfigDir()
	_, err = toml.DecodeFile(home+"/lb.db.toml", &config)
	errors.FailOnError(err)
}

func Get() Config {
	LoadToml()
	return config
}
