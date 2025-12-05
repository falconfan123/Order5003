package impl

import (
	"Order5003/internal/bizmodel"
	"Order5003/internal/dao"
	"Order5003/internal/logger"
	"errors"

	"go.uber.org/zap"
)

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

func (s *GormStore) GetAllShops() ([]bizmodel.Shop, error) {
	list, err := dao.ListShops(s.db)
	logger.Info("list shops", zap.Any("list", list))
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
			Type:          int(e.Type),
		})
	}
	return out, nil
}

func (s *GormStore) GetShopNameByShopID(shopID int) (string, error) {
	e, err := dao.GetShopByID(s.db, shopID)
	if err != nil {
		return "", errors.New("shop not found")
	}
	return e.ShopName, nil
}
