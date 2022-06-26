package task

import (
	"log"

	"github.com/li-zeyuan/zixia/config"
	"github.com/li-zeyuan/zixia/excel"
	"github.com/robfig/cron/v3"
)

func New() {
	c := cron.New()

	if config.Conf.DrivingTask == "once" {
		drivingHandle()
	} else {
		c.AddFunc(config.Conf.DrivingTask, drivingHandle)
	}

	c.Start()
}

func drivingHandle() {
	log.Println("driving handle start...")
	drivingE, err := excel.NewDriving(&config.Conf)
	if err != nil {
		return
	}

	if err = drivingE.Handle(); err != nil {
		return
	}
}
