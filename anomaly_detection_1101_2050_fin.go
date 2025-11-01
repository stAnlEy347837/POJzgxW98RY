// 代码生成时间: 2025-11-01 20:50:23
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// AnomalyRecord represents a record for anomaly detection
type AnomalyRecord struct {
    gorm.Model
    Value float64 `gorm:"column:value"`
}

// AnomalyDetector is the main struct for anomaly detection
type AnomalyDetector struct {
    DB *gorm.DB
}

// NewAnomalyDetector creates a new AnomalyDetector with a database connection
func NewAnomalyDetector() (*AnomalyDetector, error) {
    db, err := gorm.Open(sqlite.Open("anomaly.db"), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    
    // Migrate the schema
    db.AutoMigrate(&AnomalyRecord{})
    
    return &AnomalyDetector{DB: db}, nil
}

// DetectAnomaly checks if the given value is an anomaly based on the historical data
func (ad *AnomalyDetector) DetectAnomaly(value float64) (bool, error) {
    // Retrieve historical data
    var records []AnomalyRecord
    if err := ad.DB.Find(&records).Error; err != nil {
        return false, err
    }
    
    // Calculate the mean and standard deviation of historical data
    sum := 0.0
    for _, record := range records {
        sum += record.Value
    }
    mean := sum / float64(len(records))
    
    var sumSqDiff float64
    for _, record := range records {
        sumSqDiff += (record.Value - mean) * (record.Value - mean)
    }
    stdDev := math.Sqrt(sumSqDiff / float64(len(records)-1))
    
    // Define a threshold for anomaly detection (e.g., 3 standard deviations)
    threshold := 3.0 * stdDev
    
    // Check if the given value is an anomaly
    if math.Abs(value - mean) > threshold {
        return true, nil
    }
    
    return false, nil
}

func main() {
    // Create a new anomaly detector
    detector, err := NewAnomalyDetector()
    if err != nil {
        fmt.Printf("Error creating anomaly detector: %v
", err)
        return
    }
    
    // Example usage of DetectAnomaly
    value := 10.5
    isAnomaly, err := detector.DetectAnomaly(value)
    if err != nil {
        fmt.Printf("Error detecting anomaly: %v
", err)
        return
    }
    
    if isAnomaly {
        fmt.Printf("Value %v is an anomaly
", value)
    } else {
        fmt.Printf("Value %v is not an anomaly
", value)
    }
}
