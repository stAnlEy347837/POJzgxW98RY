// 代码生成时间: 2025-09-23 00:43:30
package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
    "time"
)

// ExcelData 数据结构，用于存储将要写入Excel的数据
type ExcelData struct {
    Timestamp time.Time `csv:"timestamp"`
    Value     string    `csv:"value"`
}

// GenerateExcel 生成Excel文件
func GenerateExcel(filename string, data []ExcelData) error {
    // 创建CSV文件
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    // 创建CSV写入器
    writer := csv.NewWriter(file)
    defer writer.Flush()

    // 写入标题
    header := []string{"Timestamp", "Value"}
    if err := writer.Write(header); err != nil {
        return err
    }

    // 写入数据
    for _, record := range data {
        if err := writer.Write([]string{record.Timestamp.Format(time.RFC3339), record.Value}); err != nil {
            return err
        }
    }

    // 检查是否有写入错误
    if err := writer.Error(); err != nil {
        return err
    }

    return nil
}

func main() {
    // 示例数据
    data := []ExcelData{
        {Timestamp: time.Now(), Value: "Example Value 1"},
        {Timestamp: time.Now(), Value: "Example Value 2"},
    }

    // 生成Excel文件
    filename := "example_data.csv"
    if err := GenerateExcel(filename, data); err != nil {
        log.Fatalf("Failed to generate Excel file: %v", err)
    } else {
        fmt.Printf("Excel file generated successfully: %s
", filename)
    }
}
