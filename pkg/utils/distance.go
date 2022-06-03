package distance

import (
	"math"
)

func Distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64) float64 {
	phi1 := float64(math.Pi * lat1 / 180)
	phi2 := float64(math.Pi * lat2 / 180)

	deltaPhi := (lat2 - lat1) * math.Pi / 180
	deltaLambda := (lng2 - lng1) * math.Pi / 180

	a := math.Pow(math.Sin(deltaPhi/2), 2) + (math.Cos(phi1) * math.Cos(phi2) * math.Pow(math.Sin(deltaLambda/2), 2))
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	R := 6371e3 // earth radius in metres
	d := R * c

	return d / 1000 // return distance in km

}
