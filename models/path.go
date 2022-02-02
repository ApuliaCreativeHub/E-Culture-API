package models

import (
	"time"
)

type Path struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	UserID    uint      `json:"userId"`
	User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Objects   []Object  `gorm:"many2many:is_present_in;" json:"objects"`
}

func (Path) TableName() string {
	return "path"
}

func (p *Path) Create() error {
	tx := Db.Create(p)
	return tx.Error
}

func (p *Path) AddObjectToPath(objectId, order uint) error {
	tx := Db.Create(&IsPresentIn{PathID: p.ID, ObjectID: objectId, Order: order})
	return tx.Error
}

func (p *Path) ReadPathsByPlaceId(placeId uint) ([]Path, error) {
	var paths []Path
	tx := Db.Preload("Objects").Preload("Objects.Zone", "place_id=?", placeId).Find(&paths)
	return paths, tx.Error
}

func (p *Path) ReadCuratorPathsByPlaceId(placeId uint) ([]Path, error) {
	var paths []Path
	tx := Db.Preload("Objects").Joins("INNER JOIN user u ON path.user_id=u.id").Joins("INNER JOIN place p ON u.id=p.user_id AND p.id=?", placeId).Find(&paths)
	return paths, tx.Error
}

func (p *Path) ReadByUserId(userId uint) ([]Path, error) {
	var paths []Path
	tx := Db.Where("user_id=?", userId).Preload("Objects").Find(&paths)
	return paths, tx.Error
}

func (p *Path) ReadByPathId() error {
	tx := Db.Where("id=?", p.ID).Find(p)
	return tx.Error
}

func (p *Path) Update() error {
	tx := Db.Model(p).Updates(p)
	err := p.Delete()
	if err != nil {
		return err
	}

	for i, o := range p.Objects {
		err = p.AddObjectToPath(o.ID, uint(i))
		if err != nil {
			return err
		}
	}
	return tx.Error
}

func (p *Path) Delete() error {
	tx := Db.Where("path_id=?", p.ID).Delete(&IsPresentIn{})
	if tx.Error != nil {
		return tx.Error
	}
	tx = Db.Where("id=?", p.ID).Delete(p)
	return tx.Error
}
