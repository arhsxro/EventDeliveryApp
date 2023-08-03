package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/IBM/sarama"
	"github.com/arhsxro/goprojects/eventdelivery/kafka"
)

type Event struct {
	UserID  string `json:"userID"`
	Payload string `json:"payload"`
}

func HandleSingleEvent(w http.ResponseWriter, r *http.Request) {
	var event Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, "Failed to parse event", http.StatusBadRequest)
		return
	}

	maxRetries := 5
	retryInterval := time.Second

	var producer sarama.SyncProducer
	var kafkaErr error

	for i := 1; i <= maxRetries; i++ {
		producer, kafkaErr = kafka.GetKafkaProducer()
		if kafkaErr == nil {
			break // Successfully connected, break out of the retry loop
		}

		// Retry after a delay
		log.Printf("Failed to create Kafka producer. Retry attempt %d\n", i)
		time.Sleep(retryInterval)
	}

	if kafkaErr != nil {
		// If still unsuccessful after max retries, exit with an error
		fmt.Println("Failed to create Kafka producer after multiple retries. Exiting.")
		log.Println(kafkaErr)
		return
	}

	// Publish event to Kafka topic
	err = kafka.PublishEvent(producer, kafka.Event(event))
	if err != nil {
		http.Error(w, "Failed to publish event", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, "Event received and published to Kafka")

	defer producer.Close()
}

func HandleMultipleEvents(w http.ResponseWriter, r *http.Request) {
	var events []Event
	err := json.NewDecoder(r.Body).Decode(&events)
	if err != nil {
		http.Error(w, "Failed to parse events", http.StatusBadRequest)
		return
	}

	maxRetries := 5
	retryInterval := time.Second

	var producer sarama.SyncProducer
	var kafkaErr error

	for i := 1; i <= maxRetries; i++ {
		producer, kafkaErr = kafka.GetKafkaProducer()
		if kafkaErr == nil {
			break
		}

		fmt.Printf("Failed to create Kafka producer. Retry attempt %d\n", i)
		time.Sleep(retryInterval)
	}

	if kafkaErr != nil {
		fmt.Println("Failed to create Kafka producer after multiple retries. Exiting.")
		log.Println(kafkaErr)
		return
	}

	for _, event := range events {
		err := kafka.PublishEvent(producer, kafka.Event(event))
		if err != nil {
			http.Error(w, "Failed to publish event", http.StatusInternalServerError)
			return
		}
	}

	fmt.Fprintf(w, "Events received and published to Kafka")
}
