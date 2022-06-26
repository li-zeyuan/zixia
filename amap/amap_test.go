package amap

import (
	"strings"
	"testing"

	"github.com/li-zeyuan/zixia/config"
	"github.com/li-zeyuan/zixia/model"
	"github.com/stretchr/testify/assert"
)

func TestDrivingRequest(t *testing.T) {
	config.NewCfg("../config/config.toml")

	drivingReq := new(model.DrivingReq)
	drivingReq.Key = config.Conf.Key
	// 经度,纬度
	drivingReq.Origin = strings.Join([]string{"113.902", "22.477301"}, ",")
	drivingReq.Destination = strings.Join([]string{"114.030998", "22.5394"}, ",")
	drivingReq.Strategy = 0

	duration, err := DrivingRequest(drivingReq)
	assert.Nil(t, err)
	t.Log(duration)
}
