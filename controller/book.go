package controller

import (
	"encoding/json" // package untuk enkode dan mendekode json menjadi struct dan sebaliknya
	// package yang digunakan untuk mengubah string menjadi tipe int
	"log"
	"net/http" // digunakan untuk mengakses objek permintaan dan respons dari api

	"go-postgres-crud/models" //models package dimana Buku didefinisikan

	// digunakan untuk mendapatkan parameter dari router
	_ "github.com/lib/pq" // postgres golang driver
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type Payload struct {
	Data []struct {
		Page    int    `json:"page"`
		Content string `json:"content"`
	} `json:"data"`
}

type Response struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []models.Book `json:"data"`
}
type ResponseData struct {
	Message string   `json:"message"`
	Data    []string `json:"data"`
}

func GetAllBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	title := r.URL.Query().Get("title")

	books, err := models.GetAllBook(string(title))

	if err != nil {
		log.Fatalf("Failed Get Data. %v", err)
	}

	var response Response
	response.Status = 1
	response.Message = "Success"
	response.Data = books

	json.NewEncoder(w).Encode(response)
}
