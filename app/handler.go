package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}

func GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	customer := []Customer{
		{Name: "Hery", City: "Sibolga", ZipCode: "22522"},
		{Name: "Gabriel", City: "Kolang", ZipCode: "22352"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customer)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}
