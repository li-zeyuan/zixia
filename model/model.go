package model

type DrivingReq struct {
	Key         string `url:"key"`
	Origin      string `url:"origin"`
	Destination string `url:"destination"`
	Strategy    int    `url:"strategy"`
}

type DrivingResp struct {
	Status   string           `json:"status"`
	Info     string           `json:"info"`
	Infocode string           `json:"infocode"`
	Count    string           `json:"count"`
	Route    DrivingRouteResp `json:"route"`
}

type DrivingRouteResp struct {
	Origin      string                  `json:"origin"`
	Destination string                  `json:"destination"`
	Paths       []*DrivingRoutePathResp `json:"paths"`
}

type DrivingRoutePathResp struct {
	Distance      string        `json:"distance"`
	Duration      string        `json:"duration"`
	Strategy      string        `json:"strategy"`
	Tolls         string        `json:"tolls"`
	TollDistance  string        `json:"toll_distance"`
	Steps         []interface{} `json:"steps"`
	Restriction   string        `json:"restriction"`
	TrafficLights string        `json:"traffic_lights"`
}
