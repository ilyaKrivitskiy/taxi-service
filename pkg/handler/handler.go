package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyaKrivitskiy/taxi-service/pkg/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(servicies *service.Service) *Handler {
	return &Handler{service: servicies}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	return router
}
