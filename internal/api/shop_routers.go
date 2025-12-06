package api

import (
	"Order5003/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterShopRoutes(r *gin.Engine, h *handlers.ShopHandler) {
	g := r.Group("/shop")
	g.GET("", func(c *gin.Context) { c.File("web/templates/shop.html") })
	g.POST("/login", func(c *gin.Context) { h.Login(c) })
	g.POST("/getall", func(c *gin.Context) { h.GetAll(c) })
	g.POST("/getshopnamebyshopid", func(c *gin.Context) { h.GetShopNameByShopID(c) })
	g.POST("/getdeliveryfeebyshopid", func(c *gin.Context) { h.GetDeliveryFeeByShopID(c) })
	g.GET("/getordersbyshopid", func(c *gin.Context) { h.GetOrdersByShopID(c) })
	g.POST("/getbussinesshoursbyshopid", func(c *gin.Context) { h.GetBusinessHoursByShopID(c) })
	g.POST("/getshoptypebyshopid", func(c *gin.Context) { h.GetShopTypeByShopID(c) })
	g.POST("/updateshopstatus", func(c *gin.Context) { h.UpdateShopStatus(c) })
}
