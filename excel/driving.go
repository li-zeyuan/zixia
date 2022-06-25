package excel

import (
	"errors"
	"fmt"
	"log"
	"path"
	"time"

	"github.com/li-zeyuan/zixia/model"
	"github.com/li-zeyuan/zixia/amap"
)

const defaultDrivingTitleNum = 7

var excelTitleMap = map[string]string{}

func (e *Excel) Handle() error {
	defer e.Close()
	rows, err := e.read()
	if err != nil {
		return err
	}

	for rIndex, row := range rows {
		if rIndex == 0 {
			continue
		}
		if len(row) != defaultDrivingTitleNum {
			// todo write col error
			if err = e.write(row); err != nil {
				return err
			}

			continue
		}

		// todo write
		drivingReq := new(model.DrivingReq)
		amap.DrivingRequest(drivingReq)
	}

	return nil
}

func (e *Excel) read() ([][]string, error) {
	rows, err := e.drivingReadFile.GetRows(defaultWriteSheet)
	if err != nil {
		log.Printf("Error read driving excel by %s error: %s", defaultWriteSheet, err.Error())
		return nil, err
	}

	if len(rows) == 0 {
		return nil, errors.New("error driving data empty")
	}

	return rows, nil
}

func (e *Excel) writeTitle() error {
	for col, v := range excelTitleMap {
		err := e.drivingWriteFile.SetCellValue(defaultWriteSheet, col, v)
		if err != nil {
			log.Println("Error set excel title error: ", err)
			return err
		}
	}

	return nil
}

func (e *Excel) write(rows []string) error {
	if err := e.writeTitle(); err != nil {
		return err
	}

	return nil
}

func (e *Excel) Close() {
	if err := e.drivingReadFile.Close(); err != nil {
		log.Println("Error close read file error: ", err)
	}

	outPath := path.Join(path.Dir(e.cfg.DrivingDataPath), fmt.Sprintf("%s_%d",
		path.Base(e.cfg.DrivingDataPath), time.Now().Unix()))
	if err := e.drivingWriteFile.SaveAs(outPath); err != nil {
		log.Println("Error save write file error: ", err)
	}
}
