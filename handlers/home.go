package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ernestomr87/go-rest-websockets/server"
)

type HomeResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func HomeHandler(s server.Server) http.HandlerFunc {
	log.Println("HomeHandler")
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Contet-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(HomeResponse{
			Message: "Welcome to Platzi Go",
			Status:  true,
		})
	}
}
