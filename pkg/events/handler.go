package events

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	web "github.com/gastonbordet/remembrall/pkg/web/error"
)

type IEventsHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByEventID(w http.ResponseWriter, r *http.Request)
}

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
		webErr := HandleError(ctx, "", err)
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
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	eventIDParsed, idParseErr := strconv.ParseUint(eventID, 10, 32)

	if idParseErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	event, err := handler.service.GetByEventID(ctx, uint(eventIDParsed))

	if err != nil {
		webErr := HandleError(ctx, eventID, err)
		w.WriteHeader(webErr.Status)
	}

	response, _ := json.Marshal(event)
	w.Write(response)
}

func HandleError(ctx context.Context, eventID string, err error) *web.WebError {
	code := http.StatusInternalServerError
	msg := fmt.Sprintf("Events Internal error: %s", err.Error())

	if errors.Is(err, EventNotFoundError) {
		code = http.StatusNotFound
		msg = fmt.Sprintf("Event %s not found.", eventID)
	}

	if errors.Is(err, EventInternalError) {
		code = http.StatusInternalServerError
		msg = fmt.Sprintf("Event %s generated internal error.", eventID)
	}

	return web.NewWebError(code, msg)
}
