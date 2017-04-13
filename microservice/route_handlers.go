package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {

	//healthCheckResponse is the default response for json marshalling for the /health route
	var healthCheckResponse = HealthCheckUp{
		Status: Config.Routing.HealthCheckStatusUp,
		Application: ApplicationStatus{
			Status: Config.Routing.HealthCheckStatusUp,
			Name:   Config.Microservice.Name,
		},
	}

	if err := json.NewEncoder(w).Encode(healthCheckResponse); err != nil {
		log.Println("JSON encoder error: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
