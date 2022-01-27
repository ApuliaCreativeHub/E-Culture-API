package models

import "time"

type Place struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	PhotoPath     string `json:"photoPath"`
	Address       string `json:"address"`
	Description   string `json:"description"`
	Lat           string `json:"lat"`
	Long          string `json:"long"`
	UserID        uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
	User          User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	NormalSizeImg string `gorm:"-" json:"normalSizeImg"`
	Thumbnail     string `gorm:"-"`
}

func (Place) TableName() string {
	return "place"
}

func (p *Place) Create() error {
	tx := Db.Create(p)
	return tx.Error
}

func (p *Place) Update() error {
	tx := Db.Model(p).Where(p.ID).Updates(p)
	return tx.Error
}

func (p *Place) Delete() error {
	tx := Db.Delete(p)
	return tx.Error
}

func (p *Place) Read() error {
	tx := Db.Where("id=?", p.ID).Find(p)
	return tx.Error
}

func (p *Place) ReadByUserId() ([]Place, error) {
	var places []Place
	tx := Db.Where("user_id=?", p.UserID).Find(&places)
	return places, tx.Error
}

func (p *Place) ReadByAddress() error {
	tx := Db.Where("address=?", p.Address).Find(p)
	return tx.Error
}

func (p *Place) ReadAll() ([]Place, error) {
	var places []Place
	tx := Db.Find(&places)
	return places, tx.Error
}
