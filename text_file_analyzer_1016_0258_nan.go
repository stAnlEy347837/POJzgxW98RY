// 代码生成时间: 2025-10-16 02:58:24
package main

import (
    "fmt"
    "os"
    "log"
    "strings"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// TextFileAnalyzer 结构体用于存储文本分析结果
type TextFileAnalyzer struct {
    ID        uint   "gorm:"primaryKey"
    Text      string
    CharCount int    "gorm:"column:char_count"
    WordCount int    "gorm:"column:word_count"
    LineCount int    "gorm:"column:line_count"
}

// DBConfig 数据库配置结构体
type DBConfig struct {
    DSN string
}

// NewDBConfig 创建一个新的数据库配置
func NewDBConfig(dsn string) DBConfig {
    return DBConfig{DSN: dsn}
}

// ConnectDB 连接数据库
func ConnectDB(cfg DBConfig) (*gorm.DB, error) {
    db, err := gorm.Open(sqlite.Open(cfg.DSN), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    // Migrate the schema
    db.AutoMigrate(&TextFileAnalyzer{})
    return db, nil
}

// AnalyzeText 分析文本文件内容
func AnalyzeText(db *gorm.DB, filePath string) error {
    file, err := os.ReadFile(filePath)
    if err != nil {
        return err
    }
    text := strings.TrimSpace(string(file))
    analyzer := TextFileAnalyzer{Text: text}
    analyzer.CharCount = len(text)
    analyzer.WordCount = strings.Count(text, " ") + 1 // +1 to count the first word
    analyzer.LineCount = strings.Count(text, "
")

    // Save analysis results to the database
    result := db.Create(&analyzer)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

// main 函数是程序的入口点
func main() {
    cfg := NewDBConfig("sqlite://text_analyzer.db")
    db, err := ConnectDB(cfg)
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }
    defer db.Migrator.Close()

    filePath := "example.txt" // Replace with the actual file path
    if err := AnalyzeText(db, filePath); err != nil {
        log.Fatalf("failed to analyze text: %v", err)
    }

    fmt.Println("Text analysis completed successfully.")
}