package event

import (
	"testing"
)

var events []*Event = []*Event{
	&Event{
		Timestamp: 0,
		EventUUID: "0",
	},
	&Event{
		Timestamp: 1,
		EventUUID: "1",
	},
	&Event{
		Timestamp: 2,
		EventUUID: "2",
	},
	&Event{
		Timestamp: 3,
		EventUUID: "3",
	},
	&Event{
		Timestamp: 4,
		EventUUID: "4",
	},
	&Event{
		Timestamp: 5,
		EventUUID: "5",
	},
	&Event{
		Timestamp: 6,
		EventUUID: "6",
	},
	&Event{
		Timestamp: 7,
		EventUUID: "7",
	},
}

func TestGetSurroundingEvents(t *testing.T) {
	t.Run("start", start)
	t.Run("mid", mid)
	t.Run("end", end)
}

func mid(t *testing.T) {
	toFind := &Event{
		Timestamp: 3,
		EventUUID: "3",
	}
	found, err := toFind.GetSurroundingEvents(events)
	if err != nil {
		t.Fatalf("Got an error, '%s'", err)
	}
	if found[0].EventUUID != "2" || found[0].Timestamp != 2 {
		t.Errorf(
			"0th index should match event, eventID: '%s', time: %d",
			found[1].EventUUID,
			found[1].Timestamp,
		)
	}
	if found[1].EventUUID != "3" || found[1].Timestamp != 3 {
		t.Errorf(
			"1st index should match event, eventID: '%s', time: %d",
			found[1].EventUUID,
			found[1].Timestamp,
		)
	}
	if found[2].EventUUID != "4" || found[2].Timestamp != 4 {
		t.Errorf(
			"2nd index should match event, eventID: '%s', time: %d",
			found[2].EventUUID,
			found[2].Timestamp,
		)
	}
}

func start(t *testing.T) {
	toFind := &Event{
		Timestamp: 0,
		EventUUID: "0",
	}
	found, err := toFind.GetSurroundingEvents(events)
	if err != nil {
		t.Fatalf("Got an error, '%s'", err)
	}
	if found[0] != nil {
		t.Errorf("0th index should be nil")
	}
	if found[1].EventUUID != "0" || found[1].Timestamp != 0 {
		t.Errorf(
			"1st index should match event, eventID: '%s', time: %d",
			found[1].EventUUID,
			found[1].Timestamp,
		)
	}
	if found[2].EventUUID != "1" || found[2].Timestamp != 1 {
		t.Errorf("2nd index should be just after event")
	}
}

func end(t *testing.T) {
	toFind := &Event{
		Timestamp: 7,
		EventUUID: "7",
	}
	found, err := toFind.GetSurroundingEvents(events)
	if err != nil {
		t.Fatalf("Got an error, '%s'", err)
	}
	if found[0].EventUUID != "6" || found[0].Timestamp != 6 {
		t.Errorf(
			"0th index should match event, eventID: '%s', time: %d",
			found[0].EventUUID,
			found[0].Timestamp,
		)
	}
	if found[1].EventUUID != "7" || found[1].Timestamp != 7 {
		t.Errorf(
			"1st index should match event, eventID: '%s', time: %d",
			found[1].EventUUID,
			found[1].Timestamp,
		)
	}
	if found[2] != nil {
		t.Errorf("2nd index should be nil")
	}
}
