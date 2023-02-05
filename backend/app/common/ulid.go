package common

import "github.com/oklog/ulid"

func ParseULID(s string) ulid.ULID {
	str, err := ulid.Parse(s)
	if err != nil {
		panic(err)
	}
	return str
}
