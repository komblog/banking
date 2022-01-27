package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}

func GetAllCustomer(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customer := []Customer{
		{Name: "Hery", City: "Sibolga", ZipCode: "22522"},
		{Name: "Gabriel", City: "Kolang", ZipCode: "22352"},
	}

	if request.Header.Get("Content-Type") == "application/xml" {
		writer.Header().Add("Content-Type", "application/xml")
		encoder := xml.NewEncoder(writer)
		err := encoder.Encode(&customer)

		if err != nil {
			panic(err)
		}

	} else {
		writer.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(writer)
		err := encoder.Encode(&customer)

		if err != nil {
			panic(err)
		}
	}
}
