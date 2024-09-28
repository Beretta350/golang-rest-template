package model

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type User struct {
	Id       string    `bson:"_id,omitempty" json:"id" validate:"uuid"`
	Username string    `bson:"username" json:"username" validate:"required,min=3"`
	Password string    `bson:"password" json:"password" validate:"required,min=8"`
	CreateAt time.Time `bson:"create_at" json:"create_at"`
	UpdateAt time.Time `bson:"update_at" json:"update_at"`
}

func NewUserModel(username, password string) *User {
	id := uuid.NewString()
	return &User{Id: id, Username: username, Password: password}
}

func (u User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
