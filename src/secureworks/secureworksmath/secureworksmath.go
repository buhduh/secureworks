package secureworksmath

import (
	"math"
	"secureworks/event"
)

const (
	//I'm assuming this is close enough....
	RADIUS_EARTH float64 = 6371
	//dimensional analysis of km/sec to mph
	//3600sec * .621371 miles
	KMS_TO_MPH = 2236.9356
	SUPERMAN   = 500
)

func degsToRads(degs float64) float64 {
	return (degs * math.Pi) / 180
}

//returns kms
func haverSine(loc1, loc2 *event.Event) float64 {
	lat1 := degsToRads(loc1.Location.Latitude)
	lat2 := degsToRads(loc2.Location.Latitude)
	lon1 := degsToRads(loc1.Location.Longitude)
	lon2 := degsToRads(loc2.Location.Longitude)
	sin2Lat := math.Pow(math.Sin((lat2-lat1)/2), 2)
	sin2Lon := math.Pow(math.Sin((lon2-lon1)/2), 2)
	cosTerm := math.Cos(lat1) * math.Cos(lat2)
	return 2 * RADIUS_EARTH * math.Asin(math.Sqrt(
		sin2Lat+cosTerm*sin2Lon,
	))
}

//MPH
func Speed(loc1, loc2 *event.Event) float64 {
	haverSine := haverSine(loc1, loc2)
	rad1 := float64(loc1.Location.Radius)
	rad2 := float64(loc2.Location.Radius)
	dis := haverSine - (rad1 + rad2)
	//inside radii, within error bounds
	if dis <= 0 {
		dis = 0
	}
	deltaT := float64(loc1.Timestamp) - float64(loc2.Timestamp)
	//Superman is infinitely fast
	if deltaT == 0 {
		return SUPERMAN + 1
	}
	return math.Abs((dis / deltaT) * KMS_TO_MPH)
}

func IsSuperMan(speed float64) bool {
	return speed >= SUPERMAN
}
