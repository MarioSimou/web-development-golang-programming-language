package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Collection interface {
	Name() string
	DbName() string
}

type Users struct {
	Docs []User
}

func (u Users) Name() string {
	return "users"
}

func (u Users) DbName() string {
	return "test"
}

type User struct {
	Id       *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username string              `json:"username,omitempty" bson:"username"`
	Email    string              `json:"email,omitempty" bson:"email"`
	Password string              `json:"password,omitempty" bson:"password"`
	Role     string              `json:"role,omitempty" bson:"role"`
}

func (u User) Name() string {
	return "user"
}
func (u User) CollectionName() string {
	return "users"
}

func (u User) Create() {
}
func (u User) Update() {
}
func (u User) Delete() {
}
