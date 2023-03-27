package router

import (
	"go-postgres-crud/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/books", controller.GetAllBook).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/book/{id}", controller.GetBook).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/book", controller.AddBook).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/book/{id}", controller.UpdateBook).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/book/{id}", controller.DeleteBook).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/upload/assets", controller.UploadAsset).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/book-validation", controller.AddPageValidation).Methods("POST", "OPTIONS")

	return router
}
