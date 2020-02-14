package main

type RequestJSON struct {
	UserName      string `json:"username"`
	UnixTimestamp uint   `json:"unix_timestamp"`
	EventUUID     string `json:"event_uuid"`
	IPAddress     string `json:"ip_address"`
}

type CurrGeo struct {
	Lat    float64 `json:"lat"`
	Lon    float64 `json:"lon"`
	Radius uint    `json:"radius"`
}

type IpAccess struct {
	IP     string  `json:"ip"`
	Speed  uint    `json:"speed"`
	Lat    float64 `json:"lat"`
	Lon    float64 `json:"lon"`
	Radius uint    `json:"radius"`
	Time   uint    `json:"timestamp"`
}

type ResultJSON struct {
	CurrGeo        CurrGeo  `json:"currentGeo"`
	ToSuspicious   bool     `json:"travelToCurrentGeoSuspicious,omitempty"`
	FromSuspicious bool     `json:"travelFromCurrentGeoSuspicious,omitempty"`
	Preceding      IpAccess `json:"precedingIpAccess,omitempty"`
	Subsequent     IpAccess `json:"subsequentIpAccess,omitempty"`
}
