package router

import (
	"net/http"

	"github.com/gastonbordet/remembrall/pkg/events"
	"github.com/go-chi/chi/v5"
)

func InitRouter(events_handler events.IEventsHandler) http.Handler {
	r := chi.NewRouter()
	r.Route("/events", func(r chi.Router) {
		r.Get("/", events_handler.GetEvents)
	})

	return r
}
