package kafka

import (
	"encoding/json"

	"github.com/IBM/sarama"
)

type Event struct {
	UserID  string `json:"userID"`
	Payload string `json:"payload"`
}

func PublishEvent(producer sarama.SyncProducer, event Event) error {
	message, err := json.Marshal(event)
	if err != nil {
		return err
	}

	producerMessage := &sarama.ProducerMessage{
		Topic: "events_topic",
		Value: sarama.ByteEncoder(message),
	}

	_, _, err = producer.SendMessage(producerMessage) // ignore partition and offset and focus on the error
	return err
}

func GetKafkaProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 3
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	producer, err := sarama.NewSyncProducer([]string{"kafka:9092"}, config)
	if err != nil {
		return nil, err
	}

	return producer, nil
}
