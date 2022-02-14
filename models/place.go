package models

import "time"

type Place struct {
	ID            uint      `json:"id"`
	Name          string    `json:"name"`
	PhotoPath     string    `json:"photoPath"`
	Address       string    `json:"address"`
	Description   string    `json:"description"`
	Lat           string    `json:"lat"`
	Long          string    `json:"long"`
	UserID        uint      `json:"userId"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
	FileName      string    `json:"-"`
	User          User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	NormalSizeImg string    `gorm:"-" json:"normalSizeImg"`
	Thumbnail     string    `gorm:"-" json:"-"`
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

func (p *Place) ReadByPathId(pathId uint) error {
	tx := Db.Raw("SELECT pl.* FROM place pl "+
		"JOIN `zone` z ON z.place_id  =pl.id "+
		"JOIN `object` o ON o.zone_id =z.id "+
		"JOIN is_present_in ipi ON ipi.object_id =o.id "+
		"JOIN `path` p ON p.id =ipi.path_id "+
		"WHERE p.id = ? "+
		"GROUP BY p.id", pathId).Find(p)
	return tx.Error
}
