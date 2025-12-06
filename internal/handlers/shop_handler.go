package handlers

import (
	"Order5003/internal/logger"
	"Order5003/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ShopHandler struct {
	svc service.ShopService
}

func NewShopHandler(svc service.ShopService) *ShopHandler {
	return &ShopHandler{svc: svc}
}

func (h *ShopHandler) Login(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	var loginRequest struct{ Username, Password string }
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	shop, err := h.svc.GetShopByName(loginRequest.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	if shop.Password != loginRequest.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"shop_id":  shop.ShopID,
		"username": shop.ShopName,
	})
}

func (h *ShopHandler) GetAll(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	shops, err := h.svc.GetAllShops()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, shops)
}

func (h *ShopHandler) GetShopNameByShopID(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	var request struct{ ShopID int }
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	shop, err := h.svc.GetShopNameByShopID(request.ShopID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"shop_name": shop})
}

func (h *ShopHandler) GetOrdersByShopID(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	shopIdstr := c.Query("shop_id")
	if shopIdstr == "" {
		shopIdstr = c.Query("shopid")
	}
	logger.Info("shopIdstr", shopIdstr)
	ShopID, err := strconv.Atoi(shopIdstr)
	logger.Info("ShopID", ShopID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid shop_id"})
		return
	}
	orders, err := h.svc.GetOrdersByShopID(ShopID)
	logger.Info("orders", orders)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, orders)
}
