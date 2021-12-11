package models

import "time"

type Object struct {
	ID          uint
	Name        string
	Description string
	Age         string
	PhotoPath   string
	ZoneId      uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Paths       []Path `gorm:"many2many:is_present_in;"`
}
