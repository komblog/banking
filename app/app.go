package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func StartServer() {
	router := httprouter.New()
	router.GET("/api/customers", GetAllCustomer)
	router.GET("/api/time", GetCurrentTime)

	server := http.Server{
		Addr:    "localhost:8005",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
