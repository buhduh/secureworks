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
	Latitude  float64
	Longitude float64
	Radius    uint
}

type Event struct {
	IP        string
	Timestamp uint //is it 2038 yet?
	UserName  string
	EventUUID string
	Location  *Location
}

func NewLocation(lat float64, lon float64, rad uint) *Location {
	return &Location{
		Latitude:  lat,
		Longitude: lon,
		Radius:    rad,
	}
}

//TODO, I should probably validate this
func NewEvent(
	ip, uName, eUUID string,
	time, rad uint,
	lat, lon float64,
) (*Event, error) {
	return &Event{
		IP:        ip,
		Timestamp: time,
		UserName:  uName,
		EventUUID: eUUID,
		Location: &Location{
			Latitude:  lat,
			Longitude: lon,
			Radius:    rad,
		},
	}, nil
}

//This will most likely be incomming events
func NewRawEvent(ip, uName, eUUID string, time uint) *Event {
	return &Event{
		IP:        ip,
		Timestamp: time,
		UserName:  uName,
		EventUUID: eUUID,
	}
}

func (l *Location) String() string {
	return fmt.Sprintf(`
		Latitude:  %f
    Longitude: %f
    Radius:    %d
`,
		l.Latitude,
		l.Longitude,
		l.Radius,
	)
}

func (e *Event) String() string {
	eStr := fmt.Sprintf(`
  IP:        '%s'
  Timestamp: %d
  UserName:  '%s'
  EventUUID: '%s'
`,
		e.IP, e.Timestamp, e.UserName, e.EventUUID,
	)
	if e.Location == nil {
		return eStr
	}
	return fmt.Sprintf("%s%s", eStr, e.Location)
}
