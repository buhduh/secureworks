package secureworksmath

import (
	"math"
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
