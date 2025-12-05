package api

import (
	"Order5003/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterOrderRoutes(r *gin.Engine, h *handlers.OrderHandler) {
	orders := r.Group("/orders")
	{
		orders.POST("/create", h.CreateOrder)
		orders.GET("/getall", h.GetAllOrders)
		orders.PUT("/status", h.UpdateOrderStatus)
	}
}
