package api

import (
	"AroundHomeChallenge/pkg/partner"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type request struct {
	material     string  `json:"Material"`
	addressLat   float64 `json:"Latitude"`
	addressLng   float64 `json:"Longitude"`
	squareMeters int     `json:"SquareMeters"`
	phone        uint    `json:"Phone"`
}

func StartWebservice() {
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/request", CustomerRequest).Methods("POST")
	router.HandleFunc("/partner", listPartners).Methods("GET")

	// Start Server
	http.ListenAndServe(":8080", router)
}

func CustomerRequest(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var req request
	err := json.Unmarshal(reqBody, &req)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	partners := partner.GetPartnersFiltered(req.material, req.addressLat, req.addressLng)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(partners)
}

func listPartners(w http.ResponseWriter, r *http.Request) {
	partners := partner.GetAllPartners()
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(partners)

}
