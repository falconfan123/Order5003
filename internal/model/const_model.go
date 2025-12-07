package model

type OrderStatus int

const (
	OrderStatusPending            OrderStatus = 1
	OrderStatusPreparing          OrderStatus = 8
	OrderStatusDelivering         OrderStatus = 5
	OrderStatusCompleted          OrderStatus = 6
	OrderStatusSelfCancelled      OrderStatus = 3
	OrderStatusShopCancelled      OrderStatus = 7
	OrderStatusWaitingForAccept   OrderStatus = 2
	OrderStatusWaitingForDelivery OrderStatus = 4
)
