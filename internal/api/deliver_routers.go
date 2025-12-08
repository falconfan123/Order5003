package api

import (
	"Order5003/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterDeliverRoutes(r *gin.Engine, h *handlers.DeliverHandler) {
	g := r.Group("/deliver")
	g.GET("", func(c *gin.Context) { c.File("web/templates/deliverer.html") })
	g.POST("/login", func(c *gin.Context) { h.Login(c) })
	g.POST("/getorderwatingfordeliver", func(c *gin.Context) { h.GetOrderWaitingForDeliver(c) })
	g.POST("/acceptorderdeliver", func(c *gin.Context) { h.AcceptOrderDeliver(c) })
	g.POST("/myorder", func(c *gin.Context) { h.GetMyOrder(c) })
	g.POST("/confirmdeliver", func(c *gin.Context) { h.ConfirmDeliver(c) })
}
