package exec

import (
	"context"
	"os/exec"
	"runtime"
)

// Run 在指定平台的 shell 中执行一行命令。
// Windows 用 cmd /C，其他系统（linux/darwin/bsd 等类 Unix）统一用 bash -c，
// su 为真时改用 sudo bash -c。该函数永远不返回 nil。
func Run(ctx context.Context, line string, su bool) *exec.Cmd {
	if runtime.GOOS == "windows" {
		return exec.CommandContext(ctx, "cmd", "/C", line)
	}
	if su {
		return exec.CommandContext(ctx, "sudo", "bash", "-c", line)
	}
	return exec.CommandContext(ctx, "bash", "-c", line)
}
