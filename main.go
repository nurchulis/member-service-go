package main

import (
	"fmt"
	"go-postgres-crud/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()

	fmt.Println("Running Server on Port 8081...")

	log.Fatal(http.ListenAndServe(":8081", r))
}
