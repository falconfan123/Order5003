package store

import (
	"Order5003/internal/bizmodel"
	"Order5003/internal/dao"
	"Order5003/internal/model"
	"errors"

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

func (s *GormStore) GetAllMenuItems() []bizmodel.Menu {
	list, err := dao.ListDishes(s.db)
	if err != nil {
		return []bizmodel.Menu{}
	}
	out := make([]bizmodel.Menu, 0, len(list))
	for _, e := range list {
		out = append(out, bizmodel.Menu{
			MenuID: e.DishID,
		})
	}
	return out
}

func (s *GormStore) GetMenuItemByID(id int) (bizmodel.Menu, error) {
	e, err := dao.GetDishByID(s.db, id)
	if err != nil {
		return bizmodel.Menu{}, errors.New("menu item not found")
	}
	return bizmodel.Menu{
		MenuID: e.DishID,
	}, nil
}

func (s *GormStore) CreateMenuItem(item bizmodel.Menu) bizmodel.Menu {
	e := &model.DishEntity{}
	if _, err := dao.CreateDish(s.db, e); err != nil {
		return item
	}
	return item
}

func (s *GormStore) UpdateMenuItem(id int, updatedItem bizmodel.Menu) (bizmodel.Menu, error) {
	e := &model.DishEntity{}
	if _, err := dao.UpdateDish(s.db, id, e); err != nil {
		return bizmodel.Menu{}, errors.New("menu item not found")
	}
	return updatedItem, nil
}

func (s *GormStore) DeleteMenuItem(id int) error {
	if err := dao.DeleteDish(s.db, id); err != nil {
		return errors.New("menu item not found")
	}
	return nil
}

func (s *GormStore) CreateOrder(order bizmodel.Order) bizmodel.Order {
	e := &model.OrderEntity{}
	if _, err := dao.CreateOrder(s.db, e); err != nil {
		return order
	}
	order.ID = e.OrderID
	return order
}

func (s *GormStore) GetOrderByID(id int) (bizmodel.Order, error) {
	e, err := dao.GetOrderByID(s.db, id)
	if err != nil {
		return bizmodel.Order{}, errors.New("order not found")
	}
	var items []bizmodel.OrderItem
	return bizmodel.Order{
		ID:        e.OrderID,
		Items:     items,
		Status:    bizmodel.OrderStatus(e.Status),
		CreatedAt: e.CreatedAt,
	}, nil
}

func (s *GormStore) GetAllOrders() []bizmodel.Order {
	list, err := dao.ListOrders(s.db)
	if err != nil {
		return []bizmodel.Order{}
	}
	out := make([]bizmodel.Order, 0, len(list))
	for _, e := range list {
		var items []bizmodel.OrderItem
		out = append(out, bizmodel.Order{
			ID:        e.OrderID,
			Items:     items,
			Status:    bizmodel.OrderStatus(e.Status),
			CreatedAt: e.CreatedAt,
		})
	}
	return out
}

func (s *GormStore) UpdateOrderStatus(id int, status bizmodel.OrderStatus) (bizmodel.Order, error) {
	if err := dao.UpdateOrderStatus(s.db, id, string(status)); err != nil {
		return bizmodel.Order{}, err
	}
	return s.GetOrderByID(id)
}

func (s *GormStore) GetUserByUsername(username string) (bizmodel.User, error) {
	e, err := dao.GetUserByUsername(s.db, username)
	if err != nil {
		return bizmodel.User{}, errors.New("user not found")
	}
	return bizmodel.User{
		ID:       e.UserID,
		Username: e.UserName,
		Password: e.Password,
	}, nil
}

func (s *GormStore) GetUserByID(id int) (bizmodel.User, error) {
	e, err := dao.GetUserByID(s.db, id)
	if err != nil {
		return bizmodel.User{}, errors.New("user not found")
	}
	return bizmodel.User{
		ID:       e.UserID,
		Username: e.UserName,
		Password: e.Password,
	}, nil
}

func (s *GormStore) GetShopByName(name string) (bizmodel.Shop, error) {
	e, err := dao.GetShopByName(s.db, name)
	if err != nil {
		return bizmodel.Shop{}, errors.New("shop not found")
	}
	return bizmodel.Shop{
		ShopID:        e.ShopID,
		ShopName:      e.ShopName,
		DeliveryRange: e.DeliveryRange,
		DeliveryFee:   e.DeliveryFee,
		BusinessHours: e.BusinessHours,
		CreatedAt:     &e.CreatedAt,
		Password:      e.Password,
	}, nil
}

func (s *GormStore) GetDelivererByName(name string) (bizmodel.Deliverers, error) {
	e, err := dao.GetDelivererByName(s.db, name)
	if err != nil {
		return bizmodel.Deliverers{}, errors.New("deliverer not found")
	}
	return bizmodel.Deliverers{
		DelivererID:     e.DelivererID,
		Name:            e.Name,
		Phone:           e.Phone,
		ResponsibleArea: e.ResponsibleArea,
		Password:        e.Password,
	}, nil
}

func (s *GormStore) GetAllShops() ([]bizmodel.Shop, error) {
	list, err := dao.ListShops(s.db)
	if err != nil {
		return []bizmodel.Shop{}, errors.New("shops not found")
	}
	out := make([]bizmodel.Shop, 0, len(list))
	for _, e := range list {
		out = append(out, bizmodel.Shop{
			ShopID:        e.ShopID,
			ShopName:      e.ShopName,
			DeliveryRange: e.DeliveryRange,
			DeliveryFee:   e.DeliveryFee,
			BusinessHours: e.BusinessHours,
			CreatedAt:     &e.CreatedAt,
			Password:      e.Password,
			Status:        int(e.Status),
		})
	}
	return out, nil
}
