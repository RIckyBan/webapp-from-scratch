package usecase

import (
	"context"

	"github.com/RIckyBan/webapp-from-scratch/backend/app/entity"
	"github.com/RIckyBan/webapp-from-scratch/backend/app/repository"
)

type AdminService struct {
	userRepo repository.UserRepository
}

func NewAdminService(userRepo repository.UserRepository) *AdminService {
	return &AdminService{userRepo: userRepo}
}

func (s *AdminService) GetAllUsers(ctx context.Context) ([]*entity.User, error) {
	users, err := s.userRepo.GetAllUsers(ctx)
	return users, err
}
