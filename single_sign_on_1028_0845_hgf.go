// 代码生成时间: 2025-10-28 08:45:21
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// User 定义用户模型
type User struct {
    gorm.Model
    Username string `gorm:"unique;not null"`
    Password string `gorm:"not null"`
}

// SSOService 是单点登录服务
type SSOService struct {
    db *gorm.DB
}

// NewSSOService 创建SSOService实例
func NewSSOService(db *gorm.DB) *SSOService {
    return &SSOService{db: db}
}

// Login 实现用户登录逻辑
func (s *SSOService) Login(username string, password string) error {
    // 查询用户
    var user User
    result := s.db.Where("username = ? AND password = ?", username, password).First(&user)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return fmt.Errorf("user not found")
        }
        return result.Error
    }
    return nil
}

// Register 实现用户注册逻辑
func (s *SSOService) Register(username string, password string) error {
    // 检查用户名是否已存在
    var user User
    if s.db.Where("username = ?", username).First(&user).Error == nil {
        return fmt.Errorf("username already exists")
    }
    
    // 创建新用户
    err := s.db.Create(&User{Username: username, Password: password}).Error
    if err != nil {
        return err
    }
    return nil
}

func main() {
    // 连接数据库
    db, err := gorm.Open(sqlite.Open("sso.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    
    // 自动迁移模式
    db.AutoMigrate(&User{})
    
    // 创建SSOService实例
    ssoService := NewSSOService(db)
    
    // 测试登录
    if err := ssoService.Login("admin", "password"); err != nil {
        fmt.Println("Login failed: ", err)
    } else {
        fmt.Println("Login successful")
    }
    
    // 测试注册
    if err := ssoService.Register("newuser", "newpassword"); err != nil {
        fmt.Println("Register failed: ", err)
    } else {
        fmt.Println("Register successful")
    }
}