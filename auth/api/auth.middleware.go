package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func (h *Handler) Auth(c *gin.Context) {
	// get the cookie from the request
	token, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	// decode the token from the cookie
	decodedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method : %v", t.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token",
		})
		return
	}

	var userID int64
	if claims, ok := decodedToken.Claims.(jwt.MapClaims); ok && decodedToken.Valid {
		// check if its expired or not
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "token is expired",
			})
			return
		}

		// extract token.sub to get the user id
		userID = int64(claims["sub"].(float64))

		// get the user ersource from datbase via this id
		user, err := h.UserService.FindByID(c, userID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err,
			})
		}

		// attach the resource into the req
		c.Set("user", user)

		// call the next handler in the chain
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "tampred token, security alert !!",
		})
		return
	}

}
