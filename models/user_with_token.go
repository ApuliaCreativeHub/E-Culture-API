package models

type UserWithToken struct {
	User  User  `json:"user"`
	Token Token `json:"token"`
}
