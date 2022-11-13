package routes

import (
	"backend-journey/handlers"
	"backend-journey/pkg/middleware"
	"backend-journey/pkg/mysql"
	"backend-journey/repositories"

	"github.com/gorilla/mux"
)

func JourneyRoutes(r *mux.Router) {
	journeyRepository := repositories.RepositoryJourney(mysql.DB)
	h := handlers.HandlerJourney(journeyRepository)

	r.HandleFunc("/journey", middleware.Auth(middleware.UploadFile(h.CreateJourney))).Methods("POST")
	r.HandleFunc("/journeys", h.FindJourneys).Methods("GET")
	r.HandleFunc("/journey/{id}", h.GetJourney).Methods("GET")
}
