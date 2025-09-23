// 代码生成时间: 2025-09-24 04:15:09
package main

import (
    "fmt"
    "log"
    "os"
    "time"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite" // Import the SQLite driver
)

// SystemMetrics holds the metrics to be stored in the database
type SystemMetrics struct {
    ID        uint      `gorm:"primary_key"`
    CreatedAt time.Time
    CPUUsage  float64
    MemoryUsed uint64
    DiskUsage uint64
}

// Database is a struct to hold database connection
type Database struct {
    DB *gorm.DB
}

// NewDatabase creates a new database connection
func NewDatabase() *Database {
    db, err := gorm.Open("sqlite3:performance.db", &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }

    // Migrate the schema
    db.AutoMigrate(&SystemMetrics{})

    return &Database{DB: db}
}

// RecordMetrics records the system metrics into the database
func (d *Database) RecordMetrics(cpu float64, memory uint64, disk uint64) error {
    metrics := SystemMetrics{
        CreatedAt: time.Now(),
        CPUUsage:  cpu,
        MemoryUsed: memory,
        DiskUsage: disk,
    }

    if err := d.DB.Create(&metrics).Error; err != nil {
        return err
    }

    return nil
}

func main() {
    // Create a new database
    db := NewDatabase()
    defer db.DB.Close()

    // Define the interval for recording metrics
    interval := 5 * time.Second

    ticker := time.NewTicker(interval)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            // Here you would implement the actual system performance metrics collection
            // For demonstration purposes, we use mock values
            cpu := 50.0 // CPU usage percentage
            memory := 1024 * 1024 * 512 // Memory used in bytes
            disk := 1024 * 1024 * 1024 * 100 // Disk used in bytes

            // Record the mock metrics into the database
            if err := db.RecordMetrics(cpu, memory, disk); err != nil {
                log.Printf("Failed to record metrics: %s
", err)
            }
        }
    }
}
