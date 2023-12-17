package api

import (
	"net/http"

	"github.com/FadyGamilM/rhythmify/auth/core"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Router      *gin.Engine
	AuthService core.AuthService
	UserService core.UserService
}

func newRouter() *gin.Engine {
	return gin.Default()
}

func NewHandler(as core.AuthService, us core.UserService) *Handler {
	r := newRouter()
	return &Handler{Router: r, AuthService: as, UserService: us}
}

func (h *Handler) SetupEndpoints() {
	api := h.Router.Group("/api/v1")
	api.POST("/auth/signup", h.HandleSignup)
	api.POST("/auth/login", h.HandleLogin)
	api.POST("/auth/validate", h.HandleValidation)
	api.GET("/delete-product", h.Auth, func(ctx *gin.Context) {

		loggedInUser, _ := ctx.Get("user")
		ctx.JSON(http.StatusAccepted, gin.H{
			"user": loggedInUser,
		})
	})
}
