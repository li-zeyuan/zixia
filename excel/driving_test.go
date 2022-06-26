package excel

import (
	"testing"

	"github.com/li-zeyuan/zixia/config"
	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	config.NewCfg("../config/config.toml")

	e, err := NewDriving(&config.Conf)
	assert.Nil(t, err)
	row, err := e.read()
	assert.Nil(t, err)
	t.Log(row)
}

func TestWrite(t *testing.T) {
	config.NewCfg("../config/config.toml")

	e, err := NewDriving(&config.Conf)
	assert.Nil(t, err)
	defer e.Close()

	err = e.write(0, []string{"name", "age"})
	assert.Nil(t, err)

}

func TestHandle(t *testing.T) {
	config.NewCfg("../config/config.toml")

	e, err := NewDriving(&config.Conf)
	assert.Nil(t, err)

	err = e.Handle()
	assert.Nil(t, err)

}
