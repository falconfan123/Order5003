package store

import (
	"Order5003/internal/dao"
	"Order5003/internal/models"
	"errors"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormStore struct {
	db *gorm.DB
}

var _ Store = (*GormStore)(nil)

func NewGormStore(dsn string) (*GormStore, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &GormStore{db: db}, nil
}

func (s *GormStore) GetAllMenuItems() []models.Menu {
	list, err := dao.ListDishes(s.db)
	if err != nil {
		return []models.Menu{}
	}
	out := make([]models.Menu, 0, len(list))
	for _, e := range list {
		out = append(out, models.Menu{
			MenuID: e.DishID,
		})
	}
	return out
}

func (s *GormStore) GetMenuItemByID(id int) (models.Menu, error) {
	e, err := dao.GetDishByID(s.db, id)
	if err != nil {
		return models.Menu{}, errors.New("menu item not found")
	}
	return models.Menu{
		MenuID: e.DishID,
	}, nil
}

func (s *GormStore) CreateMenuItem(item models.Menu) models.Menu {
	e := &dao.DishEntity{}
	if _, err := dao.CreateDish(s.db, e); err != nil {
		return item
	}
	return item
}

func (s *GormStore) UpdateMenuItem(id int, updatedItem models.Menu) (models.Menu, error) {
	e := &dao.DishEntity{}
	if _, err := dao.UpdateDish(s.db, id, e); err != nil {
		return models.Menu{}, errors.New("menu item not found")
	}
	return updatedItem, nil
}

func (s *GormStore) DeleteMenuItem(id int) error {
	if err := dao.DeleteDish(s.db, id); err != nil {
		return errors.New("menu item not found")
	}
	return nil
}

func (s *GormStore) CreateOrder(order models.Order) models.Order {
	e := &dao.OrderEntity{}
	if _, err := dao.CreateOrder(s.db, e); err != nil {
		return order
	}
	order.ID = e.OrderID
	return order
}

func (s *GormStore) GetOrderByID(id int) (models.Order, error) {
	e, err := dao.GetOrderByID(s.db, id)
	if err != nil {
		return models.Order{}, errors.New("order not found")
	}
	var items []models.OrderItem
	return models.Order{
		ID:        e.OrderID,
		Items:     items,
		Status:    models.OrderStatus(e.Status),
		CreatedAt: e.CreatedAt,
	}, nil
}

func (s *GormStore) GetAllOrders() []models.Order {
	list, err := dao.ListOrders(s.db)
	if err != nil {
		return []models.Order{}
	}
	out := make([]models.Order, 0, len(list))
	for _, e := range list {
		var items []models.OrderItem
		out = append(out, models.Order{
			ID:        e.OrderID,
			Items:     items,
			Status:    models.OrderStatus(e.Status),
			CreatedAt: e.CreatedAt,
		})
	}
	return out
}

func (s *GormStore) UpdateOrderStatus(id int, status models.OrderStatus) (models.Order, error) {
	if err := dao.UpdateOrderStatus(s.db, id, string(status)); err != nil {
		return models.Order{}, err
	}
	return s.GetOrderByID(id)
}

func (s *GormStore) GetUserByUsername(username string) (models.User, error) {
	return models.User{}, errors.New("unsupported: username login")
}

func (s *GormStore) GetUserByID(id int) (models.User, error) {
	e, err := dao.GetUserByID(s.db, id)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}
	return models.User{
		ID:       e.UserID,
		Username: e.UserName,
		Password: e.Password,
	}, nil
}

func (s *GormStore) GetShopByName(name string) (models.Shop, error) {
	e, err := dao.GetShopByName(s.db, name)
	if err != nil {
		return models.Shop{}, errors.New("shop not found")
	}
	return models.Shop{
		ShopID:        e.ShopID,
		ShopName:      e.ShopName,
		DeliveryRange: e.DeliveryRange,
		DeliveryFee:   e.DeliveryFee,
		BusinessHours: e.BusinessHours,
		CreatedAt:     &e.CreatedAt,
		Password:      e.Password,
	}, nil
}

func (s *GormStore) GetDelivererByName(name string) (models.Deliverers, error) {
	e, err := dao.GetDelivererByName(s.db, name)
	if err != nil {
		return models.Deliverers{}, errors.New("deliverer not found")
	}
	return models.Deliverers{
		DelivererID:     e.DelivererID,
		Name:            e.Name,
		Phone:           e.Phone,
		ResponsibleArea: e.ResponsibleArea,
		Password:        e.Password,
	}, nil
}

func (s *GormStore) GetRandomTableNumber() string {
	return strconv.Itoa(1)
}
