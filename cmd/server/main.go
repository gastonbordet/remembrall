package main

import (
	"fmt"
	"net/http"

	"github.com/gastonbordet/remembrall/pkg/util/config"
	"github.com/gastonbordet/remembrall/pkg/util/logs"

	"github.com/gastonbordet/remembrall/pkg/events"
	"github.com/gastonbordet/remembrall/pkg/http/rest"
	"github.com/gastonbordet/remembrall/pkg/storage/database"
)

type ServerConfig struct {
	port int
}

func startServer(serverConfig *ServerConfig) {
	logs.InitZapLogger()
	envConfig, errConfig := config.LoadConfig(".", "config", "env")

	if errConfig != nil {
		logs.Logger.Error(fmt.Sprintf("Error reading env configuration, err: %v", errConfig))
	}

	sqlClient := database.NewSQLClient()
	sqlConn, sqlErr := sqlClient.OpenConnection(database.GenerateMysqlURIConnection(envConfig))

	if sqlErr != nil {
		return
	}

	events_repository := events.BuildEventsRepository(sqlConn)
	events_service := events.BuildEventsService(events_repository)
	events_handlers := events.BuildEventsHandler(events_service)

	handler := rest.InitRouter(events_handlers)
	logs.Logger.Info(fmt.Sprintf("Start app on port: %d", serverConfig.port))
	http.ListenAndServe(fmt.Sprintf(":%d", serverConfig.port), handler)
}

func main() {
	config := &ServerConfig{
		port: 8000,
	}

	startServer(config)
}
