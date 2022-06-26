package excel

import (
	"errors"
	"fmt"
	"log"
	"path"
	"strings"
	"time"

	"github.com/li-zeyuan/zixia/amap"
	"github.com/li-zeyuan/zixia/config"
	"github.com/li-zeyuan/zixia/model"
	"github.com/xuri/excelize/v2"
)

const (
	defaultDrivingTitleNum = 7
	drivingTitleDuration   = "duration"
	drivingTitleComment    = "comment"
)

var (
	errDrivingCollNum = errors.New("the row is missing data")
)

type DrivingExcel struct {
	cfg          *config.Config
	reader       *excelize.File
	streamWriter *excelize.StreamWriter
	writer       *excelize.File
}

func NewDriving(cfg *config.Config) (*DrivingExcel, error) {
	readF, err := excelize.OpenFile(cfg.DrivingDataPath)
	if err != nil {
		log.Println("open driving data excel error: ", err)
		return nil, err
	}

	writeF := excelize.NewFile()
	streamWriter, err := writeF.NewStreamWriter(defaultWriteSheet)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &DrivingExcel{
		cfg:          cfg,
		reader:       readF,
		writer:       writeF,
		streamWriter: streamWriter,
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
			if err = e.write(rIndex, append(row, drivingTitleDuration, drivingTitleComment)); err != nil {
				return err
			}

			continue
		}
		if len(row) != defaultDrivingTitleNum {
			if err = e.write(rIndex, append(row, "0", errDrivingCollNum.Error())); err != nil {
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
			if err = e.write(rIndex, append(row, "0", err.Error())); err != nil {
				return err
			}
		} else {
			if err = e.write(rIndex, append(row, duration, "")); err != nil {
				return err
			}
		}

		if rIndex%100 == 0 {
			log.Printf("handle progress: %d/%d\n", rIndex, len(rows))
		}
	}

	log.Println("handle progress: done")
	return nil
}

func (e *DrivingExcel) read() ([][]string, error) {
	rows, err := e.reader.GetRows(defaultWriteSheet)
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
	rowV := make([]interface{}, 0, len(row))
	for _, colCell := range row {
		rowV = append(rowV, colCell)
	}

	if err := e.streamWriter.SetRow(fmt.Sprintf("A%d", rIndex+1), rowV); err != nil {
		log.Println("stream writer set row error: ", err)
		return err
	}

	return nil
}

func (e *DrivingExcel) Close() {
	if err := e.reader.Close(); err != nil {
		log.Println("Error close read file error: ", err)
	}

	if err := e.streamWriter.Flush(); err != nil {
		log.Println("stream writer error: ", err)
	}

	fileName := path.Base(e.cfg.DrivingDataPath)
	suffix := path.Ext(e.cfg.DrivingDataPath)
	prefix := fileName[0 : len(fileName)-len(suffix)]
	outPath := path.Join(path.Dir(e.cfg.DrivingDataPath), fmt.Sprintf("%s_%d%s",
		prefix, time.Now().Unix(), suffix))
	if err := e.writer.SaveAs(outPath); err != nil {
		log.Println("save as file error: ", err)
	}
}
