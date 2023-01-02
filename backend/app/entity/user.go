package entity

import "github.com/oklog/ulid"

type User struct {
	ID      ulid.ULID
	Name    string
	Email   string
	Address string
}

func NewUser(id ulid.ULID, name, email, address string) *User {
	return &User{
		ID:      id,
		Name:    name,
		Email:   email,
		Address: address,
	}
}
