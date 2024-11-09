package repositories

import (
	"context"
	"time"

	"github.com/animesh_30/TicketWave/models"
)

type EventRepostitory struct {
	db any
}

func(r *EventRepostitory) GetMany(ctx context.Context) ([]*models.Event,error){
	events := []*models.Event{}

	events = append(events , &models.Event{
		ID : "202151415156",
		Name: "My Fav Horse",
		Location : "Heart",
		Date: 	time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	return events,nil
}
func(r *EventRepostitory) GetOne(ctx context.Context, eventId string) (*models.Event,error){
	return nil,nil
}
func(r *EventRepostitory) CreateOne(ctx context.Context,event models.Event) (*models.Event,error){
	return nil,nil
}

func NewEventRepository(db any) models.EventRepostitory{
	return & EventRepostitory{
		db:db,
	} 
}