package models

import (
	"E-Culture-API/utils"
)

type User struct {
	ID         uint
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	IsACurator bool   `json:"isACurator"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

func (User) TableName() string {
	return "user"
}

func (u *User) Create() error {
	var err error
	u.Password, err = utils.CryptSHA1(u.Password)
	if err != nil {
		return err
	}
	tx := Db.Create(u)
	return tx.Error
}

func (u *User) Update() error {
	var err error
	u.Password, err = utils.CryptSHA1(u.Password)
	if err != nil {
		return err
	}
	tx := Db.Model(u).Updates(u)
	return tx.Error
}

func (u *User) Delete() error {
	tx := Db.Delete(u)
	return tx.Error
}

func (u *User) ReadById() error {
	tx := Db.Where("id=?", u.Email).Find(u)
	return tx.Error
}

func (u *User) ReadByEmail() error {
	tx := Db.Where("email=?", u.Email).Find(u)
	return tx.Error
}
