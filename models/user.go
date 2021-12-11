package models

type User struct {
	ID         uint
	Name       string
	Surname    string
	IsACurator bool
	Email      string
	Password   string
	Places     []Place
	Paths      []Path
}
