package models

import "time"

type Path struct {
	ID          uint
	Name        string
	Description string
	UserId      uint
	User        User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Objects     Object[] `gorm:"many2many:is_present_in;"`
}

func (Path) TableName() string {
	return "path"
}

func CreatePath(name, description string, userId uint) error {
	path := Path{Name: name, Description: description, UserId: userId}
	Db.Create(&path)
	return nil
}

func AddObjectToPath(pathId, objectId, order uint) {
	//TODO: association
}
