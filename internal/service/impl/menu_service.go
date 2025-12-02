package impl

import (
    "Order5003/internal/bizmodel"
    "Order5003/internal/dao"
    "Order5003/internal/model"
    "errors"
)

func (s *GormStore) GetAllMenuItems() []bizmodel.Menu {
    list, err := dao.ListDishes(s.db)
    if err != nil {
        return []bizmodel.Menu{}
    }
    out := make([]bizmodel.Menu, 0, len(list))
    for _, e := range list {
        out = append(out, bizmodel.Menu{MenuID: e.DishID})
    }
    return out
}

func (s *GormStore) GetMenuItemByID(id int) (bizmodel.Menu, error) {
    e, err := dao.GetDishByID(s.db, id)
    if err != nil {
        return bizmodel.Menu{}, errors.New("menu item not found")
    }
    return bizmodel.Menu{MenuID: e.DishID}, nil
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
