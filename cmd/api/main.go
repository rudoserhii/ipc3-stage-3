package main

import (
	"log"

	"github.com/obiMadu/ipc3-stage-3/internal/db"
	"github.com/obiMadu/ipc3-stage-3/internal/handlers"
	"github.com/obiMadu/ipc3-stage-3/internal/interfaces"
	"github.com/obiMadu/ipc3-stage-3/internal/models"
)

const webPort = ":8081"

type Config struct {
	Handlers interfaces.Handlers
}

// @title			Mini Shop API
// @version			1.0
// @description		This is a small Webstore API app.
// @host			ips3.obi.ninja
// @BasePath		/api/v1/
func main() {
	// Init DB.
	db := db.InitDB()
	// Make Models.
	models := models.NewModels(db)
	// Make Handlers.
	handlers := handlers.NewHandlers(models)

	app := Config{
		Handlers: handlers,
	}

	err := app.routes().Run(webPort)
	if err != nil {
		log.Fatalf("Failed to start http server: %v", err)
	}
}
