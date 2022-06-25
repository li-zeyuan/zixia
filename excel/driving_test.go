package excel

import (
	"testing"

	"github.com/li-zeyuan/zixia/config"
	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	var cfg = &config.Config{
		DrivingDataPath: "../data/21bus_test.xlsx",
	}
	e, err := New(cfg)
	assert.Nil(t, err)
	row, err := e.read()
	assert.Nil(t, err)
	t.Log(row)
}
