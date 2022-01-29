package models

import (
	"time"
)

type Object struct {
	ID            uint      `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	PhotoPath     string    `json:"photoPath"`
	ZoneID        uint      `json:"zone_id"`
	Zone          Zone      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
	FileName      string    `json:"-"`
	NormalSizeImg string    `gorm:"-" json:"normalSizeImg"`
	Paths         []Path    `gorm:"many2many:is_present_in;" json:"paths"`
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

func (o *Object) ReadAll() error {
	tx := Db.Preload("Zone").Preload("Zone.Place").Preload("Zone.Place.User").Find(&o)
	return tx.Error
}

func (o *Object) ReadByZoneId() ([]Object, error) {
	var objects []Object
	tx := Db.Where("zone_id=?", o.ZoneID).Find(&objects)
	return objects, tx.Error
}
