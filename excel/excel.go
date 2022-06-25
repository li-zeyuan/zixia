package excel

import (
	"log"

	"github.com/li-zeyuan/zixia/config"
	"github.com/xuri/excelize/v2"
)

const defaultWriteSheet = "Sheet1"

type Excel struct {
	cfg              *config.Config
	drivingReadFile  *excelize.File
	drivingWriteFile *excelize.File
}

func New(cfg *config.Config) (*Excel, error) {
	readF, err := excelize.OpenFile(cfg.DrivingDataPath)
	if err != nil {
		log.Println("open driving data excel error: ", err)
		return nil, err
	}

	writeF := excelize.NewFile()
	index := writeF.NewSheet(defaultWriteSheet)
	writeF.SetActiveSheet(index)

	return &Excel{
		cfg:              cfg,
		drivingReadFile:  readF,
		drivingWriteFile: writeF,
	}, nil
}
