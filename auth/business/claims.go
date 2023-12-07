package business

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Role string

var (
	Customer Role = "CUSTOMER"
	Premium  Role = "PREMIUM"
	Admin    Role = "ADMIN"
)

type JwtClaims struct {
	UserId int64 `json:"user_id,omitempty"`
	Role   Role  `json:"role,omitempty"`
	jwt.StandardClaims
}

// now lets implement the Claims interface from jwt package
// this method won't be used in the login api, but we have to implement the claims interface anyway
func (jc JwtClaims) Valid() error {
	now := time.Now().UTC().Unix()
	// verify if token expired or not
	// if expired -> notExpired will be false
	// if not expired -> notExpired will be true
	notExpired := jc.VerifyExpiresAt(now, true)
	if notExpired {
		return nil
	}
	return errors.New("token is expired")
}
