package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	router *gin.Engine
}

func newRouter() *gin.Engine {
	return gin.Default()
}

func NewHandler() *Handler {
	r := newRouter()
	return &Handler{
		router: r,
	}
}

func (h *Handler) SetupEndpoints() {
	gatewayApis := h.router.Group("/api/v1/gateway")
	gatewayApis.POST("/auth/signup", h.HandleSignup)
	gatewayApis.POST("/auth/login", h.HandleLogin)
	gatewayApis.POST("/auth/validate", h.Authorize)
	gatewayApis.POST("/video/upload", h.UploadVideo)
	gatewayApis.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"response": "healthy",
		})
	})

}
