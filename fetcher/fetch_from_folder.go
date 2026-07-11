package fetcher

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("韵 - 配置文件错误: ", err)
		os.Exit(1)
		return ""
	}
	return string(content)
}

// 搜索目标路径所在的文件夹内（含子文件夹）所有文件名含 name 且以 .poem 结尾的文件。
func SearchFolder(path, name string) []string {
	// 获取文件信息
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("韵 - 路径不存在")
		} else {
			fmt.Println("韵 - 获取信息时出错:", err)
		}
		return nil
	}

	// 判断是文件还是文件夹
	if !info.IsDir() {
		path = filepath.Dir(path)
	}

	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("韵 - 读取文件夹错误: ", err)
		return nil
	}

	var result []string
	for _, file := range files {
		// 目录一律递归（无论目录名是否含关键词），避免含关键词的子目录被跳过。
		if file.IsDir() {
			subDirResults := SearchFolder(path+"/"+file.Name(), name)
			result = append(result, subDirResults...)
			continue
		}
		// 文件格式是poem则添加到列表
		if strings.Contains(file.Name(), name) && strings.HasSuffix(file.Name(), ".poem") {
			// 将路径符号统一
			path = strings.ReplaceAll(path, "\\", "/")
			result = append(result, path+"/"+file.Name())
		}
	}
	return result
}

// 检查并获取文件夹中是否有指定名字的诗歌文件
func CheckAndGetPoemFile(path, name string) string {

	// 获取文件夹内的诗歌文件
	files := SearchFolder(path, name)
	if len(files) == 0 {
		return ""
	}

	// 返回第一个文件
	return files[0]
}
