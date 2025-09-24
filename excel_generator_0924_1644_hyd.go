// 代码生成时间: 2025-09-24 16:44:41
package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
    "time"

    "github.com/xuri/excelize/v2"
)

// ExcelGenerator is the struct that holds configuration for Excel generation.
type ExcelGenerator struct {
    Filepath string
    Filename string
}

// NewExcelGenerator creates a new instance of ExcelGenerator.
func NewExcelGenerator(filepath, filename string) *ExcelGenerator {
    return &ExcelGenerator{
        Filepath: filepath,
        Filename: filename,
    }
}

// Generate creates a new Excel file with the given configuration.
func (e *ExcelGenerator) Generate() error {
    // Create a new Excel file.
    f := excelize.NewFile()
    defer f.Close()

    // Generate the file name with date and time.
    filename := fmt.Sprintf("%s_%s.xlsx", e.Filename, time.Now().Format("20060102_150405"))
    filepath := filepath.Join(e.Filepath, filename)

    // Create a new sheet.
    index := f.NewSheet(excelize.DefaultSheetName)

    // Write sample data to the sheet.
    f.SetCellValue(excelize.DefaultSheetName, "A1", "ID")
    f.SetCellValue(excelize.DefaultSheetName, "B1", "Name")
    f.SetCellValue(excelize.DefaultSheetName, "C1", "Date")

    // Add data to the sheet.
    if err := f.SaveAs(filepath); err != nil {
        return err
    }

    log.Printf("Excel file '%s' has been created successfully.", filename)
    return nil
}

// WriteData adds data to the Excel file.
func (e *ExcelGenerator) WriteData(data [][]string) error {
    // Check if the generator has been initialized properly.
    if e.Filepath == "" || e.Filename == "" {
        return fmt.Errorf("file path or filename is not set")
    }

    // Create a new Excel file or open an existing one.
    f, err := excelize.OpenFile(e.Filepath)
    if err != nil {
        return err
    }
    defer f.Close()

    // Start writing data from the second row.
    for i, row := range data {
        for j, value := range row {
            cell := fmt.Sprintf("A%d", i+2)
            if j > 0 {
                cell = fmt.Sprintf("%c%d", 'A'+j-1, i+2)
            }
            f.SetCellValue(excelize.DefaultSheetName, cell, value)
        }
    }

    // Save changes to the file.
    filename := fmt.Sprintf("%s_%s.xlsx", e.Filename, time.Now().Format("20060102_150405"))
    if err := f.SaveAs(filepath.Join(e.Filepath, filename)); err != nil {
        return err
    }

    log.Printf("Data has been written to '%s' successfully.", filename)
    return nil
}

func main() {
    // Example usage of ExcelGenerator.
    filepath := "./"
    filename := "example"
    generator := NewExcelGenerator(filepath, filename)

    if err := generator.Generate(); err != nil {
        log.Fatal(err)
    }

    // Sample data to write to the Excel file.
    data := [][]string{
        {"1", "John Doe", time.Now().Format("2006-01-02")},
        {"2", "Jane Doe", time.Now().Format("2006-01-02")},
    }

    if err := generator.WriteData(data); err != nil {
        log.Fatal(err)
    }
}
