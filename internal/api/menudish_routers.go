package api

import (
	"Order5003/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterMenuDishRoutes(r *gin.Engine, h *handlers.MenuDishHandler) {
    menudishGroup := r.Group("/menudish")
    {
        menudishGroup.POST("/getallmenudishes", func(c *gin.Context) {
            h.GetAllMenuDishesByShopID(c)
        })
    }
}
