package models

type IsPresentIn struct {
	ObjectID uint `gorm:"primaryKey"`
	PathID   uint `gorm:"primaryKey"`
	Order    uint
}

func (IsPresentIn) TableName() string {
	return "is_present_in"
}
