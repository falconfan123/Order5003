package main

import (
	"Order5003/internal/handlers"
	"Order5003/internal/store"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	var s store.Store
	var dsn string
	data, err := os.ReadFile("config/mysql_dsn.txt")
	if err == nil {
		dsn = strings.TrimSpace(string(data))
	}
	log.Println("MySQL DSN:", dsn)
	if dsn != "" {
		ms, err := store.NewMySQLStore(dsn)
		if err == nil {
			s = ms
			log.Println("使用 MySQL 数据库")
		} else {
			log.Println("连接 MySQL 失败，回退到内存存储:", err)
			s = store.NewMemoryStore()
		}
	} else {
		s = store.NewMemoryStore()
		log.Println("未配置 MySQL，使用内存存储")
	}

	// 创建处理器实例
	menuHandler := handlers.NewMenuHandler(s)
	orderHandler := handlers.NewOrderHandler(s)
	userHandler := handlers.NewUserHandler(s)

	// 设置路由
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/templates/customer.html")
	})

	http.HandleFunc("/merchant", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/templates/merchant.html")
	})

	// API 路由
	http.HandleFunc("/api/menu", menuHandler.GetAllMenuItems)
	http.HandleFunc("/api/menu-item", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			menuHandler.GetMenuItemByID(w, r)
		case http.MethodPost:
			menuHandler.CreateMenuItem(w, r)
		case http.MethodPut:
			menuHandler.UpdateMenuItem(w, r)
		case http.MethodDelete:
			menuHandler.DeleteMenuItem(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/orders", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			orderHandler.CreateOrder(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/orders/all", orderHandler.GetAllOrders)

	http.HandleFunc("/api/orders/status", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			orderHandler.UpdateOrderStatus(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/table/number", orderHandler.GetRandomTableNumber)
	http.HandleFunc("/api/login", userHandler.Login)

	// 静态文件服务
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// 启动服务器
	log.Println("服务器启动在 http://localhost:8080")
	log.Println("顾客端: http://localhost:8080")
	log.Println("商家端: http://localhost:8080/merchant")
	log.Println("商家端默认账号: admin, 密码: admin")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
