package maxmind

import (
	"secureworks/constants"
	"testing"
)

func TestInitialization(t *testing.T) {
	lf, err := NewLocationFinder(constants.IP_DB)
	defer lf.Destroy()
	if err != nil {
		t.Errorf(
			"Initializing maxmind database at '%s' failed with error '%s'.",
			constants.IP_DB, err,
		)
	}
}

func TestGetLocation(t *testing.T) {
	lf, err := NewLocationFinder(constants.IP_DB)
	defer lf.Destroy()
	if err != nil {
		t.Fatalf(
			"Initializing maxmind database at '%s' failed with error '%s', halting.",
			constants.IP_DB, err,
		)
	}
	ipStr := "81.2.69.142"
	loc, err := lf.GetLocation(ipStr)
	if err != nil {
		t.Fatalf("Failed returning location '%s' with error '%s'.", ipStr, err)
	}
	if loc.Latitude != 52.139801 {
		t.Errorf(
			"Expected latitude: %f, got %f",
			52.139801, loc.Latitude,
		)
	}
	if loc.Longitude != -0.239900 {
		t.Errorf(
			"Expected longitude: %f, got %f",
			-0.239900, loc.Longitude,
		)
	}
	if loc.Radius != 100 {
		t.Errorf(
			"Expected radius: %d, got %d",
			100, loc.Radius,
		)
	}
}
