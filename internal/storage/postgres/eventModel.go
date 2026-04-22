package postgres

import "time"

type EventModel struct {
	ID         string    `gorm:"type:uuid;primaryKey"`
	ResourceID string    `gorm:"type:varchar(255);index"`
	Type       string    `gorm:"type:varchar(50)"`
	Payload    []byte    `gorm:"type:jsonb"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}

func (EventModel) TableName() string {
	return "events"
}
