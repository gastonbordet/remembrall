package events

import (
	"encoding/json"
	"log"
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
	events, eventsErr := handler.service.GetAll()

	if eventsErr != nil {
		log.Fatalf("Error happened at service.GetAll. Error: %s", eventsErr)
	}

	response, err := json.Marshal(events)

	if err != nil {
		log.Fatalf("Error happened in Events JSON marshal. Error: %s", err)
	}

	w.Write(response)
}

func (handler *EventsHandler) GetByEventID(w http.ResponseWriter, r *http.Request) {
	// TODO get eventID from url query param
	eventID := ""
	event, eventErr := handler.service.GetByEventID(eventID)

	if eventErr != nil {
		log.Fatalf("Error happened at service.GetAll. Error: %s", eventErr)
	}

	response, err := json.Marshal(event)

	if err != nil {
		log.Fatalf("Error happened in Events JSON marshal. Error: %s", err)
	}

	w.Write(response)
}
