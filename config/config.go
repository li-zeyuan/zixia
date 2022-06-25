package config

import (
	"github.com/BurntSushi/toml"
)

var (
	Conf Config
)

type Config struct {
	Key             string `toml:"key"`
	TransitDataPath string `toml:"transit_data_path"`
	DrivingDataPath string `toml:"driving_data_path"`
}

func NewCfg(cfgPath string) {
	_, err := toml.DecodeFile(cfgPath, &Conf)
	if err != nil {
		panic(err)
	}
}
