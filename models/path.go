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
	Place     Place     `gorm:"-" json:"place"`
}

func (Path) TableName() string {
	return "path"
}

func (p *Path) Create() error {
	tx := Db.Create(p)
	return tx.Error
}

func (p *Path) AddObjectToPath(objectId uint) error {
	tx := Db.Create(&IsPresentIn{PathID: p.ID, ObjectID: objectId})
	return tx.Error
}

func (p *Path) ReadPathsByPlaceId(placeId uint) ([]Path, error) {
	var paths []Path
	tx := Db.Raw("SELECT p.* FROM path p "+
		"INNER JOIN is_present_in ipi ON p.id=ipi.path_id "+
		"INNER JOIN object o ON ipi.object_id=o.id "+
		"INNER JOIN zone z ON o.zone_id=z.id "+
		"WHERE z.place_id=? "+
		"GROUP BY p.id", placeId).Find(&paths)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var err error
	paths, err = fetchPathObjects(paths)
	if err != nil {
		return nil, err
	}

	return paths, tx.Error
}

func (p *Path) ReadCuratorPathsByPlaceId(placeId uint) ([]Path, error) {
	var paths []Path
	tx := Db.Preload("Objects").Joins("INNER JOIN user u ON path.user_id=u.id").Joins("INNER JOIN place p ON u.id=p.user_id AND p.id=?", placeId).Find(&paths)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var err error
	paths, err = fetchPathObjects(paths)
	if err != nil {
		return nil, err
	}

	return paths, nil
}

func (p *Path) ReadByUserId(userId uint) ([]Path, error) {
	var paths []Path
	tx := Db.Where("user_id=?", userId).Find(&paths)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var err error
	paths, err = fetchPathObjects(paths)
	if err != nil {
		return nil, err
	}

	for i := range paths {
		err := paths[i].Place.ReadByPathId(paths[i].ID)
		if err != nil {
			return nil, err
		}
	}

	return paths, nil
}

func fetchPathObjects(paths []Path) ([]Path, error) {
	for i := range paths {
		tx := Db.Raw("SELECT * FROM object o "+
			"INNER JOIN is_present_in ipi ON o.id=ipi.object_id "+
			"WHERE ipi.path_id=? "+
			"ORDER BY ipi.order", paths[i].ID).Preload("Zone").Find(&paths[i].Objects)
		if tx.Error != nil {
			return nil, tx.Error
		}
	}

	return paths, nil
}

func (p *Path) ReadByPathId() error {
	tx := Db.Where("id=?", p.ID).Find(p)
	return tx.Error
}

func (p *Path) Update() error {
	tx := Db.Updates(p)
	if tx.Error != nil {
		return tx.Error
	}
	err := p.DeleteAllPathObjects()
	if err != nil {
		return err
	}

	for _, o := range p.Objects {
		err = p.AddObjectToPath(o.ID)
		if err != nil {
			return err
		}
	}
	return err
}

func (p *Path) Delete() error {
	tx := Db.Where("id=?", p.ID).Delete(p)
	return tx.Error
}

func (p *Path) DeleteAllPathObjects() error {
	tx := Db.Where("path_id=?", p.ID).Delete(&IsPresentIn{})
	return tx.Error
}
