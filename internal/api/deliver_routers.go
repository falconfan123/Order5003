package api

import (
    "Order5003/internal/handlers"
    "github.com/gin-gonic/gin"
)

func RegisterDeliverRoutes(r *gin.Engine, h *handlers.DeliverHandler) {
    g := r.Group("/deliver")
    g.GET("", func(c *gin.Context) { c.File("web/templates/deliverer.html") })
    g.POST("/login", func(c *gin.Context) { h.Login(c.Writer, c.Request) })
}

