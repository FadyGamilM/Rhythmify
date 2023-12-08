package user

import (
	"context"
	"fmt"

	"github.com/FadyGamilM/rhythmify/auth/core"
	"github.com/FadyGamilM/rhythmify/auth/domain"
)

type UserService struct {
	UserRepo core.UserRepo
}

// FindByID implements core.UserService.
func (us *UserService) FindByID(ctx context.Context, id int64) (*domain.User, error) {
	user, err := us.UserRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error trying to find user with id = %v , error = %v", id, err)
	}
	return user, nil
}

func NewUserService(ur core.UserRepo) core.UserService {
	return &UserService{UserRepo: ur}
}
