package impl

import (
	"Order5003/internal/bizmodel"
	"Order5003/internal/dao"
	"Order5003/internal/logger"
	"Order5003/internal/model"
	"context"
	"errors"

	"go.uber.org/zap"
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

func (s *GormStore) UpdateMenu(ctx context.Context, action string, menuID int, menuName string, status int) error {
	switch action {
	case "add":
		//新增菜单
		err := dao.CreateMenu(s.db, model.MenuEntity{
			ShopID:   menuID,
			MenuName: menuName,
			Status:   int8(status),
		})
		if err != nil {
			return err
		}
	case "delete":
		//删除菜单
		err := dao.DeleteMenu(s.db, menuID)
		if err != nil {
			return err
		}
	case "update":
		logger.Info("更新菜单", zap.Int("menuID", menuID), zap.String("menuName", menuName), zap.Int("status", status))
		//更新菜单
		err := dao.UpdateMenu(s.db, menuID, menuName, int8(status))
		if err != nil {
			return err
		}
	default:
		return errors.New("invalid action")
	}
	return nil
}
