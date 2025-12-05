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

	//在下面实现/api/orders这一接口
	//可以通过搜索前端    fetch(`/api/orders/all?shopid=${encodeURIComponent(currentShopId)}`)

}
