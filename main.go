package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go-postgres-crud/libs"
	"go-postgres-crud/models"
)

func main() {
	// Create a Redis client
	client := libs.NewClient()

	// Subscribe to the "example" channel
	pubsub := client.Subscribe(context.Background(), "notification")

	// Wait for confirmation that we're subscribed
	_, err := pubsub.Receive(context.Background())
	if err != nil {
		log.Fatalf("Failed to subscribe to channel: %v", err)
	}

	// Create a channel to receive messages on
	ch := pubsub.Channel()

	// Listen for messages on the channel
	for msg := range ch {
		// Parse the JSON string into a Person struct
		var book models.Book
		err := json.Unmarshal([]byte(msg.Payload), &book)
		if err != nil {
			log.Fatalf("Failed to unmarshal JSON data: %v", err)
		}

		// Print the name of the person
		fmt.Printf("Received message: Name=%s\n", book.Title)
	}

	// Close the pubsub channel when done
	pubsub.Close()

	// Close the Redis client when done
	client.Close()
}
