package models

type User struct {
	Id       string
	Username string
	Email    string
	Password []byte
}
