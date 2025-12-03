package handlers

import (
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
        "id":       d.DelivererID,
        "username": d.Name,
    })
}
