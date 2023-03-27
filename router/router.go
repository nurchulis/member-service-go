package router

import (
	"go-postgres-crud/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/pubsub", controller.GetAllBook).Methods("GET", "OPTIONS")

	return router
}
