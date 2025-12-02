package handlers

import (
    "Order5003/internal/service"
    "encoding/json"
    "net/http"
)

type ShopHandler struct {
    svc service.ShopService
}

func NewShopHandler(svc service.ShopService) *ShopHandler {
    return &ShopHandler{svc: svc}
}

func (h *ShopHandler) Login(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    var loginRequest struct{ Username, Password string }
    if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    shop, err := h.svc.GetShopByName(loginRequest.Username)
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
        "id":       shop.ShopID,
        "username": shop.ShopName,
    })
}

func (h *ShopHandler) GetAll(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    shops, err := h.svc.GetAllShops()
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(shops)
}
