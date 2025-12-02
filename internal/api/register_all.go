package api

import (
	"Order5003/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterAllRoutes(r *gin.Engine, h *handlers.AppHandlers) {
	RegisterUserRoutes(r, h.User)
	RegisterShopRoutes(r, h.Shop)
	RegisterDeliverRoutes(r, h.Deliver)
	RegisterMenuRoutes(r, h.Menu)
	RegisterOrderRoutes(r, h.Order)
	RegisterMenuDishRoutes(r, h.MenuDish)
}
