package task

import (
	"testing"

	"github.com/li-zeyuan/zixia/config"
)

func TestDrivingHandle(t *testing.T) {
	config.NewCfg("/Users/zeyuan.li/Desktop/workspace/code/src/github.com/li-zeyuan/sun/zixia/config.toml")
	drivingHandle()
}

func TestTransitHandle(t *testing.T) {
	config.NewCfg("/Users/zeyuan.li/Desktop/workspace/code/src/github.com/li-zeyuan/sun/zixia/config.toml")
	transitHandle()
}
