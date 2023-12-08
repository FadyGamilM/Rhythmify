package api

import (
	"log"
	"net/http"

	"github.com/FadyGamilM/rhythmify/auth/api/dtos"
	"github.com/gin-gonic/gin"
)

func (h *Handler) HandleLogin(c *gin.Context) {
	reqDto := new(dtos.LoginReqDto)
	if err := c.ShouldBindJSON(reqDto); err != nil {
		log.Printf("[handler (HandleLogin)] ➜ %v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	result, err := h.AuthService.Signin(c, reqDto)
	if err != nil {
		if err != nil {
			log.Printf("[handler (HandleSignup)] ➜ %v\n", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "couldn't signup, try later",
			})
			return
		}
	}

	// set in the cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", result.Token, 3600*24, "", "", false, true)

	c.JSON(http.StatusAccepted, gin.H{})
	return
}
