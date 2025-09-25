package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "Order5003/internal/models"
    "Order5003/internal/store"
)

// MenuHandler 处理菜单相关的HTTP请求
type MenuHandler struct {
    store *store.MemoryStore
}

// NewMenuHandler 创建新的菜单处理器
func NewMenuHandler(store *store.MemoryStore) *MenuHandler {
    return &MenuHandler{store: store}
}

// GetAllMenuItems 获取所有菜单项目
func (h *MenuHandler) GetAllMenuItems(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    
    items := h.store.GetAllMenuItems()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(items)
}

// GetMenuItemByID 根据ID获取菜单项目
func (h *MenuHandler) GetMenuItemByID(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid menu item ID", http.StatusBadRequest)
        return
    }
    
    item, err := h.store.GetMenuItemByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(item)
}

// CreateMenuItem 创建新的菜单项目（管理员功能）
func (h *MenuHandler) CreateMenuItem(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    
    var item models.MenuItem
    if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    
    createdItem := h.store.CreateMenuItem(item)
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(createdItem)
}

// UpdateMenuItem 更新菜单项目（管理员功能）
func (h *MenuHandler) UpdateMenuItem(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPut {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid menu item ID", http.StatusBadRequest)
        return
    }
    
    var item models.MenuItem
    if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    
    updatedItem, err := h.store.UpdateMenuItem(id, item)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedItem)
}

// DeleteMenuItem 删除菜单项目（管理员功能）
func (h *MenuHandler) DeleteMenuItem(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodDelete {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid menu item ID", http.StatusBadRequest)
        return
    }
    
    if err := h.store.DeleteMenuItem(id); err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    
    w.WriteHeader(http.StatusNoContent)
}