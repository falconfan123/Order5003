package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "Order5003/internal/models"
    "Order5003/internal/store"
)

// OrderHandler 处理订单相关的HTTP请求
type OrderHandler struct {
    store *store.MemoryStore
}

// NewOrderHandler 创建新的订单处理器
func NewOrderHandler(store *store.MemoryStore) *OrderHandler {
    return &OrderHandler{store: store}
}

// CreateOrder 创建新订单
func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    
    var request models.NewOrderRequest
    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    
    // 计算订单总价
    var total float64
    for _, item := range request.Items {
        total += item.Price * float64(item.Quantity)
    }
    
    order := models.Order{
        TableNumber: request.TableNumber,
        Items:       request.Items,
        Total:       total,
    }
    
    createdOrder := h.store.CreateOrder(order)
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(createdOrder)
}

// GetOrderByID 根据ID获取订单
func (h *OrderHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid order ID", http.StatusBadRequest)
        return
    }
    
    order, err := h.store.GetOrderByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(order)
}

// GetAllOrders 获取所有订单（管理员功能）
func (h *OrderHandler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    
    orders := h.store.GetAllOrders()
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(orders)
}

// UpdateOrderStatus 更新订单状态（管理员/员工功能）
func (h *OrderHandler) UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPut {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid order ID", http.StatusBadRequest)
        return
    }
    
    var statusUpdate struct {
        Status models.OrderStatus `json:"status"`
    }
    
    if err := json.NewDecoder(r.Body).Decode(&statusUpdate); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    
    updatedOrder, err := h.store.UpdateOrderStatus(id, statusUpdate.Status)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedOrder)
}

// GetRandomTableNumber 获取随机桌号（用于顾客点餐）
func (h *OrderHandler) GetRandomTableNumber(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    
    tableNumber := h.store.GetRandomTableNumber()
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"table_number": tableNumber})
}