package events

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type IEventsHandler interface {
	GetEvents(w http.ResponseWriter, r *http.Request)
}

type EventsHandler struct {
}

func BuildEventsHandler() *EventsHandler {
	eventsHandler := &EventsHandler{}
	return eventsHandler
}

func (handler *EventsHandler) GetEvents(w http.ResponseWriter, r *http.Request) {
	event := &Event{
		id:          1,
		title:       "Test event",
		date:        time.Now().Format("2006-01-02"),
		status:      true,
		description: "Event for testing",
	}

	response, marshalErr := json.Marshal(event)

	if marshalErr != nil {
		fmt.Println("Marshal error")
	}

	w.Write(response)
}
