package impl

import (
	"Order5003/internal/bizmodel"
	"Order5003/internal/dao"
	"errors"
	"fmt"
)

func (s *GormStore) GetMenuDishesByShopID(shopID int) ([]bizmodel.Dishes, error) {
	shop, err := dao.GetShopByID(s.db, shopID) //根据商家的id取出商家信息
	if err != nil {
		return nil, fmt.Errorf("查询商家失败：%w", err)
	}
	if shop.Status == 0 {
		return nil, fmt.Errorf("该商家休息中，不可点餐")
	}
	menus, err := dao.GetMenuByShopID(s.db, shopID)
	if err != nil {
		return []bizmodel.Dishes{}, errors.New("menu by shops not found")
	}
	if len(menus) == 0 {
		return []bizmodel.Dishes{}, nil
	}
	var allDishIDs []int
	for _, menu := range menus {
		menuDishes, err := dao.GetDishesByMenuID(s.db, menu.MenuID)
		if err != nil {
			return nil, fmt.Errorf("获取菜单关联菜品失败：%w", err)
		}
		for _, md := range menuDishes {
			allDishIDs = append(allDishIDs, md.DishID)
		}
	}
	dishes, err := dao.GetDishesByIDs(s.db, allDishIDs)
	if err != nil {
		return nil, fmt.Errorf("获取菜品详情失败：%w", err)
	}
	var dishList []bizmodel.Dishes
	for _, dish := range dishes {
		dishList = append(dishList, bizmodel.Dishes{
			DishID:   dish.DishID,
			ShopID:   dish.ShopID,
			DishName: dish.DishName,
			Stock:    dish.Stock,
			Status:   int(dish.Status),
			Price:    dish.Price,
		})
	}
	return dishList, nil
}
