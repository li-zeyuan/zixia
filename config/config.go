package config

import (
	"github.com/BurntSushi/toml"
)

var (
	Conf Config
)

type Transit struct {
	Task     string `toml:"task"`
	DataPath string `toml:"data_path"`
	City     string `toml:"city"`
	Date     string `toml:"date"`
	Time     string `toml:"time"`
}

type Driving struct {
	Task     string `toml:"task"`
	DataPath string `toml:"data_path"`
}

type Config struct {
	Key     string  `toml:"key"`
	Driving Driving `toml:"driving"`
	Transit Transit `toml:"transit"`
}

func NewCfg(cfgPath string) {
	_, err := toml.DecodeFile(cfgPath, &Conf)
	if err != nil {
		panic(err)
	}
}
