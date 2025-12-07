package handlers

import (
	"Order5003/internal/logger"
	"Order5003/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MenuHandler struct {
	svc service.MenuService
}

func NewMenuHandler(svc service.MenuService) *MenuHandler {
	return &MenuHandler{svc: svc}
}

func (h *MenuHandler) GetMenuByShopID(c *gin.Context) {
	if c.Request.Method != "POST" {
		c.JSON(405, gin.H{"error": "method not allowed"})
		return
	}
	shopID, err := strconv.Atoi(c.PostForm("shop_id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "shop_id is required"})
		return
	}
	menu, err := h.svc.GetMenuByShopID(c, shopID)
	logger.Info("GetMenuByShopID", zap.Int("shopID", shopID), zap.Any("menu", menu))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, menu)
}

func (h *MenuHandler) GetDishesByMenuID(c *gin.Context) {
	if c.Request.Method != "POST" {
		c.JSON(405, gin.H{"error": "method not allowed"})
		return
	}
	menuID, err := strconv.Atoi(c.PostForm("menu_id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "menu_id is required"})
		return
	}
	dishes, err := h.svc.GetDishesByMenuID(c, menuID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, dishes)
}
