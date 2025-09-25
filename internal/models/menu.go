package models

import "time"

// MenuItem 表示菜单中的一个菜品
type MenuItem struct {
    ID          int       `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Price       float64   `json:"price"`
    Category    string    `json:"category"`
    IsAvailable bool      `json:"is_available"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

// Menu 表示完整的菜单
type Menu struct {
    Items []MenuItem `json:"items"`
}