package models

type User struct {
	ID         uint
	Name       string
	Surname    string
	IsACurator bool
	Email      string
	Password   string
}

func (User) TableName() string {
	return "user"
}

func (u *User) Create() error {
	tx := Db.Create(u)
	return tx.Error
}

func (u *User) Update() error {
	tx := Db.Model(u).Updates(u)
	return tx.Error
}

func (u *User) Delete() error {
	tx := Db.Delete(u)
	return tx.Error
}

func (u *User) ReadByEmail() error {
	tx := Db.Where("email=?", u.Email).Find(u)
	return tx.Error
}
