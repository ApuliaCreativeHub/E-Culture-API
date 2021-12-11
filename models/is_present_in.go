package models

type IsPresentIn struct {
	ObjectId uint `gorm:"primaryKey"`
	PathId   uint `gorm:"primaryKey"`
	Order    uint `gorm:"primaryKey"`
}
