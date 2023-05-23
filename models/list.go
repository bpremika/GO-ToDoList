package models

import (
	"gorm.io/gorm"
)

type List struct {
	gorm.Model
	Title    string
	Position uint8
	// Priority PriorityType `gorm:"type:enum('high', 'medium', 'low')"`
	Tasks []Task `gorm:"foreignKey:ListID"`
}
