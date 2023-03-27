package models

import (
	_ "github.com/lib/pq" // driver pgsql
)

type Book struct {
	Title string `json:"title"`
}
