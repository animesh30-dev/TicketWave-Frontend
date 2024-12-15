package main

import (
	"github.com/animesh_30/TicketWave/config"
	"github.com/animesh_30/TicketWave/db"
	"github.com/animesh_30/TicketWave/handlers"
	"github.com/animesh_30/TicketWave/middlewares"
	"github.com/animesh_30/TicketWave/repositories"
	"github.com/animesh_30/TicketWave/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "TicketWave",
		ServerHeader: "Fiber",
	})

	//repositories

	eventRepostitory := repositories.NewEventRepository(db)
	ticketRepository := repositories.NewTicketRepository(db)
	authRepository := repositories.NewAuthRepository(db)

	//service
	authService := services.NewAuthService(authRepository)

	//routing
	server := app.Group("/api")
	handlers.NewAuthHandler(server.Group("/Auth"), authService)

	//routes
	privateRoutes := server.Use(middlewares.AuthProteted(db))

	//handlers

	handlers.NewEventHandler(privateRoutes.Group("/event"), eventRepostitory)
	handlers.NewTicketHandler(privateRoutes.Group("/ticket"), ticketRepository)

	app.Listen(":" + envConfig.ServerPort)

}
