package main

import (
	"fmt"

	"github.com/animesh_30/TicketWave/config"
	"github.com/animesh_30/TicketWave/db"
	"github.com/animesh_30/TicketWave/handlers"
	"github.com/animesh_30/TicketWave/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {
	envConfig := config.NewEnvConfig()
	db:= db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "TicketWave",
		ServerHeader: "Fiber",
	})

	//repositories

	eventRepostitory := repositories.NewEventRepository(db)

	server:= app.Group("/api")

	handlers.NewEventHandler(server.Group("/event"), eventRepostitory)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}