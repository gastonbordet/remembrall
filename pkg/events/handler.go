package events

import (
	"encoding/json"
	"net/http"
)

type EventsHandler struct {
	service IEventsService
}

func BuildEventsHandler(eventsService IEventsService) *EventsHandler {
	eventsHandler := &EventsHandler{
		service: eventsService,
	}
	return eventsHandler
}

func (handler *EventsHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	events, err := handler.service.GetAll(ctx)

	if err != nil {
		webErr := handleError(ctx, "", err)
		w.WriteHeader(webErr.Status)
	}

	response, _ := json.Marshal(events)
	w.Write(response)
}

func (handler *EventsHandler) GetByEventID(w http.ResponseWriter, r *http.Request) {
	// TODO get eventID from url query param
	ctx := r.Context()
	eventID := "1"

	if eventID == "" {
		// 400 Bad request error
	}

	event, err := handler.service.GetByEventID(ctx, eventID)

	if err != nil {
		// 404 error
		handleError(ctx, eventID, err)
	}

	response, _ := json.Marshal(event)

	w.Write(response)
}
