package maxmind

import (
	maxminddb "github.com/oschwald/maxminddb-golang"
	"net"
	"secureworks/event"
)

type LocationFinder interface {
	GetLocation(ip string) (*event.Location, error)
	Destroy()
}

type realDB struct {
	db *maxminddb.Reader
}

type maxmindRecord struct {
	Location struct {
		Latitude  float32 `maxminddb:"latitude"`
		Longitude float32 `maxminddb:"longitude"`
		Radius    int     `maxminddb:"accuracy_radius"`
	} `maxminddb:"location"`
}

func NewLocationFinder(path string) (LocationFinder, error) {
	db, err := maxminddb.Open(path)
	if err != nil {
		return nil, err
	}
	toRet := new(realDB)
	toRet.db = db
	return toRet, nil
}

func (r *realDB) Destroy() {
	if r.db != nil {
		r.db.Close()
	}
}

func (r *realDB) GetLocation(ipStr string) (*event.Location, error) {
	ip := net.ParseIP(ipStr)
	record := maxmindRecord{}
	err := r.db.Lookup(ip, &record)
	if err != nil {
		return nil, err
	}
	return event.NewLocation(
		record.Location.Latitude,
		record.Location.Longitude,
		record.Location.Radius,
	), nil
}
