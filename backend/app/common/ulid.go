package common

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

func ParseULID(s string) (ulid.ULID, error) {
	str, err := ulid.Parse(s)
	if err != nil {
		return ulid.ULID{}, err
	}
	return str, nil
}

func GenerateULID() ulid.ULID {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Now(), entropy)
}
