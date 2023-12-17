package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/FadyGamilM/rhythmify/auth/api/dtos"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func (h *Handler) HandleValidation(c *gin.Context) {
	reqDto := new(dtos.TokenValidationDto)
	if err := c.ShouldBindJSON(reqDto); err != nil {
		log.Printf("[auth microservice] couldn't bind the validation request : %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
	}

	token := reqDto.Token

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

		c.JSON(http.StatusAccepted, gin.H{
			"user_id": userID,
			"email":   user.Email,
		})
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "tampred token, security alert !!",
		})
		return
	}

}
