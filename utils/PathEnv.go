package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func UpdatePathEnv(dir string) error {
	// 获取当前系统的 PATH 环境变量值
	path := os.Getenv("PATH")
	// 检查路径是否已存在于 PATH 中
	if !strings.Contains(path, dir) {
		// 如果路径不存在，则将其添加到 PATH 中
		dllPath, _ := filepath.Abs(dir)
		newPath := fmt.Sprintf("%s%s%s", dllPath, string(os.PathListSeparator), path)
		// 设置更新后的 PATH 环境变量值
		err := os.Setenv("PATH", newPath)
		if err != nil {
			return err
		}
	}
	return nil
}
