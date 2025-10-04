// 代码生成时间: 2025-10-05 01:31:25
package main

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
    "os"
)

// Task represents a task with ID, Name, AssignedTo, and Status
type Task struct {
    ID         uint   `gorm:"primaryKey"`
    Name       string
    AssignedTo string
    Status     string
}

// DB is a global variable for the database connection
var DB *gorm.DB

func main() {
    // Initialize the database connection
    var err error
    DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Migrate the schema
    err = DB.AutoMigrate(&Task{})
    if err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }

    // Example usage of the task assignment system
    err = assignTask("Task 1", "User1\)
    if err != nil {
        log.Printf("Error assigning task: %v", err)
    }
}

// assignTask assigns a task to a user and saves it to the database
func assignTask(name, assignedTo string) error {
    // Create a new task
    task := Task{Name: name, AssignedTo: assignedTo, Status: "Pending"}

    // Save the task to the database
    if result := DB.Create(&task); result.Error != nil {
        return result.Error
    }

    // Log the successful assignment of the task
    log.Printf("Task %s assigned to %s", name, assignedTo)
    return nil
}
