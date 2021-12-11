package models

import "time"

type Path struct {
	ID          uint
	Name        string
	Description string
	User        uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Objects     []Object `gorm:"many2many:is_present_in;"`
}
