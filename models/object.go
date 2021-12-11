package models

import (
	"time"
)

type Object struct {
	ID          uint
	Name        string
	Description string
	Age         string
	PhotoPath   string
	ZoneId      uint
	Zone        Zone `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Paths       []Path `gorm:"many2many:is_present_in;"`
}

func (Object) TableName() string {
	return "object"
}

func CreateObject(name, description, age, photoPath string, zone uint) error {
	object := Object{Name: name, Description: description, Age: age, PhotoPath: photoPath, ZoneId: zone}
	Db.Create(&object)
	return nil
}
