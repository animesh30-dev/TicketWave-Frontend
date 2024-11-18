package repositories

import (
	"context"

	"github.com/animesh_30/TicketWave/models"
	"gorm.io/gorm"
)

type EventRepostitory struct {
	db *gorm.DB
}

func(r *EventRepostitory) GetMany(ctx context.Context) ([]*models.Event,error){
	events := []*models.Event{}

	res := r.db.Model(&models.Event{}).Find(&events)

	if res.Error != nil {
		return nil,res.Error
	}
	return events,nil
}
func(r *EventRepostitory) GetOne(ctx context.Context, eventId string) (*models.Event,error){
	return nil,nil
}
func(r *EventRepostitory) CreateOne(ctx context.Context,event models.Event) (*models.Event,error){
	return nil,nil
}

func NewEventRepository(db *gorm.DB) models.EventRepostitory{
	return & EventRepostitory{
		db:db,
	} 
}