package models

import "time"

type Token struct {
	Token     string
	CreatedAt time.Time
	UUID      string
	UserID    uint
	User      User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:Cascade;"`
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
	tx := Db.Delete(t)
	return tx.Error
}
