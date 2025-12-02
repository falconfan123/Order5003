package impl

import (
    "Order5003/internal/bizmodel"
    "Order5003/internal/dao"
    "errors"
)

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
