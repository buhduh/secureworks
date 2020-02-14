package database

import (
	"fmt"
	"secureworks/event"
	"testing"
)

var dbEvents []*event.Event = []*event.Event{
	//apple is 111.eventnum
	//bob is 222.eventnum
	&event.Event{
		IP:        "222.333",
		Timestamp: 3,
		UserName:  "bob",
		EventUUID: "bob3",
		Location: &event.Location{
			Latitude:  222.3331,
			Longitude: 222.3332,
			Radius:    333,
		},
	},
	&event.Event{
		IP:        "111.111",
		Timestamp: 1,
		UserName:  "apple",
		EventUUID: "apple1",
		Location: &event.Location{
			Latitude:  111.1111,
			Longitude: 111.1112,
			Radius:    111,
		},
	},
	&event.Event{
		IP:        "111.333",
		Timestamp: 3,
		UserName:  "apple",
		EventUUID: "apple3",
		Location: &event.Location{
			Latitude:  111.3331,
			Longitude: 111.3332,
			Radius:    333,
		},
	},
	&event.Event{
		IP:        "222.222",
		Timestamp: 2,
		UserName:  "bob",
		EventUUID: "bob2",
		Location: &event.Location{
			Latitude:  222.2221,
			Longitude: 222.2222,
			Radius:    222,
		},
	},
	&event.Event{
		IP:        "222.111",
		Timestamp: 1,
		UserName:  "bob",
		EventUUID: "bob1",
		Location: &event.Location{
			Latitude:  222.1111,
			Longitude: 222.1112,
			Radius:    111,
		},
	},
	&event.Event{
		IP:        "111.222",
		Timestamp: 2,
		UserName:  "apple",
		EventUUID: "apple2",
		Location: &event.Location{
			Latitude:  111.2221,
			Longitude: 111.2222,
			Radius:    222,
		},
	},
}

func TestDB(t *testing.T) {
	t.Run("legit", legit)
	t.Run("dupeEvent", dupeEvent)
}

func dupeEvent(t *testing.T) {
	e := &event.Event{
		IP:        "111.222",
		Timestamp: 2,
		UserName:  "apple",
		EventUUID: "apple2",
		Location: &event.Location{
			Latitude:  111.2221,
			Longitude: 111.2222,
			Radius:    222,
		},
	}
	db, err := NewDatabase()
	if err != nil {
		t.Fatalf("Need a 'real' database, got error '%s'.", err)
	}
	err = db.NewEvent(e)
	if err != nil {
		t.Errorf("dupe events are OK, got error '%s'", err)
	}
	allEvents, err := db.GetOrderedEventsForUser("apple")
	if err != nil {
		t.Fatalf("should not get an error, got '%s'", err)
	}
	if len(allEvents) != 3 {
		t.Errorf("dupe events should not be added")
	}
}

func legit(t *testing.T) {
	db, err := NewDatabase()
	if err != nil {
		t.Fatalf("Need a 'real' database, got error '%s'.", err)
	}
	for _, e := range dbEvents {
		if err = db.NewEvent(e); err != nil {
			t.Errorf("failed adding event for event:\n%swith error '%s'.", e, err)
		}
	}
	//Awesome, the events are added, let's check if they're ordered
	bobs, err := db.GetOrderedEventsForUser("bob")
	if err != nil {
		t.Errorf("I shouldn't fail, got error: '%s'.", err)
	}
	apples, err := db.GetOrderedEventsForUser("apple")
	if err != nil {
		t.Errorf("I shouldn't fail, got error: '%s'.", err)
	}
	if len(bobs) != 3 {
		t.Errorf("user bob should have 3 events, got %d events.", len(bobs))
	}
	if len(apples) != 3 {
		t.Errorf("user apple should have 3 events, got %d events.", len(apples))
	}
	for i, b := range bobs {
		num := i + 1
		if b.Timestamp != uint(num) {
			t.Errorf("Timestamp mismatch, expected %d, got %d.", num, b.Timestamp)
		}
		eventUUID := fmt.Sprintf("bob%d", num)
		if b.EventUUID != eventUUID {
			t.Errorf(
				"event ids do not match, expected '%s', got '%s'.",
				eventUUID, b.EventUUID,
			)
		}
		if b.UserName != "bob" {
			t.Errorf("Name mismatch, expected '%s', got '%s'.", "bob", b.UserName)
		}
	}
}
