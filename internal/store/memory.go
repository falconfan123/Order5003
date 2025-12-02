package store

import (
	"Order5003/internal/models"
	"errors"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// MemoryStore 实现内存数据库存储
type MemoryStore struct {
	mu             sync.RWMutex
	menuItems      map[int]models.Menu
	orders         map[int]models.Order
	users          map[int]models.User
	shops          map[string]models.Shop
	deliverers     map[string]models.Deliverers
	nextMenuItemID int
	nextOrderID    int
	nextUserID     int
}

// NewMemoryStore 创建一个新的内存数据库实例
func NewMemoryStore() *MemoryStore {
	store := &MemoryStore{
		menuItems:      make(map[int]models.Menu),
		orders:         make(map[int]models.Order),
		users:          make(map[int]models.User),
		shops:          make(map[string]models.Shop),
		deliverers:     make(map[string]models.Deliverers),
		nextMenuItemID: 1,
		nextOrderID:    1,
		nextUserID:     1,
	}

	return store
}

// GetAllMenuItems 获取所有菜单项目
func (s *MemoryStore) GetAllMenuItems() []models.Menu {
	s.mu.RLock()
	defer s.mu.RUnlock()

	items := make([]models.Menu, 0, len(s.menuItems))
	for _, item := range s.menuItems {
		items = append(items, item)
	}
	return items
}

// GetMenuItemByID 根据ID获取菜单项目
func (s *MemoryStore) GetMenuItemByID(id int) (models.Menu, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	item, exists := s.menuItems[id]
	if !exists {
		return models.Menu{}, errors.New("menu item not found")
	}
	return item, nil
}

// CreateMenuItem 创建新的菜单项目
func (s *MemoryStore) CreateMenuItem(item models.Menu) models.Menu {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.nextMenuItemID++

	return item
}

// UpdateMenuItem 更新菜单项目
func (s *MemoryStore) UpdateMenuItem(id int, updatedItem models.Menu) (models.Menu, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.menuItems[id]
	if !exists {
		return models.Menu{}, errors.New("menu item not found")
	}

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

func (s *MemoryStore) GetDelivererByName(name string) (models.Deliverers, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	d, ok := s.deliverers[name]
	if !ok {
		return models.Deliverers{}, errors.New("deliverer not found")
	}
	return d, nil
}

// GetRandomTableNumber 生成一个随机的桌号
func (s *MemoryStore) GetRandomTableNumber() string {
	// 简单生成一个1-100之间的随机桌号
	return strconv.Itoa(rand.Intn(100) + 1)
}
