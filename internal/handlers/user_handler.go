package handlers

import (
    "Order5003/internal/service"
    "encoding/json"
    "net/http"
)

type UserHandler struct {
    svc service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
    return &UserHandler{svc: svc}
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    var loginRequest struct{ Username, Password string }
    if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    user, err := h.svc.GetUserByUsername(loginRequest.Username)
    if err != nil {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }
    if user.Password != loginRequest.Password {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "id":       user.ID,
        "username": user.Username,
    })
}
