package delivery

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeliverToDestination1_Success(t *testing.T) {
	event := Event{
		UserID:  "456",
		Payload: "payload",
	}

	err := DeliverToDestination1(event)
	assert.NoError(t, err, "No errors occured when delivering to destination 1")
}

func TestDeliverToDestination1_Failure(t *testing.T) {
	event := Event{
		UserID:  "123",
		Payload: "payload",
	}

	err := DeliverToDestination1(event)
	assert.Error(t, err, "Error occured when delivering to destination 1")
}

func TestDeliverToDestination2(t *testing.T) {
	event := Event{
		UserID:  "456",
		Payload: "payload",
	}

	err := DeliverToDestination2(event)
	assert.NoError(t, err, "No errors occured when delivering to destination 2")
}

func TestDeliverToDestination3(t *testing.T) {
	event := Event{
		UserID:  "456",
		Payload: "payload",
	}

	err := DeliverToDestination3(event)
	assert.NoError(t, err, "No errors occured when delivering to destination 3")
}

func TestDeliverToDestinations_Success(t *testing.T) {
	event := Event{
		UserID:  "456",
		Payload: "payload",
	}

	err := DeliverToDestinations(event)
	assert.NoError(t, err, "No errors occured when delivering to destinations")
}

func TestDeliverToDestinations_Failure(t *testing.T) {
	event := Event{
		UserID:  "123",
		Payload: "payload",
	}

	err := DeliverToDestinations(event)
	assert.Error(t, err, "Error occured when delivering to destinations")
}
