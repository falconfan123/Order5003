package handlers

import (
    "Order5003/internal/logger"
    "Order5003/internal/service"
    "encoding/json"
    "fmt"
    "net/http"
)

type MenuDishHandler struct {
    shopSvc service.ShopService
}

func NewMenuDishHandler(s service.ShopService) *MenuDishHandler {
    return &MenuDishHandler{shopSvc: s}
}

type GetMenuDishesReq struct {
	ShopID int `json:"shop_id"` // 字段名要和前端 JSON 的 key 一致（shop_id）
}

func (h *MenuDishHandler) GetAllMenuDishesByShopID(w http.ResponseWriter, r *http.Request) {
	var req GetMenuDishesReq
	// 解析 JSON Body 到结构体
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("参数格式错误，请传 JSON 格式的 shop_id"))
		return
	}

	shopID := req.ShopID
	logger.Info("GetAllMenuDishesByShopID", ":shop_id", shopID)

	// 调用 service 层（核心逻辑全在 service，handler 不碰关联查询）
    menuDishes, err := h.shopSvc.GetAllMenuDishesByShopID(shopID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("查询失败：%s", err.Error())))
		return
	}

	// 3. 返回合并后的数据给前端
	w.Header().Set("Content-Type", "application/json") // 必须设置 JSON 响应头
	w.WriteHeader(http.StatusOK)
	// 用 json.Marshal 序列化结构体，避免直接拼接字符串报错
	logger.Info("GetAllMenuDishesByShopID", ":menu_dishes", menuDishes)
	resp, _ := json.Marshal(map[string]interface{}{
		"code": 200,
		"data": menuDishes,
	})
	w.Write(resp)
}
