package api

import (
	"log"
	"net/http"

	"github.com/FadyGamilM/rhythmify/auth/api/dtos"
	"github.com/gin-gonic/gin"
)

func (h *Handler) HandleSignup(c *gin.Context) {
	reqDto := new(dtos.SignupReqDto)

	if err := c.ShouldBindJSON(reqDto); err != nil {
		log.Printf("[handler (HandleSignup)] ➜ %v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	result, err := h.AuthService.Signup(c, reqDto)
	if err != nil {
		log.Printf("[handler (HandleSignup)] ➜ %v\n", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "couldn't signup, try later",
		})
		return
	}

	log.Printf("singup request is completed with id : %v and email %v\n", result.Id, result.Email)

	c.JSON(http.StatusCreated, result)
	return
}
