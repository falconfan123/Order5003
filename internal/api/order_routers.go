package api

import (
	"Order5003/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterOrderRoutes(r *gin.Engine, h *handlers.OrderHandler) {
	r.Any("/orders", func(c *gin.Context) {
		switch c.Request.Method {
		case "POST":
			h.CreateOrder(c)
		default:
			c.Status(405)
		}
	})
	r.GET("/orders/getall", func(c *gin.Context) { h.GetAllOrders(c) })
	r.PUT("/orders/status", func(c *gin.Context) { h.UpdateOrderStatus(c) })
}
