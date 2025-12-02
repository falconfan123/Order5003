package store

import (
    "sync"
    "time"
    "errors"
    "strconv"
    "math/rand"
    "Order5003/internal/models"
)

// MemoryStore 实现内存数据库存储
type MemoryStore struct {
    mu       sync.RWMutex
    menuItems map[int]models.MenuItem
    orders    map[int]models.Order
    users     map[int]models.User
    shops     map[string]models.Shop
    nextMenuItemID int
    nextOrderID    int
    nextUserID     int
}

// NewMemoryStore 创建一个新的内存数据库实例
func NewMemoryStore() *MemoryStore {
    store := &MemoryStore{
        menuItems:      make(map[int]models.MenuItem),
        orders:         make(map[int]models.Order),
        users:          make(map[int]models.User),
        shops:          make(map[string]models.Shop),
        nextMenuItemID: 1,
        nextOrderID:    1,
        nextUserID:     1,
    }
    
    // 添加一些初始数据
    store.seedData()
    
    return store
}

// seedData 为数据库添加一些初始数据
func (s *MemoryStore) seedData() {
    // 添加一些菜单项目
    menuItems := []models.MenuItem{
        {
            Name:        "宫保鸡丁",
            Description: "香辣可口的经典川菜",
            Price:       48.00,
            Category:    "川菜",
            IsAvailable: true,
            CreatedAt:   time.Now(),
            UpdatedAt:   time.Now(),
        },
        {
            Name:        "麻婆豆腐",
            Description: "麻辣鲜香的传统川菜",
            Price:       32.00,
            Category:    "川菜",
            IsAvailable: true,
            CreatedAt:   time.Now(),
            UpdatedAt:   time.Now(),
        },
        {
            Name:        "清蒸鱼",
            Description: "鲜嫩多汁的粤菜",
            Price:       88.00,
            Category:    "粤菜",
            IsAvailable: true,
            CreatedAt:   time.Now(),
            UpdatedAt:   time.Now(),
        },
        {
            Name:        "北京烤鸭",
            Description: "北京特色美食",
            Price:       198.00,
            Category:    "北京菜",
            IsAvailable: true,
            CreatedAt:   time.Now(),
            UpdatedAt:   time.Now(),
        },
        {
            Name:        "西湖醋鱼",
            Description: "酸甜可口的浙江菜",
            Price:       78.00,
            Category:    "浙江菜",
            IsAvailable: true,
            CreatedAt:   time.Now(),
            UpdatedAt:   time.Now(),
        },
    }
    
    for _, item := range menuItems {
        item.ID = s.nextMenuItemID
        s.menuItems[s.nextMenuItemID] = item
        s.nextMenuItemID++
    }
    
    // 添加一个管理员用户
    admin := models.User{
        Username:  "admin",
        Password:  "password", // 简单密码，实际应用中应该加密
        Role:      models.UserRoleAdmin,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }
    admin.ID = s.nextUserID
    s.users[s.nextUserID] = admin
    s.nextUserID++

    s.shops["admin"] = models.Shop{
        ID:        1,
        ShopName:  "admin",
        Password:  "password",
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }
}

// GetAllMenuItems 获取所有菜单项目
func (s *MemoryStore) GetAllMenuItems() []models.MenuItem {
    s.mu.RLock()
    defer s.mu.RUnlock()
    
    items := make([]models.MenuItem, 0, len(s.menuItems))
    for _, item := range s.menuItems {
        items = append(items, item)
    }
    return items
}

// GetMenuItemByID 根据ID获取菜单项目
func (s *MemoryStore) GetMenuItemByID(id int) (models.MenuItem, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    
    item, exists := s.menuItems[id]
    if !exists {
        return models.MenuItem{}, errors.New("menu item not found")
    }
    return item, nil
}

// CreateMenuItem 创建新的菜单项目
func (s *MemoryStore) CreateMenuItem(item models.MenuItem) models.MenuItem {
    s.mu.Lock()
    defer s.mu.Unlock()
    
    item.ID = s.nextMenuItemID
    item.CreatedAt = time.Now()
    item.UpdatedAt = time.Now()
    s.menuItems[item.ID] = item
    s.nextMenuItemID++
    
    return item
}

// UpdateMenuItem 更新菜单项目
func (s *MemoryStore) UpdateMenuItem(id int, updatedItem models.MenuItem) (models.MenuItem, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    
    _, exists := s.menuItems[id]
    if !exists {
        return models.MenuItem{}, errors.New("menu item not found")
    }
    
    updatedItem.ID = id
    updatedItem.CreatedAt = s.menuItems[id].CreatedAt
    updatedItem.UpdatedAt = time.Now()
    s.menuItems[id] = updatedItem
    
    return updatedItem, nil
}

// DeleteMenuItem 删除菜单项目
func (s *MemoryStore) DeleteMenuItem(id int) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    
    _, exists := s.menuItems[id]
    if !exists {
        return errors.New("menu item not found")
    }
    
    delete(s.menuItems, id)
    return nil
}

// CreateOrder 创建新订单
func (s *MemoryStore) CreateOrder(order models.Order) models.Order {
    s.mu.Lock()
    defer s.mu.Unlock()
    
    order.ID = s.nextOrderID
    order.Status = models.OrderStatusPending
    order.CreatedAt = time.Now()
    order.UpdatedAt = time.Now()
    s.orders[order.ID] = order
    s.nextOrderID++
    
    return order
}

// GetOrderByID 根据ID获取订单
func (s *MemoryStore) GetOrderByID(id int) (models.Order, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    
    order, exists := s.orders[id]
    if !exists {
        return models.Order{}, errors.New("order not found")
    }
    return order, nil
}

// GetAllOrders 获取所有订单
func (s *MemoryStore) GetAllOrders() []models.Order {
    s.mu.RLock()
    defer s.mu.RUnlock()
    
    orders := make([]models.Order, 0, len(s.orders))
    for _, order := range s.orders {
        orders = append(orders, order)
    }
    return orders
}

// UpdateOrderStatus 更新订单状态
func (s *MemoryStore) UpdateOrderStatus(id int, status models.OrderStatus) (models.Order, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    
    order, exists := s.orders[id]
    if !exists {
        return models.Order{}, errors.New("order not found")
    }
    
    order.Status = status
    order.UpdatedAt = time.Now()
    s.orders[id] = order
    
    return order, nil
}

// GetUserByUsername 根据用户名获取用户
func (s *MemoryStore) GetUserByUsername(username string) (models.User, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    
    for _, user := range s.users {
        if user.Username == username {
            return user, nil
        }
    }
    return models.User{}, errors.New("user not found")
}

func (s *MemoryStore) GetShopByName(name string) (models.Shop, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    sh, ok := s.shops[name]
    if !ok {
        return models.Shop{}, errors.New("shop not found")
    }
    return sh, nil
}

// GetRandomTableNumber 生成一个随机的桌号
func (s *MemoryStore) GetRandomTableNumber() string {
    // 简单生成一个1-100之间的随机桌号
    return strconv.Itoa(rand.Intn(100) + 1)
}
