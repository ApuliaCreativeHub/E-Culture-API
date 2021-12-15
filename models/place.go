package models

type Place struct {
	ID          uint
	Name        string
	PhotoPath   string
	Address     string
	Description string
	UserID      uint
	User        User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Place) TableName() string {
	return "place"
}

func (p *Place) Create() error {
	tx := Db.Create(p)
	return tx.Error
}

func (p *Place) Update() error {
	tx := Db.Model(p).Updates(p)
	return tx.Error
}

func (p *Place) Delete() error {
	tx := Db.Delete(p)
	return tx.Error
}

func (p *Place) ReadByUserId() ([]Place, error) {
	var places []Place
	tx := Db.Where("user_id=?", p.UserID).Find(&places)
	return places, tx.Error
}
