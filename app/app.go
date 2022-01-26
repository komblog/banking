package app

import "net/http"

func StartServer() {
	http.HandleFunc("/api/customers", GetAllCustomer)

	err := http.ListenAndServe("localhost:8005", nil)
	if err != nil {
		panic(err)
	}
}
