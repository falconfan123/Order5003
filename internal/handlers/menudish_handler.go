package handlers

import (
	"Order5003/internal/logger"
	"Order5003/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MenuDishHandler struct {
	shopSvc service.ShopService
}

func NewMenuDishHandler(s service.ShopService) *MenuDishHandler {
	return &MenuDishHandler{shopSvc: s}
}

type GetMenuDishesReq struct {
	ShopID int `json:"shop_id"`
}

func (h *MenuDishHandler) GetMenuDishesByShopID(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	var req GetMenuDishesReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数格式错误，请传 JSON 格式的 shop_id"})
		return
	}
	logger.Info("GetAllMenuDishesByShopID", ":shop_id", req.ShopID)
	menuDishes, err := h.shopSvc.GetMenuDishesByShopID(req.ShopID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Info("GetAllMenuDishesByShopID", ":menu_dishes", menuDishes)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": menuDishes,
	})
}
