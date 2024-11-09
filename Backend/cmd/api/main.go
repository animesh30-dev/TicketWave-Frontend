package main

import (
	"github.com/animesh_30/TicketWave/handlers"
	"github.com/animesh_30/TicketWave/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "TicketWave",
		ServerHeader: " Fiber",
	})

	//repositories

	eventRepostitory := repositories.NewEventRepository(nil)

	server:= app.Group("/api")

	handlers.NewEventHandler(server.Group("/event"), eventRepostitory)

	app.Listen(":3000")
}