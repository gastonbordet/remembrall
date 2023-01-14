package events

import "net/http"

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
	w.Write([]byte("Hola mundo"))
}
