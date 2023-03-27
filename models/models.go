package models

import (
	"go-postgres-crud/config"
	"log"
	"time"

	_ "github.com/lib/pq" // driver pgsql
)

type Book struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"Description"`
	Rating      *int       `json:"rating"`
	Image       string     `json:"image"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

// ambil satu book
func GetAllBook(title string) ([]Book, error) {
	db := config.CreateConnection()
	defer db.Close()
	var books []Book
	sqlStatement := `SELECT * FROM books WHERE title LIKE '%' || $1 || '%'`
	rows, err := db.Query(sqlStatement, title)
	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var book Book
		err = rows.Scan(&book.ID, &book.Title, &book.Description, &book.Rating, &book.Image, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			log.Fatalf("tidak bisa mengambil data. %v", err)
		}
		books = append(books, book)
	}
	return books, err
}
