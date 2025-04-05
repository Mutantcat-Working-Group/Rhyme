package main

import (
	"flag"
	"fmt"
	"org.mutantcat.rhyme/config"
	"org.mutantcat.rhyme/fetcher"
	"org.mutantcat.rhyme/poem"
	"os"
	"runtime"
	"strings"
)

func main() {
	// 获得命令行参数
	// 定义命令行标志
	file := flag.String("file", "", "目标诗歌文件路径")      // 字符串标志
	folder := flag.String("folder", "", "搜索诗歌文件夹路径") // 字符串标志
	key := flag.String("key", "", "搜索关键字")           // 字符串标志
	//web := flag.String("file", "", "目标诗歌网络路径")  // 字符串标志
	//cloudstep := flag.String("cloudstep", "", "目标诗歌云阶路径")  // 字符串标志
	depth := flag.Int("depth", 5, "最大联调深度")         // 整数标志
	su := flag.Bool("su", false, "是否使用管理员权限运行")     // 布尔标志
	search := flag.Bool("search", false, "是否为搜索模式") // 布尔标志
	flag.Parse()
	// 开始执行操作
	global_config := config.Config{}
	multi_line_string := `
██████╗ ██╗  ██╗██╗   ██╗███╗   ███╗███████╗
██╔══██╗██║  ██║╚██╗ ██╔╝████╗ ████║██╔════╝
██████╔╝███████║ ╚████╔╝ ██╔████╔██║█████╗  
██╔══██╗██╔══██║  ╚██╔╝  ██║╚██╔╝██║██╔══╝  
██║  ██║██║  ██║   ██║   ██║ ╚═╝ ██║███████╗
╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝   ╚═╝     ╚═╝╚══════╝
By: mutantcat.org         诗·韵 v1.0.20250331
`
	fmt.Println(strings.Trim(multi_line_string, "\n"))
	// 获取所有命令行参数
	args := os.Args
	fmt.Println(args)
	global_config.OS_TYPE = runtime.GOOS
	global_config.OS_ARCH = runtime.GOARCH
	global_config.OS_CORE = runtime.NumCPU()
	fmt.Println("↓↓↓↓↓↓↓↓↓↓↓ 韵 - 系统环境信息自检 ↓↓↓↓↓↓↓↓↓↓↓")
	fmt.Println("韵 - 操作系统:	", global_config.OS_TYPE)  // 获取操作系统名称（如 "linux", "windows", "darwin"）
	fmt.Println("韵 - 系统架构:	", global_config.OS_ARCH)  // 获取系统架构（如 "amd64", "386"）
	fmt.Println("韵 - CPU核心:	", global_config.OS_CORE) // 获取 CPU 核心数量

	if runtime.GOOS == "linux" {
		// 读取 os-release 文件
		data, err := os.ReadFile("/etc/os-release")
		if err != nil {
			fmt.Println("错误:", err)
			return
		}

		// 解析内容
		content := string(data)
		lines := strings.Split(content, "\n")
		info := make(map[string]string)
		for _, line := range lines {
			if strings.Contains(line, "=") {
				parts := strings.SplitN(line, "=", 2)
				key := parts[0]
				value := strings.Trim(parts[1], `"'`) // 去掉引号
				info[key] = value
			}
		}
		global_config.OS_NAME = info["NAME"]
		global_config.OS_VER = info["VERSION"]
		global_config.OS_ID = info["ID"]
		global_config.OS_LIKE = info["ID_LIKE"]
		global_config.OS_VER_ID = info["VERSION_ID"]
		// 打印发行版信息
		fmt.Println("韵 - 系统类型:	", global_config.OS_NAME)
		fmt.Println("韵 - 版本号:	", global_config.OS_VER)
		fmt.Println("韵 - ID:	", global_config.OS_ID)
		fmt.Println("韵 - ID_LIKE:	", global_config.OS_LIKE)
		fmt.Println("韵 - 主版本:	", global_config.OS_VER_ID)
	}
	fmt.Println("↑↑↑↑↑↑↑↑↑↑↑ 韵 - 系统环境信息自检 ↑↑↑↑↑↑↑↑↑↑↑")
	// 读取配置文件中所有文本
	// 进行配置读取或搜索
	if *search {
		// 搜索文件
		if *folder == "" {
			fmt.Println("韵 - 请输入文件路径来进行检索")
			return
		}
		fmt.Println("韵 - 开始搜索诗歌...")
		files := fetcher.SearchFolder(*folder, *key)
		if len(files) == 0 {
			fmt.Println("韵 - 没找到任何可用诗歌")
			return
		}
		fmt.Println("韵 - 找到以下诗歌:")

		for index, file := range files {
			fmt.Println("	", index, ": ", file)
		}
		fmt.Println()
		fmt.Println("韵 - 您可以使用 rayme -file xxx 来运行指定的诗歌文件")
	} else {
		content := fetcher.ReadFile(*file)
		fmt.Println("★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★")
		fmt.Println("↓↓↓↓↓↓↓↓↓↓↓ 诗 - 开始运行解析内容 ↓↓↓↓↓↓↓↓↓↓↓")
		// 先解析参数
		args_map := make(map[string]string)
		args_map["OS_TYPE"] = global_config.OS_TYPE
		args_map["OS_ARCH"] = global_config.OS_ARCH
		args_map["OS_CORE"] = fmt.Sprintf("%d", global_config.OS_CORE)
		args_map["OS_NAME"] = global_config.OS_NAME
		args_map["OS_VER"] = global_config.OS_VER
		args_map["OS_ID"] = global_config.OS_ID
		args_map["OS_LIKE"] = global_config.OS_LIKE
		args_map["OS_VER_ID"] = global_config.OS_VER_ID
		args_map["NEED_DEPTH"] = fmt.Sprintf("%d", *depth)
		i := flag.Args()
		for _, arg := range i {
			// 读取参数
			if strings.Contains(arg, "=") {
				// 读取参数
				parts := strings.SplitN(arg, "=", 2)
				key := parts[0]
				value := parts[1]
				args_map[key] = value
			} else {
				fmt.Println("韵 - 无效参数:", arg, "(已跳过)")
			}
		}
		// 开始解析诗词
		parsePoem := poem.ParsePoem(content, args_map)
		parsePoem.Path = *file
		// 读取诗歌的文件名称
		// 先将所有路径符号统一
		*file = strings.ReplaceAll(*file, "\\", "/")
		poemName := strings.Split(*file, "/")
		poemName = strings.Split(poemName[len(poemName)-1], ".")
		parsePoem.Name = poemName[0]
		// 运行诗词
		poem.RunPoem(*parsePoem, *su, 1)
		fmt.Println("↑↑↑↑↑↑↑↑↑↑↑ 诗 - 完成运行解析内容 ↑↑↑↑↑↑↑↑↑↑↑")
	}
}
