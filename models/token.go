package models

import "time"

type Token struct {
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"createdAt"`
	UUID      string    `json:"uuid"`
	UserID    uint      `json:"-"`
	User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:Cascade;" json:"-"`
}

func (Token) TableName() string {
	return "token"
}

func (t *Token) Create() error {
	tx := Db.Create(t)
	return tx.Error
}

func (t *Token) Update() error {
	tx := Db.Model(t).Updates(t)
	return tx.Error
}

func (t *Token) DeleteByUUID() error {
	tx := Db.Where("uuid=?", t.UUID).Delete(t)
	return tx.Error
}

func (t *Token) DeleteByToken() error {
	tx := Db.Where("token=?", t.Token).Delete(t)
	return tx.Error
}

func (t *Token) ReadByToken() (int64, error) {
	tx := Db.Where("token=?", t.Token).Find(t)
	return tx.RowsAffected, tx.Error
}
