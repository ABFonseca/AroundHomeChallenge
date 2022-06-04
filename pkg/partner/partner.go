package partner

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"

	distance "AroundHomeChallenge/pkg/utils"
)

type PartnerList struct {
	Partners []Partner `json:"Partners"`
}

type Partner struct {
	Material          []string `json:"Material"`
	AddressLatitude   float64  `json:"AddressLatitude"`
	AddressLongitude  float64  `json:"AddressLongitude"`
	OperatingRadius   int      `json:"OperatingRadius"`
	Rating            float32  `json:"Rating"`
	distanceToRequest float64  `json:"DistanceToRequest,omitempty"`
}

func (pl PartnerList) Len() int      { return len(pl.Partners) }
func (pl PartnerList) Swap(i, j int) { pl.Partners[i], pl.Partners[j] = pl.Partners[j], pl.Partners[i] }
func (pl PartnerList) Less(i, j int) bool {
	if pl.Partners[i].Rating < pl.Partners[j].Rating {
		return false
	}
	if pl.Partners[i].Rating > pl.Partners[j].Rating {
		return true
	}
	return pl.Partners[i].distanceToRequest < pl.Partners[j].distanceToRequest
}

//Returns true if the requested material is on the list of materials known to the partner
func (p Partner) KnowsMaterial(material string) bool {
	for _, mat := range p.Material {
		if mat == material {
			return true
		}
	}
	return false
}

//Calculates distance to request and stores it on the object
//returns if distance is below the OperatingRadius limit
func (p *Partner) WorksDistance(lat, lng float64) bool {
	dist := distance.Distance(lat, lng, p.AddressLatitude, p.AddressLongitude)
	p.distanceToRequest = dist
	return dist <= float64(p.OperatingRadius)
}

//Currently fetching from JSON file for challenge purposes, but in production code this would fetch from a dynamic source like a database
//Even though in the challenge I'm getting from a static source, I'll abstract from that and do all the behaviour as if I was fetching from a dynamic one
// for exapmple, for each request I will read/filter the source to get a updated list of partners and their expertises
func GetAllPartners() PartnerList {

	jsonFile, err := os.Open("pkg/partner/partner.json")
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of the jsonFile
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var partners PartnerList

	err = json.Unmarshal(byteValue, &partners)
	if err != nil {
		fmt.Println("Error fetching partner list with error:", err)
	}

	fmt.Println("in GetAllPartners:", partners)

	return partners
}

func GetPartnersFiltered(material string, lat, lng float64) PartnerList {
	//If we were fetching from a DB I would filter the ones with correct material and if possible filter the distance on DB side as well to make it more efficient
	allPartners := GetAllPartners()

	partners := []Partner{}

	for _, p := range allPartners.Partners {
		if p.KnowsMaterial(material) && p.WorksDistance(lat, lng) {
			fmt.Println("knows material as is within distance:", p)
			partners = append(partners, p)
		} else {
			fmt.Println("distance is: ", p.distanceToRequest, " and radius:", p.OperatingRadius)
		}
	}

	partnerList := PartnerList{partners}
	fmt.Println("unsorted list:", partnerList)
	sort.Sort(partnerList)
	fmt.Println("sorted list:", partnerList)

	return partnerList
}
