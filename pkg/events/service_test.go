package events_test

import (
	"context"
	"testing"

	"github.com/gastonbordet/remembrall/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockEventsRepository struct {
	mock.Mock
}

func (mock *MockEventsRepository) GetAll(ctx context.Context) ([]events.Event, error) {
	args := mock.Called()
	mock_events := args.Get(0).([]events.Event)

	return mock_events, args.Error(1)
}

func (mock *MockEventsRepository) GetByEventID(ctx context.Context, eventID uint) (*events.Event, error) {
	args := mock.Called()
	mock_event := args.Get(0).(*events.Event)

	return mock_event, args.Error(1)
}

func TestGetAll(t *testing.T) {
	// Setup
	events_response := []events.Event{{
		ID:          1,
		Title:       "test",
		Description: "test",
		Date:        "",
		Status:      true,
	}}
	eventsRepository := &MockEventsRepository{}
	eventsRepository.On("GetAll").Return(events_response, nil)
	eventsService := events.BuildEventsService(eventsRepository)

	// Act
	events, err := eventsService.GetAll(context.Background())

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, events)
	assert.Len(t, events, 1)
	assert.Equal(t, "test", events[0].Title)
}

func TestGetByEventID(t *testing.T) {
	// Setup
	eventID := uint(1)
	eventsRepository := &MockEventsRepository{}
	eventsService := events.BuildEventsService(eventsRepository)
	eventsRepository.On("GetByEventID").Return(&events.Event{
		Title: "test",
	}, nil)

	// Act
	event, err := eventsService.GetByEventID(context.Background(), eventID)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, event)
	assert.Equal(t, "test", event.Title)
}
