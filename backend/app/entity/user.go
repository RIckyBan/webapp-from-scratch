package entity

import (
	"time"

	"github.com/oklog/ulid"
)

type User struct {
	ID       ulid.ULID
	Name     string
	Email    string
	Birthday time.Time
	Admin    bool
}

func NewUser(id ulid.ULID, name, email string, birthday time.Time, admin bool) *User {
	return &User{
		ID:       id,
		Name:     name,
		Email:    email,
		Birthday: birthday,
		Admin:    admin,
	}
}
