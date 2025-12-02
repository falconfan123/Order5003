package models

import "time"

// UserRole 表示用户角色类型
type UserRole string

const (
    UserRoleCustomer UserRole = "customer" // 顾客
    UserRoleAdmin    UserRole = "admin"    // 管理员
    UserRoleStaff    UserRole = "staff"    // 员工
)

// User 表示系统中的用户
type User struct {
    ID        int       `json:"id"`
    Username  string    `json:"username"`
    Password  string    `json:"-"` // 不导出密码字段
    Role      UserRole  `json:"role"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}