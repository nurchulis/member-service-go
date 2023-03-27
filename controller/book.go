package controller

import (
	"log"

	_ "github.com/lib/pq" // postgres golang driver
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func SendEmailNotification() {
	log.Println("DISINI")
}
