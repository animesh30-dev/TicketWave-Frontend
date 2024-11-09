package handlers

import (
	"context"
	"time"

	"github.com/animesh_30/TicketWave/models"
	"github.com/gofiber/fiber/v2"
)

type EventHandler struct {
repository models.EventRepostitory
}

func (h *EventHandler) GetMany(ctx *fiber.Ctx) error{
	context,cancel := context.WithTimeout(context.Background(),time.Duration(5*time.Second))
	defer cancel()

	events,err := h.repository.GetMany(context)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status" : "fail",
			"message" : err.Error(),
				})
	}
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status" : "success",
		"message" : "",
		"data" : events, 
	})
}
func (h *EventHandler) CreateOne(ctx *fiber.Ctx) error{
	return nil
}
func (h *EventHandler) GetOne(ctx *fiber.Ctx) error{
	return nil
}

func NewEventHandler(router fiber.Router, repostitory models.EventRepostitory){
	handler := &EventHandler{
		repository: repostitory,
	}

	router.Get("/",handler.GetMany)
	router.Post("/",handler.CreateOne)
	router.Get("/eventId",handler.GetOne)
}