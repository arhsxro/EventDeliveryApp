package delivery

import (
	"errors"
	"log"
	"time"
)

type Event struct {
	UserID  string `json:"userID"`
	Payload string `json:"payload"`
}

func DeliverToDestinations(event Event) error {
	// Simulate delivery to destinations
	err := DeliverToDestination1(event)
	if err != nil {
		return err
	}

	err = DeliverToDestination2(event)
	if err != nil {
		return err
	}

	err = DeliverToDestination3(event)
	if err != nil {
		return err
	}

	return nil
}

func DeliverToDestination1(event Event) error {
	if event.UserID == "123" {
		return errors.New("delivery failure to destination 1")
	}

	// Simulate a delay of 2 seconds
	time.Sleep(2 * time.Second)
	log.Println("Event delivered to destination 1 ")

	return nil
}

func DeliverToDestination2(event Event) error {
	// Simulate a delay of 1 second
	time.Sleep(1 * time.Second)
	log.Println("Event delivered to destination 2 ")

	return nil
}

func DeliverToDestination3(event Event) error {
	// Simulate a delay of 3 seconds
	time.Sleep(3 * time.Second)

	log.Println("Event delivered to destination 3 ")

	return nil
}
