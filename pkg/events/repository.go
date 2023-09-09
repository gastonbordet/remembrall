package events

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type IEventsRepository interface {
	GetAll(ctx context.Context) ([]Event, error)
	GetByEventID(ctx context.Context, eventID uint) (*Event, error)
}

type EventsRepository struct {
	DB *gorm.DB
}

func BuildEventsRepository(dbConn *gorm.DB) *EventsRepository {
	dbConn.AutoMigrate(&Event{})
	return &EventsRepository{
		DB: dbConn,
	}
}

var (
	EventNotFoundError = errors.New("event not found")
	EventInternalError = errors.New("event internal error")
)

func (r *EventsRepository) GetAll(ctx context.Context) ([]Event, error) {
	var events []Event

	r.DB.Find(&events)

	return events, nil
}

func (r *EventsRepository) GetByEventID(ctx context.Context, eventID uint) (*Event, error) {
	event := &Event{
		ID: eventID,
	}
	r.DB.First(&event)

	if event.ValidateEvent() == false {
		return nil, EventNotFoundError
	}

	return event, nil
}

func (r *EventsRepository) CreateEvent(ctx context.Context, event *Event) (*Event, error) {
	r.DB.Create(event)

	return event, nil
}
