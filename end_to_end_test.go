package partner_test

//In terms of organization this test file should be on the partner folder, but because I'm using json the relative path fails
//since this is not how it would work in produciton I didn't dedicate much time in solving the relative path issue
//This also makes the coverage for partner package to be <20% when in reality we do cover more

import (
	. "AroundHomeChallenge/pkg/partner"
	"fmt"
	"testing"
)

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

func TestRequestListempty(t *testing.T) {
	reqLat := 2.50879681532554
	reqLng := -13.375567271135349
	partners := GetPartnersFiltered("wood", reqLat, reqLng)

	if len(partners.Partners) != 0 {
		t.Fatalf("Fail: the list should be empty")
	}

}

func TestRequestListOrder(t *testing.T) {
	reqLat := 52.50879681532554
	reqLng := 13.375567271135349
	partners := GetPartnersFiltered("wood", reqLat, reqLng)

	if len(partners.Partners) != 2 {
		failString := fmt.Sprintln("Fail: the list should be 2 and is", len(partners.Partners))
		t.Fatalf(failString)
	}

	if partners.Partners[0].Id != 624 {
		failString := fmt.Sprintln("Fail: the first partner id in the list should be 624 and is", partners.Partners[0].Id)
		t.Fatalf(failString)
	}

	if partners.Partners[1].Id != 13 {
		failString := fmt.Sprintln("Fail: the second partner id in the list should be 13 and is", partners.Partners[0].Id)
		t.Fatalf(failString)
	}

}

func TestRequestListOrderTwo(t *testing.T) {
	reqLat := 52.50879681532554
	reqLng := 13.375567271135349
	partners := GetPartnersFiltered("carpet", reqLat, reqLng)

	if len(partners.Partners) != 5 {
		failString := fmt.Sprintln("Fail: the list should be 5 and is", len(partners.Partners))
		t.Fatalf(failString)
	}

	if partners.Partners[0].Id != 624 {
		failString := fmt.Sprintln("Fail: the first partner id in the list should be 624 and is", partners.Partners[0].Id)
		t.Fatalf(failString)
	}

	if partners.Partners[1].Id != 13 {
		failString := fmt.Sprintln("Fail: the second partner id in the list should be 13 and is", partners.Partners[1].Id)
		t.Fatalf(failString)
	}

	if partners.Partners[2].Id != 552 {
		failString := fmt.Sprintln("Fail: the third partner id in the list should be 552 and is", partners.Partners[2].Id)
		t.Fatalf(failString)
	}

	if partners.Partners[3].Id != 199 {
		failString := fmt.Sprintln("Fail: the forth partner id in the list should be 199 and is", partners.Partners[3].Id)
		t.Fatalf(failString)
	}

	if partners.Partners[4].Id != 5 {
		failString := fmt.Sprintln("Fail: the fifth partner id in the list should be 5 and is", partners.Partners[4].Id)
		t.Fatalf(failString)
	}

}

func TestGetAllPartners(t *testing.T) {
	partners := GetAllPartners()
	if len(partners.Partners) != 7 {
		failString := fmt.Sprintln("Partner List not correct, should have 7 and has ", len(partners.Partners))
		t.Fatalf(failString)
	}

}

func TestGetPartnerDetail(t *testing.T) {
	p, ok := GetPartnerDetails(552)
	if !ok {
		failString := fmt.Sprintln("Partner 552 not found")
		t.Fatalf(failString)
	}

	if p.Name != "Mr. Carpet" {
		failString := fmt.Sprintln("Partner Name not correct, should be 'Mr.Carpet' and is ", p.Name)
		t.Fatalf(failString)
	}

}
