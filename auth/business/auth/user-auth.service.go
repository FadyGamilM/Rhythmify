package auth

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/FadyGamilM/rhythmify/auth/api/dtos"
	"github.com/FadyGamilM/rhythmify/auth/core"
	"github.com/FadyGamilM/rhythmify/auth/domain"
	"github.com/FadyGamilM/rhythmify/auth/utils"

	// "github.com/dgrijalva/jwt-go"
	"github.com/golang-jwt/jwt/v4"
)

type userAuthService struct {
	userRepo core.UserRepo
}

// Signin implements core.AuthService.
func (us *userAuthService) Signin(ctx context.Context, req *dtos.LoginReqDto) (*dtos.LoginResDto, error) {
	foundUser, err := us.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("error trying to retrieve user from database")
	}

	if err := utils.VerifyPassword(req.Password, foundUser.HashedPassword); err != nil {
		if err != nil {
			return nil, fmt.Errorf("error trying to verify password against database")
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": foundUser.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// now sign the token
	secret_key := os.Getenv("JWT_SECRET_KEY")
	tokenString, err := token.SignedString([]byte(secret_key))
	if err != nil {
		return nil, fmt.Errorf("error trying to encrypt the token ➜ %v", err)
	}
	return &dtos.LoginResDto{Token: tokenString}, nil
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
