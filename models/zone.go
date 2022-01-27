package models

type Zone struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	PlaceID     uint   `json:"placeId"`
	Place       Place  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}

func (Zone) TableName() string {
	return "zone"
}

func (z *Zone) Create() error {
	tx := Db.Create(z)
	return tx.Error
}

func (z *Zone) Update() error {
	tx := Db.Model(z).Updates(z)
	return tx.Error
}

func (z *Zone) Delete() error {
	tx := Db.Delete(z)
	return tx.Error
}

func (z *Zone) ReadByPlaceId() ([]Zone, error) {
	var zones []Zone
	tx := Db.Where("place_id=?", z.PlaceID).Find(&zones)
	return zones, tx.Error
}

func (z *Zone) ReadByName() error {
	tx := Db.Where("name=?", z.Name).Find(z)
	return tx.Error
}

func (z *Zone) ReadAndPreloadPlace() error {
	tx := Db.Where("id=?", z.ID).Preload("Place").Find(z)
	return tx.Error
}
