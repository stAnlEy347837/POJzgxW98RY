// 代码生成时间: 2025-10-30 15:51:10
 * It uses GORM to interact with the database and provides CRUD operations for rehabilitation exercises.
 */

package main

import (
    "fmt"
    "log"
    "gorm.io/driver/sqlite" // Assuming SQLite for simplicity
    "gorm.io/gorm"
)

// Exercise represents the model for a rehabilitation exercise
type Exercise struct {
    gorm.Model
    Name        string  `gorm:"column:name;size:255"` // Exercise name
    Description string  `gorm:"column:description;type:text"` // Exercise description
    Duration    float64 `gorm:"column:duration;type:float"` // Duration in minutes
}

// DBClient is a type that wraps the *gorm.DB connection
type DBClient struct {
    DB *gorm.DB
}

// NewDBClient initializes a new DBClient with a connection to the database
func NewDBClient() (*DBClient, error) {
    // Initialize GORM database connection
    db, err := gorm.Open(sqlite.Open("rehabilitation.db"), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    
    // Migrate the schema
    db.AutoMigrate(&Exercise{})
    
    return &DBClient{DB: db}, nil
}

// AddExercise adds a new exercise to the database
func (client *DBClient) AddExercise(exercise *Exercise) error {
    result := client.DB.Create(&exercise)
    return result.Error
}

// GetExercise retrieves an exercise by ID
func (client *DBClient) GetExercise(id uint) (*Exercise, error) {
    var exercise Exercise
    result := client.DB.First(&exercise, id)
    if result.Error != nil {
        return nil, result.Error
    }
    return &exercise, nil
}

// UpdateExercise updates an existing exercise in the database
func (client *DBClient) UpdateExercise(id uint, exercise *Exercise) error {
    result := client.DB.Model(&Exercise{}).Where("id = ?", id).Updates(exercise)
    return result.Error
}

// DeleteExercise removes an exercise from the database
func (client *DBClient) DeleteExercise(id uint) error {
    result := client.DB.Delete(&Exercise{}, id)
    return result.Error
}

func main() {
    dbClient, err := NewDBClient()
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }
    defer dbClient.DB.Close()
    
    // Example usage
    exercise := &Exercise{Name: "Deep Breathing", Description: "Breathe in deeply, hold for a few seconds, then exhale slowly.", Duration: 5}
    if err := dbClient.AddExercise(exercise); err != nil {
        log.Fatal("Failed to add exercise: ", err)
    }
    fmt.Println("Exercise added successfully.")
    
    retrievedExercise, err := dbClient.GetExercise(1)
    if err != nil {
        log.Fatal("Failed to retrieve exercise: ", err)
    }
    fmt.Printf("Retrieved Exercise: %+v
", retrievedExercise)
    
    if err := dbClient.UpdateExercise(1, &Exercise{Description: "New description for deep breathing exercise."}); err != nil {
        log.Fatal("Failed to update exercise: ", err)
    }
    fmt.Println("Exercise updated successfully.")
    
    if err := dbClient.DeleteExercise(1); err != nil {
        log.Fatal("Failed to delete exercise: ", err)
    }
    fmt.Println("Exercise deleted successfully.")
}