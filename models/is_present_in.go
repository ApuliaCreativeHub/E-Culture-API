package models

type IsPresentIn struct {
	ObjectId uint `gorm:"primaryKey"`
	PathId   uint `gorm:"primaryKey"`
	Order    uint `gorm:"primaryKey"`
}

func (IsPresentIn) TableName() string {
	return "is_present_in"
}

func AddObjectToPath(pathId, objectId, order uint) error {
	tx := Db.Create(&IsPresentIn{PathId: pathId, ObjectId: objectId, Order: order})
	return tx.Error
}
