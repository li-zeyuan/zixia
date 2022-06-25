package amap

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/li-zeyuan/zixia/model"
)

const (
	drivingUrl = "https://restapi.amap.com/v3/direction/driving?%s"
)

func DrivingRequest(req *model.DrivingReq) (*model.DrivingResp, error) {
	v, err := query.Values(req)
	if err != nil {
		log.Println("struct to query values error: ", err)
		return nil, err
	}
	resp, err := http.Get(fmt.Sprintf(drivingUrl, v.Encode()))
	defer func() {
		if resp != nil {
			_ = resp.Body.Close()
		}
	}()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Println("status code: ", resp.StatusCode)
		return nil, errors.New("driving request status code not 200")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("read all error: ", err)
		return nil, err
	}

	drivingResp := new(model.DrivingResp)
	err = json.Unmarshal(body, drivingResp)
	if err != nil {
		log.Println("json unmarshal driving response error: ", err)
		return nil, err
	}

	return drivingResp, nil
}
