package handlers

import (
	"Order5003/internal/logger"
	"Order5003/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeliverHandler struct {
	svc service.DelivererService
}

func NewDeliverHandler(svc service.DelivererService) *DeliverHandler {
	return &DeliverHandler{svc: svc}
}

func (h *DeliverHandler) Login(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	d, err := h.svc.GetDelivererByName(loginRequest.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	if d.Password != loginRequest.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"deliver_id": d.DelivererID,
		"username":   d.Name,
	})
}

func (h *DeliverHandler) GetOrderWaitingForDeliver(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	orders, err := h.svc.GetOrderWaitingForDeliver()
	logger.Info("GetOrderWaitingForDeliver: %v", orders)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"orders": orders,
	})
}

func (h *DeliverHandler) AcceptOrderDeliver(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	var acceptRequest struct {
		DeliverID int `json:"deliver_id"`
		OrderID   int `json:"order_id"`
	}
	if err := c.ShouldBindJSON(&acceptRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	err := h.svc.AcceptOrderDeliver(acceptRequest.DeliverID, acceptRequest.OrderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Order accepted",
	})
}

func (h *DeliverHandler) GetMyOrder(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	var getMyOrderRequest struct {
		DeliverID int `json:"deliver_id"`
	}
	if err := c.ShouldBindJSON(&getMyOrderRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	orders, err := h.svc.GetMyOrder(getMyOrderRequest.DeliverID)
	logger.Info("GetMyOrder: %v", orders)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"orders": orders,
	})
}

// ConfirmDeliver 确认配送
func (h *DeliverHandler) ConfirmDeliver(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	var confirmRequest struct {
		DeliverID int `json:"deliver_id"`
		OrderID   int `json:"order_id"`
	}
	if err := c.ShouldBindJSON(&confirmRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	err := h.svc.ConfirmDeliver(confirmRequest.DeliverID, confirmRequest.OrderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Order confirmed",
	})
}
