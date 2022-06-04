package api

import (
	"AroundHomeChallenge/pkg/partner"
	"encoding/json"
	"fmt"
	"net/http"

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
	router.HandleFunc("/request", CustomerRequest).Methods("POST")
	router.HandleFunc("/partner", listPartners).Methods("GET")

	// Start Server
	fmt.Println("Starting API server on port 8080")
	http.ListenAndServe(":8080", router)
}

func CustomerRequest(w http.ResponseWriter, r *http.Request) {
	//reqBody, _ := ioutil.ReadAll(r.Body)
	var req request
	//err := json.Unmarshal(reqBody, &req)
	err := json.NewDecoder(r.Body).Decode(&req)

	//fmt.Println("Got Customer requestBody: ", reqBody)
	fmt.Println("Got Customer request: ", req)

	if err != nil {
		fmt.Println("Error parsing json:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	partners := partner.GetPartnersFiltered(req.Material, req.AddressLat, req.AddressLng)

	fmt.Println("Partner list:", partners)

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(partners)
}

func listPartners(w http.ResponseWriter, r *http.Request) {
	partners := partner.GetAllPartners()

	fmt.Println("Got Partners request")
	fmt.Println("Partner list:", partners)

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(partners)

}
