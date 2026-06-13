package postgres

import (
	"context"

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
		log.ErrorF("Create Event error: %v", err.Error())
		return err
	}

	return nil
}

func (r *EventRepository) GetByResourceID(ctx context.Context, resourceID string) ([]core.Event, error) {
	var models []EventModel
	err := r.db.WithContext(ctx).Where("resource_id = ?", resourceID).Order("created_at asc").Find(&models).Error
	if err != nil {
		return nil, err
	}

	events := make([]core.Event, len(models))
	for i, m := range models {
		events[i] = core.Event{
			ID:         m.ID,
			ResourceID: m.ResourceID,
			Type:       core.EventType(m.Type),
			Payload:    m.Payload,
			CreatedAt:  m.CreatedAt,
		}
	}
	return events, nil
}
