package entity

import "github.com/oklog/ulid"

type User struct {
	ID      ulid.ULID
	Name    string
	Email   string
	Address string
}

func NewUser(id ulid.ULID) *User {
	return &User{
		ID: id,
	}
}
