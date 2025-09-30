// 代码生成时间: 2025-09-30 23:18:53
package main

import (
    "errors"
    "fmt"
    "image"
    "image/jpeg"
    "os"
    "path/filepath"
)

// ImageResizer 结构体，用于存储原始图片路径和目标尺寸
type ImageResizer struct {
    SourceDir  string
    TargetDir  string
    Width      int
    Height     int
}

// Resize 函数用于调整图片尺寸
func (ir *ImageResizer) Resize() error {
    files, err := os.ReadDir(ir.SourceDir)
    if err != nil {
        return fmt.Errorf("failed to read directory: %w", err)
    }

    for _, file := range files {
        if !file.IsDir() {
            srcPath := filepath.Join(ir.SourceDir, file.Name())
            dstPath := filepath.Join(ir.TargetDir, file.Name())

            img, err := loadImage(srcPath)
            if err != nil {
                return fmt.Errorf("failed to load image: %w", err)
            }
            
            err = resizeImage(img, dstPath, ir.Width, ir.Height)
            if err != nil {
                return fmt.Errorf("failed to resize image: %w", err)
            }
        }
    }
    return nil
}

// loadImage 函数用于加载图片
func loadImage(path string) (image.Image, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    img, _, err := image.Decode(file)
    if err != nil {
        return nil, err
    }
    return img, nil
}

// resizeImage 函数用于调整图片尺寸
func resizeImage(img image.Image, path string, width, height int) error {
    // 创建一个新图片用于存放调整后的图片
    newImg := image.NewRGBA(image.Rect(0, 0, width, height))

    // 计算缩放比例
    srcW, srcH := img.Bounds().Dx(), img.Bounds().Dy()
    scaleX, scaleY := float64(width)/srcW, float64(height)/srcH

    // 选择缩放比例较小的值，以保持图片比例
    scale := min(scaleX, scaleY)
    newW, newH := int(float64(srcW)*scale), int(float64(srcH)*scale)

    // 创建一个新图片用于存放调整后的图片
    resImg := image.NewRGBA(image.Rect(0, 0, newW, newH))

    // 将原图缩放到新图
    for x := 0; x < newW; x++ {
        for y := 0; y < newH; y++ {
            dstX, dstY := int(float64(x)/scale), int(float64(y)/scale)
            resImg.Set(x, y, img.At(dstX, dstY))
        }
    }

    // 保存新图片
    outFile, err := os.Create(path)
    if err != nil {
        return err
    }
    defer outFile.Close()

    return jpeg.Encode(outFile, resImg, nil)
}

// min 函数返回两个数中的较小值
func min(a, b float64) float64 {
    if a < b {
        return a
    }
    return b
}

func main() {
    // 创建一个ImageResizer实例
    resizer := ImageResizer{
        SourceDir: "./source",
        TargetDir: "./target",
        Width:      800,
        Height:     600,
    }

    // 调整图片尺寸
    err := resizer.Resize()
    if err != nil {
        fmt.Println("Error: ", err)
    } else {
        fmt.Println("Image resizing completed successfully")
    }
}