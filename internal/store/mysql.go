package store

import (
	"Order5003/internal/models"
	"database/sql"
	"encoding/json"
	"errors"
	"math/rand"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLStore struct {
	db *sql.DB
}

func NewMySQLStore(dsn string) (*MySQLStore, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &MySQLStore{db: db}, nil
}

func (s *MySQLStore) GetAllMenuItems() []models.MenuItem {
	rows, err := s.db.Query("SELECT id, name, description, price, category, is_available, created_at, updated_at FROM dishes")
	if err != nil {
		return []models.MenuItem{}
	}
	defer rows.Close()
	items := []models.MenuItem{}
	for rows.Next() {
		var item models.MenuItem
		var availableInt sql.NullInt64
		var desc sql.NullString
		if err := rows.Scan(&item.ID, &item.Name, &desc, &item.Price, &item.Category, &availableInt, &item.CreatedAt, &item.UpdatedAt); err != nil {
			continue
		}
		item.Description = desc.String
		item.IsAvailable = availableInt.Int64 == 1
		items = append(items, item)
	}
	return items
}

func (s *MySQLStore) GetMenuItemByID(id int) (models.MenuItem, error) {
	var item models.MenuItem
	var availableInt sql.NullInt64
	var desc sql.NullString
	err := s.db.QueryRow("SELECT id, name, description, price, category, is_available, created_at, updated_at FROM dishes WHERE id=?", id).
		Scan(&item.ID, &item.Name, &desc, &item.Price, &item.Category, &availableInt, &item.CreatedAt, &item.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.MenuItem{}, errors.New("menu item not found")
		}
		return models.MenuItem{}, err
	}
	item.Description = desc.String
	item.IsAvailable = availableInt.Int64 == 1
	return item, nil
}

func (s *MySQLStore) CreateMenuItem(item models.MenuItem) models.MenuItem {
	now := time.Now()
	res, err := s.db.Exec("INSERT INTO dishes (name, description, price, category, is_available, created_at, updated_at) VALUES (?,?,?,?,?, ?, ?)", item.Name, item.Description, item.Price, item.Category, boolToTinyInt(item.IsAvailable), now, now)
	if err != nil {
		return item
	}
	id, err := res.LastInsertId()
	if err == nil {
		item.ID = int(id)
	}
	item.CreatedAt = now
	item.UpdatedAt = now
	return item
}

func (s *MySQLStore) UpdateMenuItem(id int, updatedItem models.MenuItem) (models.MenuItem, error) {
	var createdAt time.Time
	err := s.db.QueryRow("SELECT created_at FROM dishes WHERE id=?", id).Scan(&createdAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.MenuItem{}, errors.New("menu item not found")
		}
		return models.MenuItem{}, err
	}
	now := time.Now()
	_, err = s.db.Exec("UPDATE dishes SET name=?, description=?, price=?, category=?, is_available=?, updated_at=? WHERE id=?",
		updatedItem.Name, updatedItem.Description, updatedItem.Price, updatedItem.Category, boolToTinyInt(updatedItem.IsAvailable), now, id)
	if err != nil {
		return models.MenuItem{}, err
	}
	updatedItem.ID = id
	updatedItem.CreatedAt = createdAt
	updatedItem.UpdatedAt = now
	return updatedItem, nil
}

func (s *MySQLStore) DeleteMenuItem(id int) error {
	res, err := s.db.Exec("DELETE FROM dishes WHERE id=?", id)
	if err != nil {
		return err
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New("menu item not found")
	}
	return nil
}

func (s *MySQLStore) CreateOrder(order models.Order) models.Order {
	now := time.Now()
	order.Status = models.OrderStatusPending
	itemsJSON, _ := json.Marshal(order.Items)
	res, err := s.db.Exec("INSERT INTO orders (table_number, items_json, total, status, created_at, updated_at) VALUES (?,?,?,?,?,?)",
		order.TableNumber, string(itemsJSON), order.Total, string(order.Status), now, now)
	if err != nil {
		order.CreatedAt = now
		order.UpdatedAt = now
		return order
	}
	id, err := res.LastInsertId()
	if err == nil {
		order.ID = int(id)
	}
	order.CreatedAt = now
	order.UpdatedAt = now
	return order
}

func (s *MySQLStore) GetOrderByID(id int) (models.Order, error) {
	var order models.Order
	var itemsStr sql.NullString
	var statusStr string
	err := s.db.QueryRow("SELECT id, table_number, items_json, total, status, created_at, updated_at FROM orders WHERE id=?", id).
		Scan(&order.ID, &order.TableNumber, &itemsStr, &order.Total, &statusStr, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Order{}, errors.New("order not found")
		}
		return models.Order{}, err
	}
	if itemsStr.Valid {
		_ = json.Unmarshal([]byte(itemsStr.String), &order.Items)
	}
	order.Status = models.OrderStatus(statusStr)
	return order, nil
}

func (s *MySQLStore) GetAllOrders() []models.Order {
	rows, err := s.db.Query("SELECT id, table_number, items_json, total, status, created_at, updated_at FROM orders")
	if err != nil {
		return []models.Order{}
	}
	defer rows.Close()
	orders := []models.Order{}
	for rows.Next() {
		var o models.Order
		var itemsStr sql.NullString
		var statusStr string
		if err := rows.Scan(&o.ID, &o.TableNumber, &itemsStr, &o.Total, &statusStr, &o.CreatedAt, &o.UpdatedAt); err != nil {
			continue
		}
		if itemsStr.Valid {
			_ = json.Unmarshal([]byte(itemsStr.String), &o.Items)
		}
		o.Status = models.OrderStatus(statusStr)
		orders = append(orders, o)
	}
	return orders
}

func (s *MySQLStore) UpdateOrderStatus(id int, status models.OrderStatus) (models.Order, error) {
	_, err := s.db.Exec("UPDATE orders SET status=?, updated_at=? WHERE id=?", string(status), time.Now(), id)
	if err != nil {
		return models.Order{}, err
	}
	return s.GetOrderByID(id)
}

func (s *MySQLStore) GetUserByUsername(username string) (models.User, error) {
	var u models.User
	var roleStr sql.NullString
	err := s.db.QueryRow("SELECT id, username, password, role, created_at, updated_at FROM users WHERE username=?", username).
		Scan(&u.ID, &u.Username, &u.Password, &roleStr, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, errors.New("user not found")
		}
		return models.User{}, err
	}
	if roleStr.Valid {
		u.Role = models.UserRole(roleStr.String)
	}
	return u, nil
}

func (s *MySQLStore) GetShopByName(name string) (models.Shop, error) {
	var sh models.Shop
	err := s.db.QueryRow("SELECT shop_id, shop_name, password, create_time FROM shops WHERE shop_name=?", name).
		Scan(&sh.ID, &sh.ShopName, &sh.Password, &sh.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Shop{}, errors.New("shop not found")
		}
		return models.Shop{}, err
	}
	sh.UpdatedAt = sh.CreatedAt
	return sh, nil
}

func (s *MySQLStore) GetRandomTableNumber() string {
	return strconv.Itoa(rand.Intn(100) + 1)
}

func boolToTinyInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
