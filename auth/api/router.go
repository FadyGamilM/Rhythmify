package api

import "github.com/gin-gonic/gin"

type handler struct {
	Router *gin.Engine
}

func newRouter() *gin.Engine {
	return gin.Default()
}

func NewHandler() *handler {
	r := newRouter()
	return &handler{Router: r}
}

func (h *handler) SetupEndpoints() {
	api := h.Router.Group("/api/v1")
	api.GET("/health", HandleHealthCheck)
	api.POST("/auth/login", HandleLogin)
}
