package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type TokenValidationDto struct {
	Token string `json:"token"`
}

type TokenValidationResDto struct {
	UserId int64  `json:"user_id"`
	Email  string `json:"email"`
}

func (h *Handler) Authorize(c *gin.Context) {
	// Extract the Authorization header
	authHeader := c.GetHeader("Authorization")

	// Check if the Authorization header is empty
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization header is missing",
		})
		return
	}

	// The Authorization header should be in the format "Bearer TOKEN"
	// We can split the header into a slice where the first element is "Bearer" and the second is the actual token
	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Authorization header format",
		})
		return
	}

	// The actual token is the second element in the slice
	token := splitToken[1]

	tokenValidation := &TokenValidationDto{Token: token}
	reqBodyBytes, err := json.Marshal(tokenValidation)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "error marshaling the request to send it to auth microservice",
		})
		return
	}

	authHost := os.Getenv("AUTH_HOST")
	authPort := os.Getenv("AUTH_PORT")
	authURI := os.Getenv("AUTH_URI")
	validateEndpoint := fmt.Sprintf("%v/validate", authURI)
	validationResponse, err := CommunicateSync(authHost, authPort, validateEndpoint, "POST", reqBodyBytes)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "error during communication to auth microservice",
		})
		return
	}
	defer validationResponse.Body.Close()

	// Unmarshal the response body
	resDto := new(TokenValidationResDto)
	err = json.NewDecoder(validationResponse.Body).Decode(resDto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "error decoding the response from auth microservice",
		})
		return
	}

	// set the userid and user email in the request context
	c.Set("userId", resDto.UserId)
	c.Set("email", resDto.Email)

	c.Next()
}
