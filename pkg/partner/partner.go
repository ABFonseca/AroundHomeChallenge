package partner

import (
	"encoding/json"
	"io/ioutil"
	"sort"

	distance "AroundHomeChallenge/pkg/utils"
)

type PartnerList []Partner

type Partner struct {
	Material          []string
	AddressLatitude   float64
	AddressLongitude  float64
	OperatingRadius   int
	Rating            float32
	distanceToRequest float64
}

func (pl PartnerList) Len() int      { return len(pl) }
func (pl PartnerList) Swap(i, j int) { pl[i], pl[j] = pl[j], pl[i] }
func (pl PartnerList) Less(i, j int) bool {
	if pl[i].Rating < pl[j].Rating {
		return false
	}
	if pl[i].Rating > pl[j].Rating {
		return true
	}
	return pl[i].distanceToRequest < pl[j].distanceToRequest
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
	file, _ := ioutil.ReadFile("partner.json")
	partners := PartnerList{}
	_ = json.Unmarshal([]byte(file), &partners)

	return partners
}

func GetPartnersFiltered(material string, lat, lng float64) PartnerList {
	//If we were fetching from a DB I would filter the ones with correct material and if possible filter the distance on DB side as well to make it more efficient
	allPartners := GetAllPartners()

	partners := []Partner{}

	for _, p := range allPartners {
		if p.KnowsMaterial(material) && p.WorksDistance(lat, lng) {
			partners = append(partners, p)
		}
	}

	partnerList := PartnerList(partners)
	sort.Sort(partnerList)

	return partnerList
}
