package models

type Zone struct {
	ID          uint
	Name        string
	Description string
	PlaceID     uint
	Place       Place `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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
