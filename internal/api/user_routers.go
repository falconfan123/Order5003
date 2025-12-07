package api

import (
	"Order5003/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, h *handlers.UserHandler) {
	g := r.Group("/user")
	g.GET("", func(c *gin.Context) { c.File("web/templates/user.html") })
	g.POST("/login", func(c *gin.Context) { h.LoginUser(c) })
	g.POST("/getusernamebyuserid", func(c *gin.Context) { h.GetUsernameByUserID(c) })
	//显示用户地址
	g.POST("/getuseraddressbyuserid", func(c *gin.Context) { h.GetUserAddressByUserID(c) })
}
