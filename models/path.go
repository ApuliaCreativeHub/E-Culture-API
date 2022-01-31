package models

import (
	"time"
)

type Path struct {
	ID        uint
	Name      string
	UserID    uint
	User      User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Objects   []Object `gorm:"many2many:is_present_in;"`
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
	tx := Db.Raw("SELECT DISTINCT path.* FROM path "+
		"INNER JOIN is_present_in AS ipi ON path.id=ipi.path_id "+
		"INNER JOIN object AS o ON ipi.object_id=o.id "+
		"INNER JOIN zone AS z ON o.zone_id=z.id "+
		"INNER JOIN place AS p ON z.place_id=p.id "+
		"WHERE p.id=?", placeId).Find(&paths)
	tx = Db.Preload("Objects").Preload("Objects.Zone", "place_id=?", placeId).Find(&paths)
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

func (p *Path) Update(objects []Object) error {
	tx := Db.Model(p).Updates(p)
	err := p.Delete()
	if err != nil {
		return err
	}

	for i, o := range objects {
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
