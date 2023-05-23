package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Description *string
	Duedate     *time.Time
	Position    uint8
	// Priority    PriorityType `gorm:"type:enum('high', 'medium', 'low')"`
	ListID int
}
