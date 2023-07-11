package main

import (
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	// payload response
	payload := JsonResponse{
		Error:   false,
		Message: "Hello from broker-service",
		// data belum adas
		Data: "belum ada data",
	}
	// for write json
	_ = app.writeJSON(w, http.StatusOK, payload)

	//out, _ := json.MarshalIndent(payload, "", "\t")
	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusAccepted)
	//w.Write(out)
}
