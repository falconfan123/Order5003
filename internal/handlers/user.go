package handlers

import (
	"Order5003/internal/models"
	"Order5003/internal/store"
	"encoding/json"
	"net/http"
)

// UserHandler 处理用户相关的HTTP请求
type UserHandler struct {
	store store.Store
}

// NewUserHandler 创建新的用户处理器
func NewUserHandler(store store.Store) *UserHandler {
	return &UserHandler{store: store}
}

// Login 处理用户登录请求
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	shop, err := h.store.GetShopByName(loginRequest.Username)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}
	if shop.Password != loginRequest.Password {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":       shop.ID,
		"username": shop.ShopName,
		"role":     models.UserRoleAdmin,
	})
}
