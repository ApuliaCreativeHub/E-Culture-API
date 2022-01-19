package models

type IsPresentIn struct {
	ObjectID uint `gorm:"primaryKey"`
	PathID   uint `gorm:"primaryKey"`
	Order    uint `gorm:"primaryKey"`
}

func (IsPresentIn) TableName() string {
	return "is_present_in"
}
