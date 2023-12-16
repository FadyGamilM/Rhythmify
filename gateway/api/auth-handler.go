package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type SignupReqDto struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type SignupResDto struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
}

// receive the body from the request
// make a sync call to the auth service and forward the body to it
// return the response to the user
func (h *Handler) HandleSignup(c *gin.Context) {
	signupReq := new(SignupReqDto)
	if err := c.ShouldBindJSON(signupReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	reqBodyBytes, err := json.Marshal(signupReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error marshling the request to send it to auth microservice",
		})
		return
	}

	authHost := os.Getenv("AUTH_HOST")
	authPort := os.Getenv("AUTH_PORT")
	authURI := os.Getenv("AUTH_URI")
	signupEndpoint := fmt.Sprintf("%v/signup", authURI)
	signupResponse, err := CommunicateSync(authHost, authPort, signupEndpoint, "POST", reqBodyBytes)
	log.Printf("the error : %v\n", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error during communication to auth microservice",
		})
		return
	}
	defer signupResponse.Body.Close()

	// Extract the response body from http.Response
	responseBody, err := ioutil.ReadAll(signupResponse.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error reading signup response body",
		})
		return
	}

	signupResDto := new(SignupResDto)
	if err := json.Unmarshal(responseBody, signupResDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error unmarshal the response body to dto",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"response": signupResDto,
	})
}
