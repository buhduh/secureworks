package event

import (
	"fmt"
	"sort"
)

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

//Get SurroundingEvents is MUST be in order
//an error if e is not in events
//0th return is preceeding event
//1st return is event
//2nd return is proceeding
//preceeding and proceeding can be nil
//timestamps can't be identical for events
func (e *Event) GetSurroundingEvents(events []*Event) ([3]*Event, error) {
	toRet := [3]*Event{}
	if len(events) == 0 {
		return toRet, fmt.Errorf("passed len(events) == 0")
	}
	index := sort.Search(
		len(events),
		func(i int) bool {
			return events[i].Timestamp >= e.Timestamp
		},
	)
	if index == len(events) {
		return toRet,
			fmt.Errorf(
				"could not locate event with timestamp %d",
				e.Timestamp,
			)
	}
	if e.EventUUID != events[index].EventUUID {
		return toRet, fmt.Errorf("event uuids do not match!!!")
	}
	if index > 0 {
		toRet[0] = events[index-1]
	}
	toRet[1] = events[index]
	if index+1 < len(events) {
		toRet[2] = events[index+1]
	}
	return toRet, nil
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
