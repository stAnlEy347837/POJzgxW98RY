// 代码生成时间: 2025-10-20 20:57:20
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// SensorData represents the data from an IoT sensor.
type SensorData struct {
    gorm.Model
    SensorID   string `gorm:"primaryKey"` // Unique ID for the sensor
    FieldData string
    Timestamp int64  // Unix timestamp
}

// DBConfig holds the database connection configuration.
type DBConfig struct {
    DSN string
}

// Database represents the database connection.
var db *gorm.DB

func main() {
    // Initialize database configuration
    dbConfig := DBConfig{
        DSN: "file:agriculture_iot.db?cache=shared&mode=rwc",
    }

    // Connect to the database
    db, err := gorm.Open(sqlite.Open(dbConfig.DSN), &gorm.Config{})
    if err != nil {
        fmt.Printf("Error connecting to the database: %v
", err)
        return
    }
    defer db.Close()

    // Migrate the database schema
    if err := db.AutoMigrate(&SensorData{}); err != nil {
        fmt.Printf("Error migrating database schema: %v
", err)
        return
    }

    // Example of adding a new sensor data entry
    sampleData := SensorData{
        SensorID:   "sensor_1",
        FieldData: "temperature: 25°C, humidity: 60%",
        Timestamp: time.Now().Unix(),
    }

    if err := db.Create(&sampleData).Error; err != nil {
        fmt.Printf("Error creating sensor data: %v
", err)
        return
    }

    fmt.Println("Sensor data added successfully")
}

// Note: This example uses SQLite for simplicity, but you can replace it with any other GORM-supported database.
// Ensure to import the appropriate driver and update the DSN accordingly.
