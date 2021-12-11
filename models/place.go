package models

type Place struct {
	ID          uint
	Name        string
	PhotoPath   string
	Address     string
	Description string
	UserId      uint
	User        User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Place) TableName() string {
	return "place"
}

func CreatePlace(name, photoPath, address, description string, userId uint) error {
	place := Place{Name: name, PhotoPath: photoPath, Address: address, Description: description, UserId: userId}
	tx := Db.Create(&place)
	return tx.Error
}
