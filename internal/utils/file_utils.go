package utils

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

// CreateFile 创建文件
func CreateFile(filePath string) error {
	// 检查文件路径是否存在，如果不存在则创建目录
	dirPath := filepath.Dir(filePath)
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0755) // 创建所有必要的父目录
		if err != nil {
			return err
		}
	}

	// 创建文件
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			slog.Error(fmt.Sprintf("Error closing file: %s", err))
		}
	}(file) // 确保在函数返回时关闭文件
	return nil
}

// IsExist 检查给定的路径是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// SplitFilePath 分解文件路径为文件夹路径、文件名和文件后缀
func SplitFilePath(filePath string) (dirPath, baseName, ext string) {
	// 使用 filepath.Clean 来规范化路径，这会处理相对路径和绝对路径
	cleanedPath := filepath.Clean(filePath)
	// 提取文件的目录路径
	dirPath = filepath.Dir(cleanedPath)

	// 提取文件的完整名字
	baseName = filepath.Base(filePath)

	// 分割文件名和后缀
	if strings.Contains(baseName, ".") {
		ext = strings.Join(strings.Split(baseName, ".")[1:], ".")
		baseName = strings.Split(baseName, ".")[0]
	} else {
		// 没有后缀的情况
		ext = ""
	}
	return
}
