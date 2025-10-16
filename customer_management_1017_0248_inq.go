// 代码生成时间: 2025-10-17 02:48:30
package main

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
)

// Customer represents a customer in the CRM system
type Customer struct {
    gorm.Model
    Name        string `gorm:"type:varchar(100);unique"`
    Email       string `gorm:"type:varchar(100);unique"`
    PhoneNumber string `gorm:"type:varchar(20)"`
    Address     string `gorm:"type:text"`
}

// InitializeDB sets up the database connection
func InitializeDB() *gorm.DB {
    // Connect to an SQLite database named "crm.db"
    db, err := gorm.Open(sqlite.Open("crm.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: \{\{err\}\}")
    }
    return db
}

// CreateCustomer inserts a new customer into the database
func CreateCustomer(db *gorm.DB, customer Customer) error {
    result := db.Create(&customer)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

// GetCustomers retrieves all customers from the database
func GetCustomers(db *gorm.DB) ([]Customer, error) {
    var customers []Customer
    result := db.Find(&customers)
    if result.Error != nil {
        return nil, result.Error
    }
    return customers, nil
}

// UpdateCustomer updates an existing customer in the database
func UpdateCustomer(db *gorm.DB, customer Customer) error {
    result := db.Save(&customer)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

// DeleteCustomer removes a customer from the database
func DeleteCustomer(db *gorm.DB, id uint) error {
    result := db.Delete(&Customer{}, id)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func main() {
    db := InitializeDB()
    defer db.Migrator().DropTable(&Customer{})
    db.AutoMigrate(&Customer{})
    
    // Example usage of the CRM system
    customer := Customer{Name: "John Doe", Email: "johndoe@example.com"}
    if err := CreateCustomer(db, customer); err != nil {
        log.Printf("Failed to create customer: \{\{err\}\}")
    } else {
        log.Println("Customer created successfully")
    }
    
    customers, err := GetCustomers(db)
    if err != nil {
        log.Printf("Failed to get customers: \{\{err\}\}")
    } else {
        log.Printf("Customers: \{\{customers\}\}")
    }
}
