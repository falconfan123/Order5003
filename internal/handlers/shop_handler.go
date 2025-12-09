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
	logger.Info("Shop_id ", shop.ShopID, "shop_name ", shop.ShopName, "password ", shop.Password)
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

func (h *ShopHandler) GetDeliveryFeeByShopID(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	var request struct {
		ShopID int `json:"shop_id"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	fee, err := h.svc.GetDeliveryFeeByShopID(request.ShopID)
	logger.Info("fee", fee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"delivery_fee": fee})
}

func (h *ShopHandler) GetBusinessHoursByShopID(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	var request struct {
		ShopID int `json:"shop_id"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	businessHours, err := h.svc.GetBusinessHoursByShopID(request.ShopID)
	logger.Info("businessHours", businessHours)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"business_hours": businessHours})
}

func (h *ShopHandler) GetShopTypeByShopID(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	var request struct {
		ShopID int `json:"shop_id"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	logger.Info("GetShopTypeByShopID ShopID", request.ShopID)
	shopType, err := h.svc.GetShopTypeByShopID(request.ShopID)
	logger.Info("shopType", shopType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"shop_type": shopType})
}

func (h *ShopHandler) UpdateShopStatus(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	var request struct {
		ShopID int
		Status int
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	status, err := h.svc.UpdateShopStatus(request.ShopID, request.Status)
	logger.Info("status", status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": status})
}

func (h *ShopHandler) GetOrderDishesByOrderID(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	var request struct {
		OrderID int
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	dishes, err := h.svc.GetDishesByOrderID(request.OrderID)
	logger.Info("dishes", dishes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, dishes)
}

func (h *ShopHandler) AcceptOrder(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	var request struct {
		OrderID int
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	err := h.svc.AcceptOrder(request.OrderID)
	logger.Info("err", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order accepted"})
}

func (h *ShopHandler) WaitingForDeliveryOrder(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	var request struct {
		OrderID int
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	err := h.svc.WaitingForDeliveryOrder(request.OrderID)
	logger.Info("err", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order waiting for preparing"})
}

func (h *ShopHandler) GetShopStatusByShopID(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	var request struct {
		ShopID int `json:"shop_id"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	status, err := h.svc.GetShopStatusByShopID(request.ShopID)
	logger.Info("status", status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": status})
}

func (h *ShopHandler) StopDish(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	var request struct {
		DishID int `json:"dish_id"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	err := h.svc.StopDish(request.DishID)
	logger.Info("err", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Dish stopped"})
}

func (h *ShopHandler) StartDish(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	var request struct {
		DishID int `json:"dish_id"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	err := h.svc.StartDish(request.DishID)
	logger.Info("err", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Dish started"})
}

func (h *ShopHandler) GetTodayOrderCountByShopID(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	var request struct {
		ShopID int `json:"shop_id"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	count, err := h.svc.GetTodayOrderCountByShopID(request.ShopID)
	//同时写入到日志中去
	h.svc.WriteTodayOrderCountByShopID(request.ShopID, count)
	logger.Info("count", count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count})
}

func (h *ShopHandler) GetTodayRevenueByShopID(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	var request struct {
		ShopID int `json:"shop_id"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	revenue, err := h.svc.GetTodayRevenueByShopID(request.ShopID)
	//同时写入到日志中去
	h.svc.WriteTodayRevenueByShopID(request.ShopID, revenue)
	logger.Info("revenue", revenue)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"revenue": revenue})
}

func (h *ShopHandler) GetAllOrderCountByShopID(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	var request struct {
		ShopID int `json:"shop_id"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	count, err := h.svc.GetAllOrderCountByShopID(request.ShopID)
	logger.Info("count", count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count})
}

// GetAllRevenueByShopID 获取指定店铺的所有营业额记录
func (h *ShopHandler) GetAllRevenueByShopID(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	var request struct {
		ShopID int `json:"shop_id"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	revenue, err := h.svc.GetAllRevenueByShopID(request.ShopID)
	logger.Info("revenue", revenue)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"revenue": revenue})
}

// UpdateShopInfo 更新店铺资料
func (h *ShopHandler) UpdateShopInfo(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	var request struct {
		ShopID        int     `json:"shop_id"`
		ShopName      string  `json:"shop_name"`
		DeliveryRange float64 `json:"delivery_range"`
		DeliveryFee   float64 `json:"delivery_fee"`
		BusinessHours string  `json:"business_hours"`
		Type          int     `json:"type"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	err := h.svc.UpdateShopInfo(request.ShopID, request.ShopName, request.DeliveryRange, request.DeliveryFee, request.BusinessHours, request.Type)
	logger.Info("err", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Shop info updated"})
}

// SaveDish 保存菜品
func (h *ShopHandler) SaveDish(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	var request struct {
		DishID   int     `json:"dish_id"`
		DishName string  `json:"dish_name"`
		Price    float64 `json:"price"`
		Stock    int     `json:"stock"`
		Status   int     `json:"status"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	err := h.svc.SaveDish(request.DishID, request.DishName, request.Price, request.Stock, request.Status)
	logger.Info("err", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Dish saved"})
}

// GetShopInfoForUser 获取指定店铺的对用户展示的资料
func (h *ShopHandler) GetShopInfoForUser(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	var request struct {
		ShopID int `json:"shop_id"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	shop, err := h.svc.GetShopInfoForUser(request.ShopID)
	logger.Info("shop", shop)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"shop_phone": shop.Phone, "shop_name": shop.ShopName})
}

func (h *ShopHandler) AddDish(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	var request struct {
		DishName string  `json:"dish_name"`
		Price    float64 `json:"price"`
		Stock    int     `json:"stock"`
		Status   int     `json:"status"`
		ShopID   int     `json:"shop_id"`
		MenuName string  `json:"menu_name"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	err := h.svc.AddDish(request.ShopID, request.MenuName, request.DishName, request.Price, request.Stock, request.Status)
	logger.Info("err", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Dish added"})
}

// RefuseOrderByShop 商家拒单
func (h *ShopHandler) RefuseOrderByShop(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	var request struct {
		OrderID int `json:"order_id"`
		ShopID  int `json:"shop_id"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	err := h.svc.RefuseOrderByShop(request.OrderID, request.ShopID)
	logger.Info("err", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order refused"})
}
