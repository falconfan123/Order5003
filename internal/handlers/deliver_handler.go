package handlers

import (
    "Order5003/internal/service"
    "encoding/json"
    "net/http"
)

type DeliverHandler struct {
    svc service.DelivererService
}

func NewDeliverHandler(svc service.DelivererService) *DeliverHandler {
    return &DeliverHandler{svc: svc}
}

func (h *DeliverHandler) Login(w http.ResponseWriter, r *http.Request) {
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
    d, err := h.svc.GetDelivererByName(loginRequest.Username)
    if err != nil {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }
    if d.Password != loginRequest.Password {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "id":       d.DelivererID,
        "username": d.Name,
    })
}
