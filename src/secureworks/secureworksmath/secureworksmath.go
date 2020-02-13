package secureworksmath

import (
	"math"
	"secureworks/event"
)

const (
	//I'm assuming this is close enough....
	RADIUS_EARTH float64 = 6371
	//dimensional analysis of km/sec to mph
	//60sec * .621371 miles
	KMS_TO_MPH = 37.28226
	SUPERMAN   = 500
)

func degsToRads(degs float64) float64 {
	return (degs * math.Pi) / 180
}

//returns kms
//I should test this.... ugh
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

func IsSuperMan(loc1, loc2 *event.Event) bool {
	dis := math.Abs(
		haverSine(loc1, loc2) + float64(loc1.Location.Radius) + float64(loc2.Location.Radius),
	)
	deltaT := math.Abs(
    float64(loc1.Timestamp) - float64(loc2.Timestamp),
  )
	return (dis/deltaT)*KMS_TO_MPH >= SUPERMAN
}

//kms/sec -> miles/hours
