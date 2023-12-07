package core

import (
	"context"

	"github.com/FadyGamilM/rhythmify/auth/domain"
)

type AuthService interface {
	Signup(ctx context.Context, req interface{}) (res interface{}, err error)
	Signin(ctx context.Context, req interface{}) (res interface{}, err error)
	Signout(ctx context.Context, req interface{}) (res interface{}, err error)
}

type UserRepo interface {
	Insert(context.Context, domain.User) (*domain.User, error)
	GetByID(ctx context.Context, id int) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
}