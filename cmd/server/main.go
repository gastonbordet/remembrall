package main

import (
	"fmt"
	"net/http"

	"github.com/gastonbordet/remembrall/pkg/events"
	"github.com/gastonbordet/remembrall/pkg/http/rest"
)

type ServerConfig struct {
	port int
}

func startServer(config *ServerConfig) {
	events_service := events.BuildEventsService()
	events_handlers := events.BuildEventsHandler(events_service)

	handler := rest.InitRouter(events_handlers)
	fmt.Println(fmt.Sprintf("Start app on port: %d", config.port))
	http.ListenAndServe(fmt.Sprintf(":%d", config.port), handler)
}

func main() {
	config := &ServerConfig{
		port: 8000,
	}

	startServer(config)
}
