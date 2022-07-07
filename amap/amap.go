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

var (
	ErrQueryNumExceedLimit = errors.New("【宝】 😘Requests over the limit, come back tomorrow")
)

const (
	queryNumExceedLimit        = "10044"
	queryDrivingNumExceedLimit = 10003
	drivingUrl                 = "https://restapi.amap.com/v3/direction/driving?%s"
	transitUrl                 = "https://restapi.amap.com/v5/direction/transit/integrated?%s"
)

func DrivingRequest(req *model.DrivingReq) (string, error) {
	v, err := query.Values(req)
	if err != nil {
		log.Println("struct to query values error: ", err)
		return "", err
	}
	resp, err := http.Get(fmt.Sprintf(drivingUrl, v.Encode()))
	defer func() {
		if resp != nil {
			_ = resp.Body.Close()
		}
	}()
	if err != nil {
		log.Println(err)
		return "", err
	}

	if resp.StatusCode == queryDrivingNumExceedLimit {
		return "", ErrQueryNumExceedLimit
	}
	if resp.StatusCode != 200 {
		log.Println("status code: ", resp.StatusCode)
		return "", errors.New("driving request status code not 200")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("read all error: ", err)
		return "", err
	}

	drivingResp := new(model.DrivingResp)
	err = json.Unmarshal(body, drivingResp)
	if err != nil {
		log.Printf("req: %+v, json unmarshal driving response error: %s\n", req, err.Error())
		return "", err
	}

	if len(drivingResp.Route.Paths) == 0 {
		log.Printf("req: %+v", req)
		return "", errors.New("driving route path not found")
	}

	return drivingResp.Route.Paths[0].Duration, nil
}

func TransitRequest(req *model.TransitReq) (string, error) {
	req.ShowFields = "cost"
	v, err := query.Values(req)
	if err != nil {
		log.Println("struct to query values error: ", err)
		return "", err
	}
	resp, err := http.Get(fmt.Sprintf(transitUrl, v.Encode()))
	defer func() {
		if resp != nil {
			_ = resp.Body.Close()
		}
	}()
	if err != nil {
		log.Println(err)
		return "", err
	}

	if resp.StatusCode != 200 {
		log.Println("status code: ", resp.StatusCode)
		return "", errors.New("transit request status code not 200")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("read all error: ", err)
		return "", err
	}

	drivingResp := new(model.TransitResp)
	err = json.Unmarshal(body, drivingResp)
	if err != nil {
		log.Println("json unmarshal transit response error: ", err)
		return "", err
	}

	if drivingResp.Infocode == queryNumExceedLimit {
		return "", ErrQueryNumExceedLimit
	}
	if len(drivingResp.Route.Transits) == 0 {
		return "", errors.New("transit route path not found")
	}

	return drivingResp.Route.Transits[0].Cost.Duration, nil
}
