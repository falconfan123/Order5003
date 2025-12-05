package api

import (
	"Order5003/internal/handlers"
	"Order5003/internal/logger"

	"github.com/gin-gonic/gin"
)

func RegisterShopRoutes(r *gin.Engine, h *handlers.ShopHandler) {
	g := r.Group("/shop")
	logger.Info("开始注册商家端路由，分组路径：/shop") // 加日志

	g.GET("", func(c *gin.Context) {
		logger.Info("访问/shop根路径，返回shop.html")
		c.File("web/templates/shop.html")
	})
	g.POST("/login", func(c *gin.Context) { h.Login(c) })
	g.POST("/getall", func(c *gin.Context) { h.GetAll(c) })
	g.POST("/getshopnamebyshopid", func(c *gin.Context) {
		logger.Info("接收到/getshopnamebyshopid请求，开始处理")
		h.GetShopNameByShopID(c)
	})

	//在下面实现/shop/getordersbyshopid这一接口
	//可以通过搜索前端    fetch(`/shop/getordersbyshopid?shopid=${encodeURIComponent(currentShopId)}`)
	g.GET("/getordersbyshopid", func(c *gin.Context) { h.GetOrdersByShopID(c) })
	logger.Info("商家端路由注册完成，包含GET /shop/getordersbyshopid")

}
