package repository

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/FadyGamilM/rhythmify/auth/core"
	"github.com/FadyGamilM/rhythmify/auth/db/postgres"
	"github.com/FadyGamilM/rhythmify/auth/domain"
)

const (
	INSERT_USER_QUERY = `
			INSERT INTO users (email, hashed_password) 
			VALUES ($1, $2) 
			RETURNING id, email, hashed_password
	`

	GET_USER_BY_ID_QUERY = `
		SELECT id, email, hashed_password
		FROM users 
		WHERE id = $1
	`
	GET_USER_BY_EMAIL_QUERY = `
		SELECT id, email, hashed_password
		FROM users 
		WHERE email = $1
	`
)

type PG_UserRepo struct {
	db *postgres.PG
}

// GetByEmail implements core.UserRepo.
// TODO => handle not found errors
func (ur *PG_UserRepo) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := &domain.User{}
	if err := ur.db.DB.QueryRowContext(ctx, GET_USER_BY_EMAIL_QUERY, email).Scan(&user.Id, &user.Email, &user.HashedPassword); err != nil {
		log.Println("[pg-user-rep (GetByEmail)] ➜ ", err.Error())
		return nil, fmt.Errorf("error trying to get user by email ➜ %v", err)
	}
	return user, nil
}

// GetByID implements core.UserRepo.
// TODO => handle not found errors
func (ur *PG_UserRepo) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	user := &domain.User{}
	if err := ur.db.DB.QueryRowContext(ctx, GET_USER_BY_ID_QUERY, id).Scan(&user.Id, &user.Email, &user.HashedPassword); err != nil {
		log.Println("[pg-user-rep (GetByID)]")
		return nil, fmt.Errorf("error trying to get user by id ➜ %v", err)
	}

	return user, nil
}

// Insert implements core.UserRepo.
// TODO => handle duplicate and null-value-resource errors
func (ur *PG_UserRepo) Insert(ctx context.Context, u domain.User) (*domain.User, error) {
	user := &domain.User{}
	if err := ur.db.DB.QueryRowContext(ctx, INSERT_USER_QUERY, u.Email, u.HashedPassword).Scan(&user.Id, &user.Email, &user.HashedPassword); err != nil {
		log.Println("[pg-user-rep (Insert)] ➜ ", err)
		return nil, errors.New(fmt.Sprintf("error trying to insert new user ➜ %v", err))
	}
	return user, nil
}

// the factory return the port not the concrete implementation
func NewPgUserRepo(db *postgres.PG) core.UserRepo {
	return &PG_UserRepo{
		db: db,
	}
}
