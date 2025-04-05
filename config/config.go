package config

type Config struct {
	OS_TYPE   string // 操作系统 linux, windows, darwin
	OS_ARCH   string // 系统架构 amd64, 386
	OS_CORE   int    // CPU核心数量
	OS_NAME   string // 系统类型
	OS_VER    string // 版本号
	OS_ID     string // ID
	OS_LIKE   string // ID_LIKE
	OS_VER_ID string // 主版本
}
