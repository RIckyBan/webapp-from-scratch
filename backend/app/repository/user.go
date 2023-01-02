package repository

import (
	"context"

	"github.com/RIckyBan/webapp-from-scratch/backend/app/entity"
)

type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]*entity.User, error)
}
