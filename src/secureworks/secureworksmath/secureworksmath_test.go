package secureworksmath

import (
	"math"
	"secureworks/event"
	"testing"
)

func TestDegsToRads(t *testing.T) {
	pi := degsToRads(180)
	halfpi := degsToRads(90)
	threepiover2 := degsToRads(270)
	if pi != math.Pi {
		t.Errorf("Expected %f, got %f.", math.Pi, pi)
	}
	if halfpi != math.Pi/2 {
		t.Errorf("Expected %f, got %f.", math.Pi/2, halfpi)
	}
	if threepiover2 != (3*math.Pi)/2 {
		t.Errorf("Expected %f, got %f.", (3*math.Pi)/2, threepiover2)
	}
}

func TestHaverSine(t *testing.T) {
	locs := [][2]*event.Event{
		[2]*event.Event{
			&event.Event{
				Location: &event.Location{
					Latitude:  0,
					Longitude: 90,
				},
			},
			&event.Event{
				Location: &event.Location{
					Latitude:  0,
					Longitude: 0,
				},
			},
		},
	}
	for _, l := range locs {
		dis := haverSine(l[0], l[1])
		exp := (math.Pi * RADIUS_EARTH) / 2
		//close enough
		if int(dis) != int(exp) {
			t.Errorf("expected %f, got %f", exp, dis)
		}
	}
}
