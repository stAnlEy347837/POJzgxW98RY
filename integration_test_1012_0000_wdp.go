// 代码生成时间: 2025-10-12 00:00:31
package main

import (
    "fmt"
    "os"
    "testing"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Product 模型
type Product struct {
    gorm.Model
    Code  string
    Price uint
}

// setupDB 初始化数据库
func setupDB() *gorm.DB {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // 迁移 schema
    db.AutoMigrate(&Product{})

    return db
}

// teardownDB 清理数据库
func teardownDB(db *gorm.DB) {
    db.Migrator().DropTable(&Product{})
    os.Remove("test.db")
}

// TestProductCreate 测试创建产品
func TestProductCreate(t *testing.T) {
    db := setupDB()
    defer teardownDB(db)

    // 创建产品
    product := Product{Code: "D42", Price: 100}
    result := db.Create(&product)

    // 检查是否有错误和产品是否创建
    if result.Error != nil {
        t.Errorf("expected no error, got %v", result.Error)
    }
    if product.ID == 0 {
        t.Errorf("expected product with ID, got %v", product.ID)
    }
}

func main() {
    // 运行测试
    fmt.Println("Running tests...")
    testing.Main()
}
