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
	errTransitCollNum = errors.New("the row is missing data")
)

type TransitExcel struct {
	cfg   *config.Config
	eFile *excelize.File
}

func NewTransit(cfg *config.Config) (*TransitExcel, error) {
	readF, err := excelize.OpenFile(cfg.Transit.DataPath)
	if err != nil {
		log.Println("open transit data excel error: ", err)
		return nil, err
	}

	return &TransitExcel{
		cfg:   cfg,
		eFile: readF,
	}, nil
}

func (t *TransitExcel) Handle() error {
	defer t.Close()
	rows, err := t.read()
	if err != nil {
		return err
	}

	for rIndex, row := range rows {
		if rIndex == 0 {
			if err = t.write(rIndex, []string{drivingTitleDuration, drivingTitleComment}); err != nil {
				return err
			}

			continue
		}

		if len(row) > defaultDrivingTitleNum {
			continue
		}

		if len(row) != defaultDrivingTitleNum {
			if err = t.write(rIndex, []string{"0", errTransitCollNum.Error()}); err != nil {
				return err
			}

			continue
		}

		TransitReq := new(model.TransitReq)
		TransitReq.Key = t.cfg.Key
		// 经度,纬度
		TransitReq.Origin = strings.Join([]string{row[4], row[3]}, ",")
		TransitReq.Destination = strings.Join([]string{row[6], row[5]}, ",")
		TransitReq.City1 = t.cfg.Transit.City
		TransitReq.City2 = t.cfg.Transit.City
		TransitReq.Date = t.cfg.Transit.Date
		TransitReq.Time = t.cfg.Transit.Time
		duration, err := amap.TransitRequest(TransitReq)
		if err != nil {
			if err == amap.ErrQueryNumExceedLimit {
				log.Println(err.Error())
				break
			}

			if err = t.write(rIndex, []string{"0", err.Error()}); err != nil {
				return err
			}
		} else {
			if err = t.write(rIndex, []string{duration, ""}); err != nil {
				return err
			}
		}

		if rIndex%100 == 0 {
			log.Printf("transit handle progress: %d/%d\n", rIndex, len(rows))
		}
	}

	log.Println("【宝】 transit handle progress: 100%")
	return nil
}

func (t *TransitExcel) read() ([][]string, error) {
	rows, err := t.eFile.GetRows(defaultWriteSheet)
	if err != nil {
		log.Printf("Error read transit excel by %s error: %s", defaultWriteSheet, err.Error())
		return nil, err
	}

	if len(rows) == 0 {
		return nil, errors.New("error transit data empty")
	}

	return rows, nil
}

func (t *TransitExcel) write(rIndex int, row []string) error {
	for i, colCell := range row {
		err := t.eFile.SetCellDefault(defaultWriteSheet, fmt.Sprintf("%s%d", string(rune(72+i)), rIndex+1), colCell)
		if err != nil {
			log.Println("set cel default error: ", err)
			return err
		}
	}

	return nil
}

func (t *TransitExcel) Close() {
	if err := t.eFile.Save(); err != nil {
		log.Println("Error close read file error: ", err)
	}
}
