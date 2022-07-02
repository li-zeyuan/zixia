package excel

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/li-zeyuan/zixia/amap"
	"github.com/li-zeyuan/zixia/config"
	"github.com/li-zeyuan/zixia/model"
	"github.com/xuri/excelize/v2"
)

var (
	errDrivingCollNum = errors.New("the row is missing data")
)

type DrivingExcel struct {
	cfg   *config.Config
	eFile *excelize.File
}

func NewDriving(cfg *config.Config) (*DrivingExcel, error) {
	readF, err := excelize.OpenFile(cfg.Driving.DataPath)
	if err != nil {
		log.Println("open driving data excel error: ", err)
		return nil, err
	}

	return &DrivingExcel{
		cfg:   cfg,
		eFile: readF,
	}, nil
}

func (e *DrivingExcel) Handle() error {
	defer e.Close()
	rows, err := e.read()
	if err != nil {
		return err
	}

	for rIndex, row := range rows {
		if rIndex == 0 {
			if err = e.write(rIndex, []string{drivingTitleDuration, drivingTitleComment}); err != nil {
				return err
			}

			continue
		}

		if len(row) > defaultDrivingTitleNum {
			continue
		}

		if len(row) != defaultDrivingTitleNum {
			if err = e.write(rIndex, []string{"0", errDrivingCollNum.Error()}); err != nil {
				return err
			}

			continue
		}

		drivingReq := new(model.DrivingReq)
		drivingReq.Key = e.cfg.Key
		// 经度,纬度
		drivingReq.Origin = strings.Join([]string{row[4], row[3]}, ",")
		drivingReq.Destination = strings.Join([]string{row[6], row[5]}, ",")
		drivingReq.Strategy = 0
		duration, err := amap.DrivingRequest(drivingReq)
		if err != nil {
			if err == amap.ErrQueryNumExceedLimit {
				log.Println(err.Error())
				break
			}

			if err = e.write(rIndex, []string{"0", err.Error()}); err != nil {
				return err
			}
		} else {
			if err = e.write(rIndex, []string{duration, ""}); err != nil {
				return err
			}
		}

		if rIndex%100 == 0 {
			log.Printf("driving handle progress: %d/%d\n", rIndex, len(rows))
		}
	}

	log.Println("宝 driving handle progress: done")
	return nil
}

func (e *DrivingExcel) read() ([][]string, error) {
	rows, err := e.eFile.GetRows(defaultWriteSheet)
	if err != nil {
		log.Printf("Error read driving excel by %s error: %s", defaultWriteSheet, err.Error())
		return nil, err
	}

	if len(rows) == 0 {
		return nil, errors.New("error driving data empty")
	}

	return rows, nil
}

func (e *DrivingExcel) write(rIndex int, row []string) error {
	for i, colCell := range row {
		err := e.eFile.SetCellDefault(defaultWriteSheet, fmt.Sprintf("%s%d", string(rune(72+i)), rIndex+1), colCell)
		if err != nil {
			log.Println("set cel default error: ", err)
			return err
		}
	}

	return nil
}

func (e *DrivingExcel) Close() {
	if err := e.eFile.Save(); err != nil {
		log.Println("Error close read file error: ", err)
	}
}
