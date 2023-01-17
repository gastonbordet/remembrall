package events

import "net/http"

type IEventsHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByEventID(w http.ResponseWriter, r *http.Request)
}

type IEventsService interface {
	GetAll() ([]Event, error)
	GetByEventID(eventID string) (*Event, error)
}
