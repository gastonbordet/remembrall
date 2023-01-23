package events

import (
	"context"
)

type IEventsService interface {
	GetAll(ctx context.Context) ([]Event, error)
	GetByEventID(ctx context.Context, eventID uint) (*Event, error)
}

type EventsService struct {
	Repository IEventsRepository
}

func BuildEventsService(eventsRepository IEventsRepository) *EventsService {
	service := &EventsService{
		Repository: eventsRepository,
	}
	return service
}

func (service *EventsService) GetAll(ctx context.Context) ([]Event, error) {
	events, err := service.Repository.GetAll(ctx)

	return events, err
}

func (service *EventsService) GetByEventID(ctx context.Context, eventID uint) (*Event, error) {
	event, err := service.Repository.GetByEventID(ctx, eventID)

	return event, err
}
