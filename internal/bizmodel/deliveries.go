package bizmodel

import "time"

type Deliveries struct {
	DeliveryID     int
	OrderID        int
	DelivererID    int
	PickUpTime     *time.Time
	DeliverTime    *time.Time
	DeliveryStatus int
}
