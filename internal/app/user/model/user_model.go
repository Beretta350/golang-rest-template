package model

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type User struct {
	Id        string    `bson:"_id,omitempty" db:"id" json:"id,omitempty" validate:"uuid"`
	Username  string    `bson:"username" db:"username" json:"username" validate:"required,min=3"`
	Password  string    `bson:"password" db:"password" json:"password,omitempty" validate:"required,min=8"`
	CreatedAt time.Time `bson:"createdAt" db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" db:"updated_at" json:"updatedAt"`
}

func NewUserModel(username, password string) *User {
	id := uuid.NewString()
	return &User{Id: id, Username: username, Password: password}
}

func (u User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

func (u User) ValidateUsername() error {
	validate := validator.New()
	return validate.StructPartial(u, "Username")
}

func (u User) ValidatePassword() error {
	validate := validator.New()
	return validate.StructPartial(u, "Password")
}
