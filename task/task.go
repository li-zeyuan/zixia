package task

import (
	"log"

	"github.com/li-zeyuan/zixia/config"
	"github.com/li-zeyuan/zixia/excel"
	"github.com/robfig/cron/v3"
)

func New() {
	img2txt("./.zixia.png",
		150, []string{"*", "%", "+", ",", ".", " "}, "\n", "./zixia")

	c := cron.New()

	if config.Conf.Driving.Task == "once" {
		drivingHandle()
	} else if len(config.Conf.Driving.Task) > 0 {
		c.AddFunc(config.Conf.Driving.Task, drivingHandle)
	}

	if config.Conf.Transit.Task == "once" {
		transitHandle()
	} else if len(config.Conf.Transit.Task) > 0 {
		c.AddFunc(config.Conf.Transit.Task, transitHandle)
	}

	c.Start()
}

func drivingHandle() {
	log.Println("driving handle start...")
	driving, err := excel.NewDriving(&config.Conf)
	if err != nil {
		return
	}

	if err = driving.Handle(); err != nil {
		return
	}
}

func transitHandle() {
	log.Println("transit handle start...")
	transit, err := excel.NewTransit(&config.Conf)
	if err != nil {
		return
	}

	if err = transit.Handle(); err != nil {
		return
	}
}
