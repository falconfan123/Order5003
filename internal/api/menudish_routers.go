package api

import (
	"Order5003/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterMenuDishRoutes(r *gin.Engine, h *handlers.MenuDishHandler) {
	menudishGroup := r.Group("/menudish")
	{
		//得到shopid下的所有菜品
		menudishGroup.POST("/getmenudishesbyid", func(c *gin.Context) {
			h.GetMenuDishesByShopID(c)
		})
	}
}
