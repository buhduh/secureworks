package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"secureworks/constants"
	"secureworks/database"
	"secureworks/event"
	"secureworks/maxmind"
	"secureworks/secureworksmath"
)

var db database.Database
var locFinder maxmind.LocationFinder

func init() {
	var err error
	db, err = database.NewDatabase()
	if err != nil {
		log.Fatalf("failed initializing db with error, '%s'", err)
	}
	locFinder, err = maxmind.NewLocationFinder(constants.IP_DB)
	if err != nil {
		log.Fatalf("failed initializing maxmind database with error, '%s'", err)
	}
}

func doWork(req *RequestJSON) (*ResultJSON, error) {
	incEvent := event.NewRawEvent(
		req.IPAddress, req.UserName, req.EventUUID, req.UnixTimestamp,
	)
	location, err := locFinder.GetLocation(req.IPAddress)
	if err != nil {
		return nil, err
	}
	incEvent.Location = location
	err = db.NewEvent(incEvent)
	if err != nil {
		return nil, err
	}
	events, err := db.GetOrderedEventsForUser(req.UserName)
	if err != nil {
		return nil, err
	}
	found, err := incEvent.GetSurroundingEvents(events)
	if err != nil {
		return nil, err
	}
	toRet := ResultJSON{}
	currGeo := CurrGeo{
		Lat:    found[1].Location.Latitude,
		Lon:    found[1].Location.Longitude,
		Radius: found[1].Location.Radius,
	}
	toRet.CurrGeo = currGeo
	if found[0] != nil {
		speed := secureworksmath.Speed(found[0], found[1])
		precIP := IpAccess{
			IP:     found[0].IP,
			Speed:  uint(speed),
			Lat:    found[0].Location.Latitude,
			Lon:    found[0].Location.Longitude,
			Radius: found[0].Location.Radius,
			Time:   found[0].Timestamp,
		}
		isSuperMan := secureworksmath.IsSuperMan(speed)
		toRet.FromSuspicious = isSuperMan
		toRet.Preceding = precIP
	}
	if found[2] != nil {
		speed := secureworksmath.Speed(found[1], found[2])
		subsIP := IpAccess{
			IP:     found[2].IP,
			Speed:  uint(speed),
			Lat:    found[2].Location.Latitude,
			Lon:    found[2].Location.Longitude,
			Radius: found[2].Location.Radius,
			Time:   found[2].Timestamp,
		}
		isSuperMan := secureworksmath.IsSuperMan(speed)
		toRet.FromSuspicious = isSuperMan
		toRet.Subsequent = subsIP
	}
	return &toRet, nil
}

func handler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(
			w, fmt.Sprintf("%s not supported", req.Method),
			http.StatusNotImplemented,
		)
		return
	}
	body, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, "could not process body", http.StatusInternalServerError)
		return
	}
	rJSON := RequestJSON{}
	err = json.Unmarshal(body, &rJSON)
	if err != nil {
		http.Error(w, "could not process body", http.StatusInternalServerError)
		return
	}
	resJSON, err := doWork(&rJSON)
	if err != nil {
		msg := fmt.Sprintf("could not process request with error '%s'", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	data, err := json.MarshalIndent(*resJSON, "", "  ")
	if err != nil {
		msg := fmt.Sprintf("could not process request with error '%s'", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(data))
}

//go:generate /root/generate.sh
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%s", constants.PORT), nil)
}
