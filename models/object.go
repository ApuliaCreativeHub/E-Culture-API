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
	ZoneID      uint
	Zone        Zone `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Paths       []Path `gorm:"many2many:is_present_in;"`
}

func (Object) TableName() string {
	return "object"
}

func (o *Object) Create() error {
	tx := Db.Create(o)
	return tx.Error
}

func (o *Object) Update() error {
	tx := Db.Model(o).Updates(o)
	return tx.Error
}

func (o *Object) Delete() error {
	tx := Db.Delete(o)
	return tx.Error
}

func (o *Object) ReadByZoneId() ([]Object, error) {
	var objects []Object
	tx := Db.Where("zone_id=?", o.ZoneID).Find(&objects)
	return objects, tx.Error
}
