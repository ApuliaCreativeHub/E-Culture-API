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

func CreateUser(name, surname, email, password string, isACurator bool) error {
	user := User{Name: name, Surname: surname, IsACurator: isACurator, Email: email, Password: password}
	Db.Create(&user)
	return nil
}
