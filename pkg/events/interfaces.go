package events

import (
	"context"
	"net/http"
)

type IEventsHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByEventID(w http.ResponseWriter, r *http.Request)
}

type IEventsService interface {
	GetAll(ctx context.Context) ([]Event, error)
	GetByEventID(ctx context.Context, eventID string) (*Event, error)
}
