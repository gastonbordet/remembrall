package main

import (
	"fmt"
	"net/http"

	"github.com/gastonbordet/remembrall/cmd/util"

	"github.com/gastonbordet/remembrall/pkg/events"
	"github.com/gastonbordet/remembrall/pkg/http/rest"
	"github.com/gastonbordet/remembrall/pkg/storage/database"
)

type ServerConfig struct {
	port int
}

func startServer(config *ServerConfig) {
	envConfig, errConfig := util.LoadConfig(".", "config", "env")
	if errConfig != nil {
		fmt.Println("Error reading env configuration, err: ", errConfig)
	}
	fmt.Println("env: ", envConfig)
	sqlClient := database.NewSQLClient()
	sqlConn, sqlErr := sqlClient.OpenConnection(database.GenerateMysqlURIConnection(envConfig))

	if sqlErr != nil {
		return
	}

	events_repository := events.BuildEventsRepository(sqlConn)
	events_service := events.BuildEventsService(events_repository)
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
