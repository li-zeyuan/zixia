package excel

import (
	"testing"

	"github.com/li-zeyuan/zixia/config"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	var cfg = &config.Config{
		DrivingDataPath: "../data/21bus.xlsx",
	}

	e, err := New(cfg)
	assert.Nil(t, err)
	t.Log(e)
}
