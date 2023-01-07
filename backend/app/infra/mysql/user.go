package mysql

import (
	"context"
	"database/sql"

	"github.com/RIckyBan/webapp-from-scratch/backend/app/entity"
	"github.com/RIckyBan/webapp-from-scratch/backend/app/repository"
	"github.com/RIckyBan/webapp-from-scratch/backend/db/models"
	"github.com/oklog/ulid"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (r userRepository) GetAllUsers(ctx context.Context) ([]*entity.User, error) {
	users := []*entity.User{}

	us, err := models.Users().All(ctx, r.db)
	if err != nil {
		return users, err
	}

	for _, u := range us {
		userID, err := ulid.Parse(u.ID)
		if err != nil {
			panic(err)
		}

		users = append(users, entity.NewUser(
			userID,
			u.Name,
			u.Email,
			u.Address,
		))
	}

	return users, nil
}
