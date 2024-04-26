package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func CreateLogDir(Path string) {
	_, err := os.Stat(Path)

	// 如果文件夹不存在，则创建它
	if os.IsNotExist(err) {
		err := os.Mkdir(Path, 0755)
		if err != nil {
			fmt.Println("Error creating folder:", err)
			return
		}
	} else if err == nil {
		fmt.Println("Folder already exists. Skipping creation.")
		return
	} else {
		fmt.Println("Error checking folder existence:", err)
		return
	}
}
func CreateProject(Path string, ProjectName string) {
	projectPath := filepath.Join(Path, ProjectName)
	if _, err := os.Stat(projectPath); os.IsNotExist(err) {
		err := os.MkdirAll(projectPath, 0755)
		if err != nil {
			fmt.Println("Error creating project folder:", err)
			return
		}
	}

}
func ListDirectory(Path string) []string {
	Projects := []string{"Create New Project"}
	d, err := os.Open(Path)
	if err != nil {
		fmt.Println("Error opening directory:", err)
		return nil
	}
	defer d.Close()

	// 读取目录中的文件和文件夹列表
	files, err := d.Readdir(-1)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return nil
	}

	// 遍历文件和文件夹列表
	for _, file := range files {
		// 如果是文件夹，则打印文件夹名
		if file.IsDir() {
			Projects = append(Projects, file.Name())
		}
	}
	return Projects
}
func CreateLog(TargetIp, LogDir, ProjectName string) {
	now := time.Now()
	formattedTime := now.Format("2006-01-02.15.04.05.000000")
	path, err := filepath.Abs(filepath.Join(LogDir, ProjectName, fmt.Sprintf("z%s", TargetIp)))
	for _, v := range ListDirectory(path) {
		if strings.Contains(v, "fuzzbunch") {
			return
		}
	}
	if err != nil {
		return
	}
	filename := fmt.Sprintf("fuzzbunch-%s.log", formattedTime)
	rawData := fmt.Sprintf(`
[*] Initializing Global State
[+] Set TargetIp => %s
[+] Set CallbackIp => 127.0.0.1

[!] Redirection OFF
[+] Set LogDir => %s
`, TargetIp, path)
	file, err := os.Create(filepath.Join(path, filename))
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // 确保文件关闭

	// 写入内容到文件
	_, err = file.WriteString(rawData)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
func CheckPathLegal(Path string) bool {
	_, err := os.Stat(Path)
	if err != nil {
		if os.IsNotExist(err) {
			return true // Path does not exist
		}
		return false
	}
	return true
}
