package main

import (
	"fmt"
	"net/http"

	"github.com/gastonbordet/remembrall/pkg/events"
	"github.com/gastonbordet/remembrall/pkg/http/router"
)

type ServerConfig struct {
	port int
}

func startServer(config *ServerConfig) {
	events_handlers := events.BuildEventsHandler()

	handler := router.InitRouter(events_handlers)
	fmt.Println(fmt.Sprintf("Start app on port: %d", config.port))
	http.ListenAndServe(fmt.Sprintf(":%d", config.port), handler)
}

func main() {
	config := &ServerConfig{
		port: 8000,
	}

	startServer(config)
}
