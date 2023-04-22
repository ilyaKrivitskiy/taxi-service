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

	authUser := router.Group("/auth")
	{
		authUser.POST("/sign-up", h.userSignUp)
		authUser.POST("/sign-in", h.userSignIn)
	}

	authDriver := router.Group("/auth")
	{
		authDriver.POST("/sign-up", h.driverSignUp)
		authDriver.POST("/sign-in", h.driverSignIn)
	}

	api := router.Group("/api")
	{
		orders := api.Group("/orders")
		{
			orders.POST("/", h.createOrder)
			orders.GET("/", h.getAllOrders)
			orders.GET("/:id", h.getOrderById)
			orders.PUT("/:id", h.updateOrder)
			orders.DELETE("/:id", h.deleteOrder)
		}
	}

	return router
}
