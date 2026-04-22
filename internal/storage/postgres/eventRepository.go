package postgres

import (
	"context"

	"github.com/Aaron-GMM/DockOps/internal/config/logger"
	"github.com/Aaron-GMM/DockOps/internal/core"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) Save(ctx context.Context, event core.Event) error {
	log = logger.NewLogger("EventRepository")
	log.Info("Saving Event ")
	dbModel := EventModel{
		ID:         event.ID,
		ResourceID: event.ResourceID,
		Type:       string(event.Type),
		Payload:    event.Payload,
		CreatedAt:  event.CreatedAt,
	}
	err := r.db.Create(&dbModel).Error
	if err != nil {
		log.Error("Create Event error: %v", err.Error())
		return err
	}

	return nil
}
