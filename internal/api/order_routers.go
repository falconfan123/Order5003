package api

import (
    "Order5003/internal/handlers"
    "github.com/gin-gonic/gin"
)

func RegisterOrderRoutes(r *gin.Engine, h *handlers.OrderHandler) {
    r.Any("/api/orders", func(c *gin.Context) {
        switch c.Request.Method {
        case "POST":
            h.CreateOrder(c.Writer, c.Request)
        default:
            c.Status(405)
        }
    })
    r.GET("/api/orders/all", func(c *gin.Context) { h.GetAllOrders(c.Writer, c.Request) })
    r.PUT("/api/orders/status", func(c *gin.Context) { h.UpdateOrderStatus(c.Writer, c.Request) })
}

