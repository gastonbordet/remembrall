package events

import "time"

type IEventsService interface {
	GetAll() ([]Event, error)
}

type EventsService struct{}

func BuildEventsService() *EventsService {
	service := &EventsService{}
	return service
}

func (service *EventsService) GetAll() ([]Event, error) {
	var events []Event

	event := &Event{
		ID:          1,
		Title:       "Test event",
		Date:        time.Now().Format("2006-01-02"),
		Status:      true,
		Description: "Event for testing",
	}

	events = append(events, *event)

	return events, nil
}
