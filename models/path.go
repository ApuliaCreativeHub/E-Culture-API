package models

import (
	"time"
)

type Path struct {
	ID          uint
	Name        string
	Description string
	UserID      uint
	User        User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Objects     []Object `gorm:"many2many:is_present_in;"`
}

func (Path) TableName() string {
	return "path"
}

func CreatePath(name, description string, userId uint) error {
	err := Db.SetupJoinTable(&Path{}, "Objects", &IsPresentIn{})
	if err != nil {
		return err
	}
	path := Path{Name: name, Description: description, UserID: userId}
	tx := Db.Create(&path)
	return tx.Error
}

func AddObjectToPath(pathId, objectId, order uint) error {
	tx := Db.Create(&IsPresentIn{PathID: pathId, ObjectID: objectId, Order: order})
	return tx.Error
}
