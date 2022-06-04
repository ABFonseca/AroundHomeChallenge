package distance

import (
	"fmt"
	"math"
	"testing"
)

type distanceTests struct {
	lat1 float64
	lng1 float64
	lat2 float64
	lng2 float64
	want float64
	name string
}

func assertFloat(t *testing.T, got, want float64, testName string) {
	t.Helper()
	if got != want {
		failString := fmt.Sprintln("Test", testName, "got:", got, "wanted:", want)
		t.Fatalf(failString)
	}
}

func TestDistance(t *testing.T) {
	originLat := 52.50879681532554
	originLng := 13.375567271135349
	testCases := []distanceTests{
		{originLat, originLng, originLat, originLng, 0, "Test lattitude and longitue are equal"},
		{originLat, originLng, originLat, 13.9666065, 40, "Test lattitude change"},
		{originLat, originLng, 52.382, originLng, 14.099, "Test Longitude change"},
		{originLat, originLng, 52.382, 13.9666065, 42.466, "Test Longitude and lattitude change"},
	}

	for _, test := range testCases {
		//test with a percision of 3 decimal places

		got := Distance(test.lat1, test.lng1, test.lat2, test.lng2)
		got = math.Round(got*1000) / 1000
		assertFloat(t, got, test.want, test.name)
	}
}
