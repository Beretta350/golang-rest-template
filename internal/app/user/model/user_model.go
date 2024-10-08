package model

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type User struct {
	Id       string    `bson:"_id,omitempty" json:"id,omitempty" validate:"uuid"`
	Username string    `bson:"username" json:"username" validate:"required,min=3"`
	Password string    `bson:"password" json:"password,omitempty" validate:"required,min=8"`
	CreateAt time.Time `bson:"createAt" json:"createAt"`
	UpdateAt time.Time `bson:"updateAt" json:"updateAt"`
}

func NewUserModel(username, password string) *User {
	id := uuid.NewString()
	return &User{Id: id, Username: username, Password: password}
}

func (u User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

func (u User) ValidatePassword() error {
	validate := validator.New()
	return validate.StructPartial(u, "Password")
}
