package impl

import (
	"Order5003/internal/bizmodel"
	"Order5003/internal/dao"
	"context"
	"errors"
)

func (s *GormStore) GetMenuByShopID(ctx context.Context, shopID int) ([]bizmodel.Menu, error) {
	list, err := dao.GetMenuByShopID(s.db, shopID)
	if err != nil {
		return []bizmodel.Menu{}, errors.New("menu not found")
	}
	out := make([]bizmodel.Menu, 0, len(list))
	for _, e := range list {
		out = append(out,
			bizmodel.Menu{
				MenuID:     e.MenuID,
				ShopID:     e.ShopID,
				MenuName:   e.MenuName,
				Status:     int(e.Status),
				CreateTime: e.CreateTime,
				UpdateTime: e.UpdateTime,
			})
	}
	return out, nil
}

func (s *GormStore) GetDishesByMenuID(ctx context.Context, menuID int) ([]bizmodel.Dishes, error) {
	//先从menu_dishes取dishesid
	Dishes_ids, err := dao.GetDishesIDByMenuID(s.db, menuID)
	if err != nil {
		return []bizmodel.Dishes{}, errors.New("dishes not found")
	}
	//再从Dishes_ids的每一项取出对应的Dish详情
	out := make([]bizmodel.Dishes, 0, len(Dishes_ids))
	for _, dishID := range Dishes_ids {
		dish, err := dao.GetBothDishByDishID(ctx, s.db, dishID)
		if err != nil {
			return []bizmodel.Dishes{}, errors.New("dishes not found")
		}
		out = append(out, *dish)
	}
	return out, nil
}
