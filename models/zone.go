package models

type Zone struct {
	ID          uint
	Name        string
	Description string
	PlaceId     uint
	Objects     []Object
}
