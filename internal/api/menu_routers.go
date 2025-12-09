package api

import (
	"Order5003/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterMenuRoutes(r *gin.Engine, h *handlers.MenuHandler) {
	g := r.Group("/menu")
	//得到shopid下的所有菜单
	g.POST("/getmenubyshopid", func(c *gin.Context) { h.GetMenuByShopID(c) })
	//得到menuid下的所有菜品
	g.POST("/getdishesbymenuid", func(c *gin.Context) { h.GetDishesByMenuID(c) })
	// 更改菜单
	g.POST("/updatemenu", func(c *gin.Context) { h.UpdateMenu(c) })
}
