package entities

import (
	"time"
)

type Swipe struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	ProfileID uint
	Action    string
	CreatedAt time.Time
}
