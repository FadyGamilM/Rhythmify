package user

import (
	"context"
	"fmt"

	"github.com/FadyGamilM/rhythmify/auth/api/dtos"
	"github.com/FadyGamilM/rhythmify/auth/core"
	"github.com/FadyGamilM/rhythmify/auth/domain"
	"github.com/FadyGamilM/rhythmify/auth/utils"
)

type userAuthService struct {
	userRepo core.UserRepo
}

// Signin implements core.AuthService.
func (*userAuthService) Signin(ctx context.Context, req interface{}) (res interface{}, err error) {
	panic("unimplemented")
}

// Signout implements core.AuthService.
func (*userAuthService) Signout(ctx context.Context, req interface{}) (res interface{}, err error) {
	panic("unimplemented")
}

// Signup implements core.AuthService.
func (us *userAuthService) Signup(ctx context.Context, req *dtos.SignupReqDto) (*dtos.SignupResDto, error) {
	// hash the password
	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("error trying to hash the password")
	}
	createdUser, err := us.userRepo.Insert(ctx, domain.User{
		Email:          req.Email,
		HashedPassword: hashed,
	})
	if err != nil {
		return nil, fmt.Errorf("error trying to persist user into database")
	}
	res := &dtos.SignupResDto{Id: createdUser.Id, Email: createdUser.Email}
	return res, nil
}

func NewUserAuthService(ur core.UserRepo) core.AuthService {
	return &userAuthService{
		userRepo: ur,
	}
}
