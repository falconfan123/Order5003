package main

import (
	"Order5003/internal/handlers"
	"Order5003/internal/store"
	"log"
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
	log.Println("MySQL DSN:", dsn)
	if dsn == "" {
		log.Fatal("未配置 MySQL DSN，无法启动")
	}
	gs, err := store.NewGormStore(dsn)
	if err != nil {
		log.Fatal("连接 MySQL 失败:", err)
	}
	s = gs
	log.Println("使用 MySQL(GORM) 数据库")

	// 创建处理器实例
	menuHandler := handlers.NewMenuHandler(s)
	orderHandler := handlers.NewOrderHandler(s)
	userHandler := handlers.NewUserHandler(s)
	shopHandler := handlers.NewShopHandler(s)
	deliverHandler := handlers.NewDeliverHandler(s)

	r := gin.Default()

	userG := r.Group("/user")
	userG.GET("", func(c *gin.Context) { c.File("web/templates/customer.html") })
	userG.POST("/login", func(c *gin.Context) { userHandler.LoginUser(c.Writer, c.Request) })

	shopG := r.Group("/shop")
	shopG.GET("", func(c *gin.Context) { c.File("web/templates/merchant.html") })
	shopG.POST("/login", func(c *gin.Context) { shopHandler.Login(c.Writer, c.Request) })

	deliverG := r.Group("/deliver")
	deliverG.GET("", func(c *gin.Context) { c.File("web/templates/deliverer.html") })
	deliverG.POST("/login", func(c *gin.Context) { deliverHandler.Login(c.Writer, c.Request) })

	r.GET("/api/menu", func(c *gin.Context) { menuHandler.GetAllMenuItems(c.Writer, c.Request) })
	r.Any("/api/menu-item", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			menuHandler.GetMenuItemByID(c.Writer, c.Request)
		case "POST":
			menuHandler.CreateMenuItem(c.Writer, c.Request)
		case "PUT":
			menuHandler.UpdateMenuItem(c.Writer, c.Request)
		case "DELETE":
			menuHandler.DeleteMenuItem(c.Writer, c.Request)
		default:
			c.Status(405)
		}
	})

	r.Any("/api/orders", func(c *gin.Context) {
		switch c.Request.Method {
		case "POST":
			orderHandler.CreateOrder(c.Writer, c.Request)
		default:
			c.Status(405)
		}
	})
	r.GET("/api/orders/all", func(c *gin.Context) { orderHandler.GetAllOrders(c.Writer, c.Request) })
	r.PUT("/api/orders/status", func(c *gin.Context) { orderHandler.UpdateOrderStatus(c.Writer, c.Request) })
	r.GET("/api/table/number", func(c *gin.Context) { orderHandler.GetRandomTableNumber(c.Writer, c.Request) })

	r.Static("/static", "web/static")

	log.Println("服务器启动在 http://localhost:8080")
	log.Println("用户端: http://localhost:8080/user")
	log.Println("商家端: http://localhost:8080/shop")
	log.Println("配送员端: http://localhost:8080/deliver")
	r.Run(":8080")
}
