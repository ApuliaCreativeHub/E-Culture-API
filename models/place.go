package models

type Place struct {
	ID          uint
	Name        string
	PhotoPath   string
	Address     string
	Description string
	UserId      uint
	Zones       []Zone
}
