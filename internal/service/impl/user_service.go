package impl

import (
	"Order5003/internal/bizmodel"
	"Order5003/internal/dao"
	"Order5003/internal/model"
	"errors"
)

func (s *GormStore) GetUserByUsername(username string) (bizmodel.User, error) {
	e, err := dao.GetUserByUsername(s.db, username)
	if err != nil {
		return bizmodel.User{}, errors.New("user not found")
	}
	return bizmodel.User{ID: e.UserID, Username: e.UserName, Password: e.Password}, nil
}

func (s *GormStore) GetUserByID(id int) (bizmodel.User, error) {
	e, err := dao.GetUserByID(s.db, id)
	if err != nil {
		return bizmodel.User{}, errors.New("user not found")
	}
	return bizmodel.User{ID: e.UserID, Username: e.UserName, Password: e.Password}, nil
}

func (s *GormStore) GetUsernameByUserID(userID int) (string, error) {
	e, err := dao.GetUserByID(s.db, userID)
	if err != nil {
		return "", errors.New("user not found")
	}
	return e.UserName, nil
}

func (s *GormStore) GetUserAddressByUserID(userID int) (string, error) {
	e, err := dao.GetUserAddressByUserID(s.db, userID)
	if err != nil {
		return "", errors.New("user not found")
	}
	return e.MainAddress, nil
}

func (s *GormStore) GetUserPhoneByUserID(userID int) (string, error) {
	e, err := dao.GetUserPhoneByUserID(s.db, userID)
	if err != nil {
		return "", errors.New("user not found")
	}
	return e.Phone, nil
}

// PayOrder 支付订单
func (s *GormStore) PayOrder(userID int, orderID int) error {
	//更改order表下的信息
	err := dao.UpdateOrderStatus(s.db, orderID, int(model.OrderStatusWaitingForAccept))
	if err != nil {
		return errors.New("internal server error")
	}
	return nil
}
