package main

import (
	"log"
	"net/http"
	"time"

	"github.com/IBM/sarama"
	"github.com/arhsxro/goprojects/eventdelivery/api"
	"github.com/arhsxro/goprojects/eventdelivery/kafka"
)

type Event struct {
	UserID  string `json:"userID"`
	Payload string `json:"payload"`
}

func main() {

	// Set up Kafka consumer,max retries 5 before terminate

	maxRetries := 5
	retryInterval := time.Second

	var consumer sarama.Consumer
	var kafkaErr error

	for i := 1; i <= maxRetries; i++ {
		consumer, kafkaErr = kafka.GetKafkaConsumer()
		if kafkaErr == nil {
			break // Successfully connected, break out of the retry loop
		}

		// Retry after a delay
		log.Printf("Failed to create Kafka consumer. Retry attempt %d ", i)
		time.Sleep(retryInterval)
	}

	if kafkaErr != nil {
		// If still unsuccessful after max retries, exit
		log.Printf("Failed to create Kafka consumer after multiple retries. Exiting.")
		log.Println(kafkaErr)
		return
	}

	// Start consuming events
	go kafka.ConsumeEvents(consumer)

	// Set up HTTP endpoint to receive events
	http.HandleFunc("/api/singleEvent", api.HandleSingleEvent)
	http.HandleFunc("/api/multipleEvents", api.HandleMultipleEvents)

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
