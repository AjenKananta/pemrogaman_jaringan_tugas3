package routes

import (
	"backend/controllers"

	"github.com/gorilla/mux"
)

func RegisterEventRoutes(router *mux.Router) {
	router.HandleFunc("/events", controllers.GetAllEvents).Methods("GET")
	router.HandleFunc("/events/{id}", controllers.GetEventByID).Methods("GET")
	router.HandleFunc("/events", controllers.CreateEvent).Methods("POST")
	router.HandleFunc("/events/{id}", controllers.UpdateEvent).Methods("PUT")
	router.HandleFunc("/events/{id}", controllers.DeleteEvent).Methods("DELETE")
}
