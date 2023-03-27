package models

import (
	"database/sql"
	"fmt"
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

type Image struct {
	ID  int64  `json:"id"`
	Url string `json:"url"`
}

type Asset struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	Url       string     `json:"url"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func AddAssets(name string, url string) int64 {
	db := config.CreateConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO assets (name, url)
	VALUES ($1, $2) RETURNING id`
	var id int64
	err := db.QueryRow(sqlStatement,
		name, url).Scan(&id)
	if err != nil {
		log.Fatalf("Failed Exec query. %v", err)
	}
	fmt.Printf("Insert data single record %v", id)
	return id
}

func AddBook(book Book) int64 {
	db := config.CreateConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO books (title, Description, Rating, Image)
	VALUES ($1, $2, $3, $4) RETURNING id`
	var id int64
	err := db.QueryRow(sqlStatement,
		book.Title, book.Description,
		book.Rating, book.Image).Scan(&id)
	if err != nil {
		log.Fatalf("Failed Exec query. %v", err)
	}
	fmt.Printf("Insert data single record %v", id)
	return id
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

// mengambil satu book
func GetBook(id int64) (Book, error) {
	db := config.CreateConnection()
	defer db.Close()
	var book Book
	sqlStatement := `SELECT * FROM books WHERE id=$1`
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&book.ID, &book.Title, &book.Description, &book.Rating, &book.Image, &book.CreatedAt, &book.UpdatedAt)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("Data Not Found")
		return book, nil
	case nil:
		return book, nil
	default:
		log.Fatalf("Cannot get Data. %v", err)
	}
	return book, err
}

// update user in the DB
func UpdateBook(id int64, book Book) int64 {
	db := config.CreateConnection()
	defer db.Close()
	sqlStatement := `UPDATE books SET title=$2, description=$3, rating=$4, image=$5 WHERE id=$1`
	res, err := db.Exec(sqlStatement, id, book.Title, book.Description, book.Rating, book.Image)
	if err != nil {
		log.Fatalf("Failed Exec query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error update. %v", err)
	}
	fmt.Printf("Success Update Data %v\n", rowsAffected)
	return rowsAffected
}

func DeleteBook(id int64) int64 {
	db := config.CreateConnection()
	defer db.Close()
	sqlStatement := `DELETE FROM books WHERE id=$1`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatalf("Failed Exec query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Data not found. %v", err)
	}
	fmt.Printf("Success Delete Data%v", rowsAffected)
	return rowsAffected
}
