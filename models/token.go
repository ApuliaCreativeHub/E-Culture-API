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

func (t *Token) Delete() error {
	tx := Db.Where("uuid=?", t.UUID).Delete(t)
	return tx.Error
}
