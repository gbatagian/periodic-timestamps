package ptlist

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func PtListGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			// Map unknown errors to 400
			if r := recover(); r != nil {
				http.Error(w, fmt.Sprintf("%v", r), http.StatusBadRequest)
				log.Printf("Recovered from panic: %v", r)
			}
		}()

		params, err := (&QueryParams{}).ParseURL(r.URL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		timestamps := PeriodicTimestamps(
			params.Period,
			params.T1,
			params.T2,
			params.Tz,
		)

		response := (&PtListGetResponse{}).FromTimestampsSlice(timestamps)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response.Ptlist)
	}
}
