// 代码生成时间: 2025-10-04 03:26:27
package main

import (
	"fmt"
	"log"
	"net/http"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ModalDialog represents the structure of a modal dialog.
type ModalDialog struct {
	ID       uint   "gorm:"primaryKey" json:"id"
	Title    string "json:"title"
	Content  string "json:"content"
	IsOpen   bool   "json:"isOpen"
}

// ModalDialogService provides operations on modal dialogs.
type ModalDialogService struct {
	db *gorm.DB
}

// NewModalDialogService creates a new ModalDialogService with a database connection.
func NewModalDialogService(db *gorm.DB) *ModalDialogService {
	return &ModalDialogService{db: db}
}

// CreateDialog creates a new modal dialog and saves it to the database.
func (s *ModalDialogService) CreateDialog(title, content string, isOpen bool) (*ModalDialog, error) {
	dialog := &ModalDialog{Title: title, Content: content, IsOpen: isOpen}
	result := s.db.Create(dialog)
	if result.Error != nil {
		return nil, result.Error
	}
	return dialog, nil
}

// OpenDialog opens a modal dialog by setting its `IsOpen` field to true.
func (s *ModalDialogService) OpenDialog(id uint) (*ModalDialog, error) {
	var dialog ModalDialog
	result := s.db.First(&dialog, id)
	if result.Error != nil {
		return nil, result.Error
	}
	dialog.IsOpen = true
	result = s.db.Save(&dialog)
	if result.Error != nil {
		return nil, result.Error
	}
	return &dialog, nil
}

// CloseDialog closes a modal dialog by setting its `IsOpen` field to false.
func (s *ModalDialogService) CloseDialog(id uint) (*ModalDialog, error) {
	var dialog ModalDialog
	result := s.db.First(&dialog, id)
	if result.Error != nil {
		return nil, result.Error
	}
	dialog.IsOpen = false
	result = s.db.Save(&dialog)
	if result.Error != nil {
		return nil, result.Error
	}
	return &dialog, nil
}

// DeleteDialog removes a modal dialog from the database.
func (s *ModalDialogService) DeleteDialog(id uint) error {
	var dialog ModalDialog
	result := s.db.Delete(&dialog, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func main() {
	// Initialize the database connection.
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema.
	db.AutoMigrate(&ModalDialog{})

	// Create a new ModalDialogService with the database connection.
	service := NewModalDialogService(db)

	// Create a new modal dialog.
	dialog, err := service.CreateDialog("Warning", "This is a warning message.", false)
	if err != nil {
		log.Fatal("failed to create dialog")
	}
	fmt.Printf("Created dialog with ID: %d
", dialog.ID)

	// Open the dialog.
	openedDialog, err := service.OpenDialog(dialog.ID)
	if err != nil {
		log.Fatal("failed to open dialog")
	}
	fmt.Printf("Opened dialog with ID: %d
", openedDialog.ID)

	// Close the dialog.
	closedDialog, err := service.CloseDialog(dialog.ID)
	if err != nil {
		log.Fatal("failed to close dialog")
	}
	fmt.Printf("Closed dialog with ID: %d
", closedDialog.ID)

	// Delete the dialog.
	err = service.DeleteDialog(dialog.ID)
	if err != nil {
		log.Fatal("failed to delete dialog")
	}
	fmt.Printf("Deleted dialog with ID: %d
", dialog.ID)
}