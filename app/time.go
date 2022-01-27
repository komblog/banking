package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
)

type CurrentTime struct {
	Location    string    `json:"location"`
	CurrentTime time.Time `json:"current_time"`
}

func GetCurrentTime(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	tz := request.URL.Query().Get("tz")
	response := make(map[string]string, 0)
	timeZone := strings.Split(tz, ",")

	if len(timeZone) <= 1 {
		loc, err := time.LoadLocation(tz)
		if err != nil {
			writer.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(writer, "Invalid timezone %s", tz)
		} else {
			response[tz] = time.Now().In(loc).String()
			writer.Header().Add("Content-Type", "application/json")
			json.NewEncoder(writer).Encode(response)
		}
	} else {
		for _, tZone := range timeZone {
			loc, err := time.LoadLocation(tZone)
			if err != nil {
				writer.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(writer, "invalid timezone %s", tZone)
				return
			} else {
				response[tZone] = time.Now().In(loc).String()
			}
		}
		writer.Header().Add("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(response)
	}
}
