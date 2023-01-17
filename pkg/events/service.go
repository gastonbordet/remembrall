package events

import "time"

type EventsService struct{}

func BuildEventsService() *EventsService {
	service := &EventsService{}
	return service
}

func (service *EventsService) GetAll() ([]Event, error) {
	events := []Event{
		{
			ID:          1,
			Title:       "Test event 1",
			Date:        time.Now().Format("2006-01-02"),
			Status:      true,
			Description: "Event for testing",
		},
		{
			ID:          2,
			Title:       "Test event 2",
			Date:        time.Now().Format("2006-01-02"),
			Status:      true,
			Description: "Event for testing",
		},
	}

	return events, nil
}

func (service *EventsService) GetByEventID(eventID string) (*Event, error) {
	event := &Event{
		ID:          1,
		Title:       "Test event 1",
		Date:        time.Now().Format("2006-01-02"),
		Status:      true,
		Description: "Event for testing",
	}

	return event, nil
}
