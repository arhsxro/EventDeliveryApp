package kafka

import (
	"encoding/json"
	"log"
	"time"

	"github.com/IBM/sarama"

	"github.com/arhsxro/goprojects/eventdelivery/delivery"
)

func ConsumeEvents(consumer sarama.Consumer) {
	retryInterval := time.Second

	partitionConsumer, err := consumer.ConsumePartition("events_topic", 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatal("Failed to start consuming events:", err)
	}
	defer partitionConsumer.Close()

	for message := range partitionConsumer.Messages() {
		var event Event
		err := json.Unmarshal(message.Value, &event)
		if err != nil {
			log.Println("Failed to unmarshal event:", err)
			continue
		}

		// Simulate event delivery to destinations
		err = delivery.DeliverToDestinations(delivery.Event(event))
		if err != nil {
			log.Println("Failed to deliver event to destinations:", err)
			// Implement retry logic

			maxRetries := 5
			for retry := 1; retry <= maxRetries; retry++ {
				time.Sleep(retryInterval * time.Second)

				err = delivery.DeliverToDestinations(delivery.Event(event))
				if err == nil {
					break
				}

				if retry == maxRetries {
					log.Println("Giving up on event delivery after maximum retries")
				}
				retryInterval++
			}
		}

		log.Println("Event delivered:", event)
	}
}

func GetKafkaConsumer() (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, err := sarama.NewConsumer([]string{"kafka:9092"}, config)
	if err != nil {
		return nil, err
	}

	return consumer, nil
}
