package handlers

import (
    "Order5003/internal/service"
    "net/http"

    "github.com/gin-gonic/gin"
)

type UserHandler struct {
    svc service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
    return &UserHandler{svc: svc}
}

func (h *UserHandler) LoginUser(c *gin.Context) {
    if c.Request.Method != http.MethodPost {
        c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
        return
    }
    var loginRequest struct{ Username, Password string }
    if err := c.ShouldBindJSON(&loginRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }
    user, err := h.svc.GetUserByUsername(loginRequest.Username)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }
    if user.Password != loginRequest.Password {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "id":       user.ID,
        "username": user.Username,
    })
}
