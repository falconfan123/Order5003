package main

import (
	"Order5003/internal/api"
	"Order5003/internal/handlers"
	"Order5003/internal/logger"
	"Order5003/internal/store"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	var s store.Store
	var dsn string
	data, err := os.ReadFile("config/mysql_dsn.txt")
	if err == nil {
		dsn = strings.TrimSpace(string(data))
	}
	_ = logger.Init()
	logger.Info("MySQL DSN:", dsn)
	if dsn == "" {
		logger.Error("未配置 MySQL DSN，无法启动")
		return
	}
	gs, err := store.NewGormStore(dsn)
	if err != nil {
		logger.Error("连接 MySQL 失败:", err)
		return
	}
	s = gs
	logger.Info("使用 MySQL(GORM) 数据库")

	appHandlers := handlers.NewAppHandlers(s)

	r := gin.Default()

	api.RegisterAllRoutes(r, appHandlers)

	r.Static("/static", "web/static")

	logger.Info("服务器启动在 http://localhost:8080")
	logger.Info("用户端: http://localhost:8080/user")
	logger.Info("商家端: http://localhost:8080/shop")
	logger.Info("配送员端: http://localhost:8080/deliver")
	r.Run(":8080")
}
