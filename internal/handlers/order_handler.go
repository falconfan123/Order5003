package handlers

import (
	"Order5003/internal/bizmodel"
	"Order5003/internal/logger"
	"Order5003/internal/service"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type OrderHandler struct {
	svc service.OrderService
}

func NewOrderHandler(svc service.OrderService) *OrderHandler {
	return &OrderHandler{svc: svc}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	var req bizmodel.NewOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效请求体: " + err.Error()})
		return
	}
	logger.Info("Receive CreateOrder request: %v", req)
	if req.UserID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户ID不合法（必须大于0）"})
		return
	}
	// 3.2 校验菜品列表非空、数量合法
	if len(req.Dishes) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "订单不能为空，至少选择一道菜品"})
		return
	}
	for _, item := range req.Dishes {
		if item.Quantity < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "菜品数量不能小于1"})
			return
		}
		if item.DishId <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "菜品ID不合法（必须大于0）"})
			return
		}
		if item.Price.IsNegative() {
			c.JSON(http.StatusBadRequest, gin.H{"error": "菜品单价不能为负数"})
			return
		}
	}
	var createdOrderID int

	err := h.svc.WithTransaction(c.Request.Context(), func(tx *gorm.DB) error {
		logger.Info("CreateOrder transaction started")
		//1 批量查询菜品信息（获取名称、单价、所属商家，用于冗余存储和校验）
		dishMap := make(map[int]*bizmodel.Dishes)
		var targetShopID int
		for _, item := range req.Dishes {
			if _, exist := dishMap[item.DishId]; exist {
				continue
			}
			logger.Info("Query dish ID: %d", item.DishId)
			//从dishes表中查询菜品详情
			dish, err := h.svc.GetDishByID(c.Request.Context(), tx, item.DishId)
			if err != nil {
				logger.Error("GetDishByID failed, dishID: %d, err: %v", item.DishId, err) // 新增：打印查询错误
				return fmt.Errorf("查询菜品ID %d 失败: %w", item.DishId, err)
			}
			if dish == nil {
				logger.Error("Dish not found, dishID: %d", item.DishId) // 新增：打印菜品不存在
				return fmt.Errorf("菜品ID %d 不存在", item.DishId)
			}
			dishMap[item.DishId] = dish
			if targetShopID == 0 {
				targetShopID = dish.ShopID
			} else if targetShopID != dish.ShopID {
				logger.Error("Cross shop error, targetShopID: %d, dishShopID: %d", targetShopID, dish.ShopID) // 新增：打印跨商家
				return fmt.Errorf("订单中包含来自不同商家的菜品，无法创建订单")
			}
		}
		logger.Info("Query dish map: %v", dishMap)
		//2 计算总金额 + 组装订单明细（修正字段名一致性）
		var totalAmount decimal.Decimal
		var orderDishDetails []bizmodel.OrderDishDetail
		for _, item := range req.Dishes {
			dish := dishMap[item.DishId]
			quantityDec := decimal.NewFromInt(int64(item.Quantity))
			subtotal := dish.Price.Mul(quantityDec) // 无报错
			totalAmount = totalAmount.Add(subtotal)
			orderDishDetails = append(orderDishDetails, bizmodel.OrderDishDetail{
				DishID:    item.DishId,
				DishName:  dish.DishName,
				Quantity:  item.Quantity,
				UnitPrice: dish.Price,
				Subtotal:  subtotal,
			})
		}
		logger.Info("Total amount: %v", totalAmount)
		//3 创建订单主表
		orderMaster := &bizmodel.Order{
			UserID:      req.UserID,
			ShopID:      targetShopID,
			TotalAmount: totalAmount,
			Status:      int(bizmodel.OrderStatusPending),
			CreatedAt:   time.Now(),
		}
		logger.Info("Create order master: %+v", orderMaster) // 改%v为%+v：打印所有字段
		orderID, err := h.svc.CreateOrderMaster(c.Request.Context(), tx, orderMaster)
		if err != nil {
			return fmt.Errorf("创建订单主表失败：%v", err)
		}
		createdOrderID = orderID // 存储订单ID，用于响应
		logger.Info("Created order ID: %d", createdOrderID)
		//4 创建订单明细表
		for i := range orderDishDetails {
			orderDishDetails[i].OrderID = orderID // 绑定订单ID
			if err := h.svc.CreateOrderDish(c.Request.Context(), tx, &orderDishDetails[i]); err != nil {
				return fmt.Errorf("创建订单明细失败：%v", err)
			}
		}
		return nil
	})
	//5 事务处理结果
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建订单失败: " + err.Error()})
		return
	}
	//6 成功响应
	c.JSON(http.StatusCreated, gin.H{"order_id": createdOrderID, "message": "订单创建成功"})
}

func (h *OrderHandler) GetAllOrders(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	orderIDStr := c.Query("user_id")
	orderID, err := strconv.Atoi(orderIDStr)
	logger.Info("GetAllOrders, userID: %d", orderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}
	order, err := h.svc.GetOrderByUserID(orderID)
	logger.Info("GetOrderByUserID, order: %+v", order)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}
