package handlers

import (
	"Order5003/internal/store"
	"encoding/json"
	"net/http"
)

// UserHandler 处理用户相关的HTTP请求
type UserHandler struct {
	store *store.MemoryStore
}

// NewUserHandler 创建新的用户处理器
func NewUserHandler(store *store.MemoryStore) *UserHandler {
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

	// 检查用户是否存在
	user, err := h.store.GetUserByUsername(loginRequest.Username)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// 简单比较密码（实际应用中应该使用密码哈希）
	if user.Password != loginRequest.Password {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// 登录成功，返回用户信息（不包含密码）
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
		"role":     user.Role,
	})
}
