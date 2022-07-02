package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/li-zeyuan/zixia/config"
	"github.com/li-zeyuan/zixia/task"
)

/*
build: GOARCH="amd64" GOOS="windows" go build -o ./bin/zixia-win-v2.exe main.go
*/
func main() {
	fSet := flag.NewFlagSet("", flag.ExitOnError)
	cfgPath := fSet.String("config", "./config.toml", "path to config file")

	err := fSet.Parse(os.Args[1:])
	if err != nil {
		fmt.Println("flag parse error")
		os.Exit(1)
	}

	config.NewCfg(*cfgPath)
	task.New()

	var sigChan = make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case <-sigChan:
			log.Println("received SIGTERM, gracefully exit...")
			return
		}
	}
}
