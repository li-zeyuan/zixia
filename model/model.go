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

type TransitReq struct {
	Key         string `url:"key"`
	Origin      string `url:"origin"`
	Destination string `url:"destination"`
	City1       string `url:"city1"`
	City2       string `url:"city2"`
	Date        string `url:"date"`
	Time        string `url:"time"`
	ShowFields  string `url:"show_fields"`
}

type TransitResp struct {
	Status   string           `json:"status"`
	Info     string           `json:"info"`
	Infocode string           `json:"infocode"`
	Count    string           `json:"count"`
	Route    TransitRouteResp `json:"route"`
}

type TransitRouteResp struct {
	Origin      string                 `json:"origin"`
	Destination string                 `json:"destination"`
	Distance    string                 `json:"distance"`
	TaxiCost    string                 `json:"taxi_cost"`
	Transits    []*TransitRouteTransit `json:"transits"`
}

type TransitRouteTransit struct {
	Cost            TransitRouteTransitCost `json:"cost"`
	Nightflag       string                  `json:"nightflag"`
	WalkingDistance string                  `json:"walking_distance"`
	Distance        string                  `json:"distance"`
	Missed          string                  `json:"missed"`
	Segments        []interface{}           `json:"segments"`
}

type TransitRouteTransitCost struct {
	Duration string `json:"duration"`
}
