package event

import (
	"fmt"
)

/*
	Just about all packages in this project
	will need this, just centralize it in its own
	package
*/

type Location struct {
	Latitude  float32
	Longitude float32
	Radius    int
}

type Event struct {
	IP        string
	Timestamp int //is it 2038 yet?
	*Location
}

func NewLocation(lat float32, lon float32, rad int) *Location {
	return &Location{
		Latitude:  lat,
		Longitude: lon,
		Radius:    rad,
	}
}

func (l *Location) String() string {
	return fmt.Sprintf(
		"Latitude: %f\nLongitude: %f\nRadius: %d",
		l.Latitude, l.Longitude, l.Radius,
	)
}
