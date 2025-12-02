package api

import (
	"Order5003/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, h *handlers.UserHandler) {
	g := r.Group("/user")
	g.GET("", func(c *gin.Context) { c.File("web/templates/user.html") })
	g.POST("/login", func(c *gin.Context) { h.LoginUser(c.Writer, c.Request) })
}
