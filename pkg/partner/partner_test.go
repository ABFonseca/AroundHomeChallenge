package partner

import (
	"fmt"
	"testing"
)

// request address: 52.50879681532554, 13.375567271135349

type knowsMaterialTests struct {
	partner  Partner
	material string
	want     bool
	name     string
}

type worksDistanceTests struct {
	partner Partner
	lat     float64
	lng     float64
	want    bool
	name    string
}

func assertBool(t *testing.T, got, want bool, testName string) {
	if got != want {
		failString := fmt.Sprintln("Test", testName, "got:", got, "wanted:", want)
		t.Fatalf(failString)
	}
}

func TestKnowsMaterial(t *testing.T) {
	testCases := []knowsMaterialTests{
		{Partner{Material: []string{"wood", "tiles", "carpet"}}, "wood", true, "Search in first element"},
		{Partner{Material: []string{"wood", "tiles", "carpet"}}, "carpet", true, "Search in last element"},
		{Partner{Material: []string{"wood", "tiles", "carpet"}}, "tiles", true, "Search in middle element"},
		{Partner{Material: []string{"wood", "tiles", "carpet"}}, "concrete", false, "Searched element doesn't exist"},
		{Partner{Material: []string{}}, "wood", false, "Search empty list"},
	}

	for _, test := range testCases {
		got := test.partner.KnowsMaterial(test.material)
		assertBool(t, got, test.want, test.name)
	}

}

func TestWorksDistance(t *testing.T) {
	reqLat := 52.50879681532554
	reqLng := 13.375567271135349
	testCases := []worksDistanceTests{
		{
			Partner{AddressLatitude: reqLat, AddressLongitude: reqLng, OperatingRadius: 25},
			reqLat,
			reqLng,
			true,
			"Request address 0m away from partner",
		},
		{
			Partner{AddressLatitude: reqLat, AddressLongitude: 13.575567271135349, OperatingRadius: 25},
			reqLat,
			reqLng,
			true,
			"Position to East within range",
		},
		{
			Partner{AddressLatitude: reqLat, AddressLongitude: 13.975567271135349, OperatingRadius: 25},
			reqLat,
			reqLng,
			true,
			"Position to East out of range",
		},
		{
			Partner{AddressLatitude: reqLat, AddressLongitude: 13.175567271135349, OperatingRadius: 25},
			reqLat,
			reqLng,
			true,
			"Position to West within range",
		},
		{
			Partner{AddressLatitude: reqLat, AddressLongitude: 13.005567271135349, OperatingRadius: 25},
			reqLat,
			reqLng,
			true,
			"Position to West barely out of range",
		},
		{
			Partner{AddressLatitude: 52.70879681532554, AddressLongitude: reqLng, OperatingRadius: 25},
			reqLat,
			reqLng,
			true,
			"Position to North within range",
		},
		{
			Partner{AddressLatitude: 52.70879681532554, AddressLongitude: reqLng, OperatingRadius: 20},
			reqLat,
			reqLng,
			true,
			"Position to North out of range",
		},
		{
			Partner{AddressLatitude: 52.30879681532554, AddressLongitude: reqLng, OperatingRadius: 25},
			reqLat,
			reqLng,
			true,
			"Position to North within range",
		},
		{
			Partner{AddressLatitude: 52.30879681532554, AddressLongitude: reqLng, OperatingRadius: 20},
			reqLat,
			reqLng,
			true,
			"Position to North out of range",
		},
		{
			Partner{AddressLatitude: 52.40879681532554, AddressLongitude: 13.675567271135349, OperatingRadius: 25},
			reqLat,
			reqLng,
			true,
			"Change Latitude and Longitude within range",
		},
		{
			Partner{AddressLatitude: 52.40879681532554, AddressLongitude: 13.775567271135349, OperatingRadius: 25},
			reqLat,
			reqLng,
			true,
			"Change Latitude and Longitude out of range",
		},
	}

	for _, test := range testCases {
		got := test.partner.WorksDistance(test.lat, test.lng)
		assertBool(t, got, test.want, test.name)
	}

}
