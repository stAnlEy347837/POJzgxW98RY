// 代码生成时间: 2025-09-23 22:21:37
package main

import (
    "fmt"
    "math"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// MathTool 是一个包含数学计算工具集的结构体
type MathTool struct {
    db *gorm.DB
}

// NewMathTool 初始化并返回一个 *MathTool 实例
func NewMathTool(db *gorm.DB) *MathTool {
    return &MathTool{db: db}
}

// Add 实现加法运算
func (mt *MathTool) Add(a, b float64) (float64, error) {
    if mt.db != nil {
        // 可选：将操作记录到数据库
        err := mt.db.Exec("INSERT INTO math_operations (operation, a, b, result) VALUES ('ADD', ?, ?, ?)", a, b, a+b).Error
        if err != nil {
            return 0, err
        }
    }
    return a + b, nil
}

// Subtract 实现减法运算
func (mt *MathTool) Subtract(a, b float64) (float64, error) {
    if mt.db != nil {
        // 可选：将操作记录到数据库
        err := mt.db.Exec("INSERT INTO math_operations (operation, a, b, result) VALUES ('SUBTRACT', ?, ?, ?)", a, b, a-b).Error
        if err != nil {
            return 0, err
        }
    }
    return a - b, nil
}

// Multiply 实现乘法运算
func (mt *MathTool) Multiply(a, b float64) (float64, error) {
    if mt.db != nil {
        // 可选：将操作记录到数据库
        err := mt.db.Exec("INSERT INTO math_operations (operation, a, b, result) VALUES ('MULTIPLY', ?, ?, ?)", a, b, a*b).Error
        if err != nil {
            return 0, err
        }
    }
    return a * b, nil
}

// Divide 实现除法运算，并处理除以零的情况
func (mt *MathTool) Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    if mt.db != nil {
        // 可选：将操作记录到数据库
        err := mt.db.Exec("INSERT INTO math_operations (operation, a, b, result) VALUES ('DIVIDE', ?, ?, ?)", a, b, a/b).Error
        if err != nil {
            return 0, err
        }
    }
    return a / b, nil
}

// SquareRoot 计算平方根，并处理负数的情况
func (mt *MathTool) SquareRoot(a float64) (float64, error) {
    if a < 0 {
        return 0, fmt.Errorf("cannot calculate square root of negative number")
    }
    return math.Sqrt(a), nil
}

func main() {
    // 初始化数据库连接
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()

    // 自动迁移模式
    db.AutoMigrate(&MathOperation{})

    // 创建 MathTool 实例
    mathTool := NewMathTool(db)

    // 执行数学运算
    result, err := mathTool.Add(10, 5)
    if err != nil {
        fmt.Println("Error: ", err)
    } else {
        fmt.Printf("Add: 10 + 5 = %v
", result)
    }

    result, err = mathTool.Subtract(10, 5)
    if err != nil {
        fmt.Println("Error: ", err)
    } else {
        fmt.Printf("Subtract: 10 - 5 = %v
", result)
    }

    result, err = mathTool.Multiply(10, 5)
    if err != nil {
        fmt.Println("Error: ", err)
    } else {
        fmt.Printf("Multiply: 10 * 5 = %v
", result)
    }

    result, err = mathTool.Divide(10, 5)
    if err != nil {
        fmt.Println("Error: ", err)
    } else {
        fmt.Printf("Divide: 10 / 5 = %v
", result)
    }

    result, err = mathTool.SquareRoot(25)
    if err != nil {
        fmt.Println("Error: ", err)
    } else {
        fmt.Printf("SquareRoot: sqrt(25) = %v
", result)
    }
}

// MathOperation 用于存储数据库中的数学操作记录
type MathOperation struct {
    ID        uint   "gorm:column:id;primaryKey"
n    Operation string
    A, B      float64
    Result    float64
}
