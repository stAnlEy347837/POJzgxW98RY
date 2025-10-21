// 代码生成时间: 2025-10-22 07:32:12
// slow_query_analyzer.go

package main

import (
    "fmt"
    "log"
    "os"
    "time"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "gorm.io/gorm/schema"
)

// SlowQuery is a struct to store slow query data
type SlowQuery struct {
    ID       uint64    `gorm:"primary_key"`
    SQL      string    `gorm:""`
    Duration time.Duration
    CreatedAt time.Time
}

// DBConfig is a struct to store database configuration
type DBConfig struct {
    Username string
    Password string
    Host     string
    Port     int
    DBName   string
}

// NewGormClient creates a new GORM client with logging
func NewGormClient(cfg DBConfig) (*gorm.DB, error) {
   dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
      cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
     Logger: logger.Default.LogMode(logger.Info),
  })
  if err != nil {
     return nil, err
  }
  return db, nil
}

// AnalyzeSlowQueries performs slow query analysis
func AnalyzeSlowQueries(db *gorm.DB, threshold time.Duration) error {
  var slowQueries []SlowQuery
  results := db.Raw("SHOW PROFILES").Scan(&slowQueries)
  if results.Error != nil {
    return results.Error
  }
  for _, q := range slowQueries {
    if q.Duration > threshold {
      fmt.Printf("Slow query: "%s" took %s
", q.SQL, q.Duration)
    }
  }
  return nil
}

// main function to run the slow query analyzer
func main() {
  cfg := DBConfig{
    Username: "user",
    Password: "password",
    Host:     "localhost",
    Port:     3306,
    DBName:   "database",
  }
  db, err := NewGormClient(cfg)
  if err != nil {
    log.Fatalf("Failed to connect to database: %v", err)
  }
  defer db.Close()
  
  threshold := 100 * time.Millisecond
  if err := AnalyzeSlowQueries(db, threshold); err != nil {
    log.Printf("Error analyzing slow queries: %v", err)
  }
}
