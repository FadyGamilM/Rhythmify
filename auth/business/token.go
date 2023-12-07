package business

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	jwt_private_key = os.Getenv("JWT_PRIVATE_KEY")
)

func GenerateJwtToken(claims *JwtClaims, expirationTime time.Time) (token string, err error) {
	// setup claims props
	claims.ExpiresAt = expirationTime.Unix()
	// sign the token
	signedTokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// encrypt it with private key
	token, err = signedTokenWithClaims.SignedString([]byte(jwt_private_key))
	// handle errors if any
	if err != nil {
		return "", errors.New("error during token encryption")
	}
	// return token
	return token, nil
}
