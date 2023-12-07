package api

import (
	"github.com/FadyGamilM/rhythmify/auth/core"
	"github.com/gin-gonic/gin"
)

type handler struct {
	Router      *gin.Engine
	UserService core.AuthService
}

func newRouter() *gin.Engine {
	return gin.Default()
}

func NewHandler(us core.AuthService) *handler {
	r := newRouter()
	return &handler{Router: r, UserService: us}
}

func (h *handler) SetupEndpoints() {
	api := h.Router.Group("/api/v1")
	api.POST("/auth/signup", h.HandleSignup)
}
