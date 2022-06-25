package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/li-zeyuan/zixia/config"
)

func main() {
	fSet := flag.NewFlagSet("", flag.ExitOnError)
	cfgPath := fSet.String("config", "./config/config.toml", "path to config file")

	err := fSet.Parse(os.Args[1:])
	if err != nil {
		fmt.Println("flag parse error")
		os.Exit(1)
	}

	config.NewCfg(*cfgPath)

}
