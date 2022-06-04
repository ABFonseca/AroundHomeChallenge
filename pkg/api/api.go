package api

import (
	"AroundHomeChallenge/pkg/partner"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type request struct {
	Material     string  `json:"Material"`
	AddressLat   float64 `json:"Latitude"`
	AddressLng   float64 `json:"Longitude"`
	SquareMeters int     `json:"SquareMeters"`
	Phone        uint    `json:"Phone"`
}

func StartWebservice() {
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/request", customerRequest).Methods("POST")
	router.HandleFunc("/partner", listPartners).Methods("GET")
	router.HandleFunc("/partner/{partner_id}", getPartner).Methods("GET")

	// Start Server
	fmt.Println("Starting API server on port 8080")
	http.ListenAndServe(":8080", router)
}

func customerRequest(w http.ResponseWriter, r *http.Request) {
	//reqBody, _ := ioutil.ReadAll(r.Body)
	var req request
	//err := json.Unmarshal(reqBody, &req)
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		fmt.Println("Error parsing json:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	partners := partner.GetPartnersFiltered(req.Material, req.AddressLat, req.AddressLng)

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(partners)
}

func listPartners(w http.ResponseWriter, r *http.Request) {
	partners := partner.GetAllPartners()

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(partners)

}

func getPartner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	partner_id, err := strconv.Atoi(vars["partner_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	partner, ok := partner.GetPartnerDetails(partner_id)
	if ok == false {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(partner)
}
