package events_test

import (
	"context"
	"testing"

	"github.com/gastonbordet/remembrall/pkg/events"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	// Setup
	eventsService := events.BuildEventsService()

	// Act
	events, err := eventsService.GetAll(context.Background())

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, events)
	assert.Len(t, events, 2)
}

func TestGetByEventID(t *testing.T) {
	// Setup
	eventID := "1"
	eventsService := events.BuildEventsService()

	// Act
	event, err := eventsService.GetByEventID(context.Background(), eventID)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, event)
}
