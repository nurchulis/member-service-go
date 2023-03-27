package controller

import (
	"encoding/json" // package untuk enkode dan mendekode json menjadi struct dan sebaliknya
	"fmt"
	"strconv" // package yang digunakan untuk mengubah string menjadi tipe int
	"sync"

	"log"
	"net/http" // digunakan untuk mengakses objek permintaan dan respons dari api

	"go-postgres-crud/models" //models package dimana Buku didefinisikan

	"github.com/gorilla/mux" // digunakan untuk mendapatkan parameter dari router
	_ "github.com/lib/pq"    // postgres golang driver
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

// AddBook
func AddBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Fatalf("error encode body.  %v", err)
	}
	insertID := models.AddBook(book)
	res := response{
		ID:      insertID,
		Message: "Success add Book",
	}
	json.NewEncoder(w).Encode(res)
}

// AddBook
func AddPageValidation(w http.ResponseWriter, r *http.Request) {

	var payload Payload
	var data []string
	var wg sync.WaitGroup
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, d := range payload.Data {
		wg.Add(1)
		go func(index string) {
			defer wg.Done()
			data = append(data, index)
		}(d.Content)
	}
	wg.Wait()

	res := ResponseData{
		Message: "Success",
		Data:    data,
	}
	json.NewEncoder(w).Encode(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Invalid Format ID.  %v", err)
	}
	book, err := models.GetBook(int64(id))
	if err != nil {
		log.Fatalf("Failed Get Data. %v", err)
	}
	json.NewEncoder(w).Encode(book)
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

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Invalid Format ID.  %v", err)
	}
	var book models.Book
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Fatalf("error encode body.  %v", err)
	}
	updatedRows := models.UpdateBook(int64(id), book)
	msg := fmt.Sprintf("Success Update Data %v rows/record", updatedRows)
	res := response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Invalid Format ID.  %v", err)
	}
	deletedRows := models.DeleteBook(int64(id))
	msg := fmt.Sprintf("Success Delete Data %v", deletedRows)
	res := response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}
