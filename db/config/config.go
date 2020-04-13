package config

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/irnak4t/leaderboards/errors"
)

type Config struct {
	MySQL MySQLConfig `toml:"mysql"`
}

type MySQLConfig struct {
	Database string `toml:"database"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

var config Config

func LoadToml() {
	cfgdir, err := os.UserConfigDir()
	_, err = toml.DecodeFile(cfgdir+"/lb.db.toml", &config)
	errors.FailOnError(err)
}

func GetMySQLConfig() MySQLConfig {
	LoadToml()
	return config.MySQL
}
