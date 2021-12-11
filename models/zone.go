package models

type Zone struct {
	ID          uint
	Name        string
	Description string
	PlaceId     uint
	Place       Place `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Zone) TableName() string {
	return "zone"
}

func CreateZone(name, description string, placeID uint) error {
	zone := Zone{Name: name, Description: description, PlaceId: placeID}
	tx := Db.Create(&zone)
	return tx.Error
}
