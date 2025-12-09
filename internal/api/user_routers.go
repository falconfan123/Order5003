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
	// 更改用户地址
	g.POST("/updateuseraddressbyuserid", func(c *gin.Context) { h.UpdateUserAddressByUserID(c) })
	//显示用户电话
	g.POST("/getuserphonebyuserid", func(c *gin.Context) { h.GetUserPhoneByUserID(c) })
	//支付订单
	g.POST("/payorder", func(c *gin.Context) { h.PayOrder(c) })
	//取消订单
	g.POST("/cancelorder", func(c *gin.Context) { h.CancelOrder(c) })
}
